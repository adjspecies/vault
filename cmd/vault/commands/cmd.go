// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package commands

import (
	"fmt"
	"os"
	"text/tabwriter"

	"go.uber.org/zap"
	errgo "gopkg.in/errgo.v1"

	"github.com/adjspecies/vault/cmd/vault/commands/add"
	"github.com/adjspecies/vault/cmd/vault/commands/command"
	"github.com/adjspecies/vault/cmd/vault/commands/hierarchy"
	"github.com/adjspecies/vault/cmd/vault/commands/serve"
	"github.com/adjspecies/vault/config"
	"github.com/adjspecies/vault/logging"
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

func RegisterCommand(cmd *command.RegisteredCommand) {
	log.Debugf("registering command %s", cmd.Command)
	registeredCommands[cmd.Command] = cmd
	commandList = append(commandList, cmd.Command)
}

// RegisterCommands adds subcommands to the registry of available commands.
func RegisterCommands() {
	log.Debug("registering commands")
	RegisterCommand(add.NewAddSourceCommand())
	RegisterCommand(add.NewAddSurveyCommand())
	RegisterCommand(hierarchy.NewAddSourceToSourceCommand())
	RegisterCommand(hierarchy.NewRemoveSourceFromSourceCommand())
	RegisterCommand(hierarchy.NewAddSurveyToSourceCommand())
	RegisterCommand(hierarchy.NewRemoveSurveyFromSourceCommand())
	RegisterCommand(serve.NewServeCommand())
}

// Main takes a command and a list of arguments and attempts to run the matching
// subcommand, using those arguments.
func Main(cfg *config.Config, command string, args []string) error {
	log = logging.Logger()
	RegisterCommands()

	// Run help if needed
	if command == "help" {
		return Help(args)
	}
	// Attempt to find the requested command.
	cmd, ok := registeredCommands[command]
	if !ok {
		err := errgo.Newf("could not find command %s", command)
		log.Errorf(err.Error())
		return err
	}
	// Initialize and run the command.
	err := cmd.Entry.Init(cfg, args)
	if err != nil {
		return err
	}
	err = cmd.Entry.Run()
	if err != nil {
		return err
	}
	return nil
}

func Help(args []string) error {
	// If we have no topic, print the master help.
	if len(args) == 0 {
		fmt.Print(masterHelp)
		return nil
	}
	topic := args[0]

	// If the topic was 'commands', list the commands and their names.
	if topic == "commands" {
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		for _, command := range commandList {
			fmt.Fprintf(w, "%s\t%s\n", command, registeredCommands[command].Name)
		}
		w.Flush()
		return nil
	}

	// Otherwise, attempt to print the help for a given command.
	cmd, ok := registeredCommands[topic]
	if !ok {
		err := errgo.Newf("could not find command %s", topic)
		log.Errorf(err.Error())
		return err
	}
	fmt.Print(cmd.Help)
	return nil
}
