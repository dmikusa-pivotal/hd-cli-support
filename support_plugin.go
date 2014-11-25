package main

import (
	"fmt"
	"github.com/cloudfoundry/cli/plugin"
	"github.com/dmikusa-pivotal/support_plugin/ticket_log"
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

func (sp *SupportPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	fmt.Println("Don't Panic!  We're gathering some information. Please hold.")
	sp.TicketLog = ticket_log.NewTicketLog()
	sp.TicketLog.Append(runSingleCommand(cliConnection, "target"))
	sp.TicketLog.Append(runSingleCommand(cliConnection, "apps"))
	sp.TicketLog.Append(runSingleCommand(cliConnection, "services"))
	fmt.Println("Ticket data has been gathered.  It's located here: ", sp.TicketLog.Name)
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
