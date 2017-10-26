// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package commands

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/adjspecies/vault/cmd/vault/commands/add"
	"github.com/adjspecies/vault/cmd/vault/commands/command"
	"github.com/adjspecies/vault/cmd/vault/commands/hierarchy"
	"github.com/adjspecies/vault/cmd/vault/commands/serve"
	"github.com/adjspecies/vault/config"
	"github.com/adjspecies/vault/logging"
	"go.uber.org/zap"
)

var log *zap.SugaredLogger

var masterHelp = `vault [help] <command> [arguments]

Vault provides a means of importing data and editing structure through the
command line.

Common commands:

    serve                       Start the Vault server
    add-source                  Adds a data source
    add-survey                  Adds a dataset
    add-source-to-source        Adds an existing source to a hierarchy
    add-survey-to-source        Adds a dataset to a source
    remove-source-from-source   Removes a source from a hierarchy
    remove-survey-from-source   Removes a dataset from a source

Example help commands:

    vault help
    vault help commands
    vault help add-source
`

var registeredCommands = make(map[string]*command.RegisteredCommand)
var commandList []string

func registerCommand(cmd *command.RegisteredCommand) {
	log.Debugf("registering command %s", cmd.Command)
	registeredCommands[cmd.Command] = cmd
	commandList = append(commandList, cmd.Command)
}

// RegisterCommands adds subcommands to the registry of available commands.
func RegisterCommands() {
	log.Debug("registering commands")
	registerCommand(add.NewAddSourceCommand())
	registerCommand(add.NewAddSurveyCommand())
	registerCommand(hierarchy.NewAddSourceToSourceCommand())
	registerCommand(hierarchy.NewRemoveSourceFromSourceCommand())
	registerCommand(hierarchy.NewAddSurveyToSourceCommand())
	registerCommand(hierarchy.NewRemoveSurveyFromSourceCommand())
	registerCommand(serve.NewServeCommand())
}

// Main takes a command and a list of arguments and attempts to run the matching
// subcommand, using those arguments.
func Main(cfg *config.Config, command string, args []string) {
	log = logging.Logger()
	RegisterCommands()

	// Run help if needed
	if command == "help" {
		help(args)
	}
	// Attempt to find the requested command.
	cmd, ok := registeredCommands[command]
	if !ok {
		log.Errorf("could not find command %s", command)
		os.Exit(3)
	}
	// Initialize and run the command.
	cmd.Entry.Init(cfg, args)
	cmd.Entry.Run()
}

func help(args []string) {
	// If we have no topic, print the master help.
	if len(args) == 0 {
		fmt.Print(masterHelp)
		log.Sync()
		os.Exit(0)
	}
	topic := args[0]

	// If the topic was 'commands', list the commands and their names.
	if topic == "commands" {
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		for _, command := range commandList {
			fmt.Fprintf(w, "%s\t%s\n", command, registeredCommands[command].Name)
		}
		w.Flush()
		os.Exit(0)
	}

	// Otherwise, attempt to print the help for a given command.
	cmd, ok := registeredCommands[topic]
	if !ok {
		log.Errorf("could not find command %s", topic)
		os.Exit(3)
	}
	fmt.Print(cmd.Help)
	os.Exit(0)
}
