package main

import (
	"bufio"
	"fmt"
	"github.com/cloudfoundry/cli/plugin"
	"github.com/dmikusa-pivotal/support_plugin/ticket_log"
	"io"
	"io/ioutil"
	"strings"
)

type SupportPlugin struct {
	TicketLog ticket_log.TicketLog
}

func runSingleCommand(cliConnection plugin.CliConnection, cmd string) (te ticket_log.TicketEntry) {
	output, err := cliConnection.CliCommandWithoutTerminalOutput(cmd)
	if err != nil {
		fmt.Println("Plugin Error:", err)
	}
	te = ticket_log.TicketEntry{
		Description: fmt.Sprintf("Output from `%s`", cmd),
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
	fmt.Println("line:", line)
	if cnt > 0 && (line == "y" || line == "Y") {
		return true
	} else {
		return false
	}
}

func (sp *SupportPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	fmt.Println("Don't Panic!  We're gathering some information. Please hold.\n")
	sp.TicketLog = ticket_log.NewTicketLog()
	sp.TicketLog.Append(runSingleCommand(cliConnection, "target"))
	sp.TicketLog.Append(runSingleCommand(cliConnection, "apps"))
	sp.TicketLog.Append(runSingleCommand(cliConnection, "services"))
	fmt.Println("We've gathered the following information.  Please review.\n")
	fmt.Println("-------------------------------------------------------------")
	data, _ := ioutil.ReadFile(sp.TicketLog.Name)
	fmt.Println(string(data[:]))
	fmt.Println("-------------------------------------------------------------")
	fmt.Println("")
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
