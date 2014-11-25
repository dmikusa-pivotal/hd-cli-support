package main

import (
	"github.com/cloudfoundry/cli/plugin"
)

type Support struct{}

func (c *Support) Run(cliConnection plugin.CliConnection, args []string) {
}

func (c *Support) GetMetadata() plugin.PluginMetadata {
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
	plugin.Start(new(Support))
}
