package main

import (
	"github.com/cloudfoundry/cli/plugin"
)

type SupportPlugin struct{}

func (c *SupportPlugin) Run(cliConnection plugin.CliConnection, args []string) {
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
