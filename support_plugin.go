package main

import (
	"bufio"
	"fmt"
	"github.com/cloudfoundry/cli/plugin"
	"github.com/dmikusa-pivotal/support_plugin/ticket_log"
	"github.com/sendgrid/sendgrid-go"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type SupportPlugin struct {
	TicketLog ticket_log.TicketLog
}

func runSingleCommand(cliConnection plugin.CliConnection, cmd ...string) (te ticket_log.TicketEntry) {
	output, err := cliConnection.CliCommandWithoutTerminalOutput(cmd...)
	if err != nil {
		fmt.Println("Plugin Error:", err)
	}
	te = ticket_log.TicketEntry{
		Description: fmt.Sprintf("Output from `%s`", strings.Join(cmd, " ")),
		Body:        output,
	}
	return
}

func PromptForYesNo(reader io.Reader, question string) bool {
	fmt.Println(question, "(y/n)")
	cnt := 4
	r := bufio.NewReader(reader)
	line, err := r.ReadString('\n')
	for cnt > 0 && err != nil {
		fmt.Println("Sorry, I didn't get that.  Please try again.")
		fmt.Println(question, "(y/n)")
		line, err = r.ReadString('\n')
		cnt -= 1
	}
	line = strings.TrimSpace(line)
	if cnt > 0 && (line == "y" || line == "Y") {
		return true
	} else {
		return false
	}
}

func PromptForString(reader io.Reader, question string) []string {
	// TODO: refactor the PromptFor methods
	fmt.Println(question)
	cnt := 4
	r := bufio.NewReader(reader)
	line, err := r.ReadString('\n')
	for cnt > 0 && err != nil {
		fmt.Println("Sorry, I didn't get that.  Please try again.")
		fmt.Println(question)
		line, err = r.ReadString('\n')
		cnt -= 1
	}
	// TODO: support multiple lines
	return []string{strings.TrimSpace(line)}
}

func (sp *SupportPlugin) OpenTicket() {
	sg := sendgrid.NewSendGridClient("uJgAGdippC", "2AyGJ6Wc3d")
	message := sendgrid.NewMail()
	message.AddTo("svennela@pivotal.io")
	message.AddToName("Pivotal Support")
	message.SetSubject("New CF Support Ticket")
	message.SetText("New Ticket from User.  See attachment.")
	fi, err := os.Open(sp.TicketLog.Name)
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(fi)

	message.AddAttachment(sp.TicketLog.Name, r)
	message.SetFrom("support@run.pivotal.io")
	if r := sg.Send(message); r == nil {
		fmt.Println("Ticket Opened!  You should receive an email confirmation shortly.")
	} else {
		fmt.Println(r)
	}
}

func (sp *SupportPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	if args[0] == "help-me" {
		fmt.Println("Don't Panic!  We're gathering some information. Please hold.\n")
		sp.TicketLog = ticket_log.NewTicketLog()
		sp.TicketLog.Append(runSingleCommand(cliConnection, "target"))
		sp.TicketLog.Append(runSingleCommand(cliConnection, "apps"))
		sp.TicketLog.Append(runSingleCommand(cliConnection, "services"))
		ans := PromptForYesNo(os.Stdin,
			"Are you currently experiencing a problem with one appliation in particular?")
		if ans {
			// TODO: check that app name exists
			// TODO: print list of known apps
			appName := PromptForString(os.Stdin,
				"Please enter the name of the failing application?")
			sp.TicketLog.Append(runSingleCommand(cliConnection, "app", appName[0]))
			sp.TicketLog.Append(runSingleCommand(cliConnection, "logs", appName[0], "--recent"))
		}
		description := PromptForString(os.Stdin,
			"Please enter anything else you'd like to mention about the issue.")
		sp.TicketLog.Append(ticket_log.TicketEntry{
			Description: "Customer's Problem Summary",
			Body:        description,
		})
		fmt.Println("We've gathered the following information.  Please review.\n")
		fmt.Println("-------------------------------------------------------------")
		data, _ := ioutil.ReadFile(sp.TicketLog.Name)
		fmt.Println(string(data[:]))
		fmt.Println("-------------------------------------------------------------")
		fmt.Println("")
		openTicket := PromptForYesNo(os.Stdin,
			"Last chance, are you sure you want to open a ticket?")
		if openTicket {
			sp.OpenTicket()
		}
	}
}

func (c *SupportPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "Support",
		Commands: []plugin.Command{
			{
				Name:     "help-me",
				HelpText: "Get help using Cloud Foundry",
			},
		},
	}
}

func main() {
	plugin.Start(new(SupportPlugin))
}
