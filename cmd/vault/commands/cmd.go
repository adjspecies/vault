// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

// Package commands manages the command line Vault tool.
//
// It allows for subcommands to be loaded into a registry. These are called
// via the first argument to the `vault` command. It also provides the means to
// get help on the `vault` command as well as all subcommands.
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

// registerCommand adds a single subcommand to the registry of commands.
func registerCommand(cmd *command.RegisteredCommand) {
	log.Debugf("registering command %s", cmd.Command)
	registeredCommands[cmd.Command] = cmd
	commandList = append(commandList, cmd.Command)
}

// registerCommands adds subcommands to the registry of available commands.
func registerCommands() {
	log.Debug("registering commands")
	registerCommand(add.NewSourceCommand())
	registerCommand(add.NewSummaryCommand())
	registerCommand(add.NewSurveyCommand())
	registerCommand(add.NewResponseCommand())
	registerCommand(add.NewResponsesCommand())
	registerCommand(add.NewRespondentCommand())
	registerCommand(hierarchy.NewAddSourceToSourceCommand())
	registerCommand(hierarchy.NewRemoveSourceFromSourceCommand())
	registerCommand(hierarchy.NewAddSummaryToSourceCommand())
	registerCommand(hierarchy.NewRemoveSummaryFromSourceCommand())
	registerCommand(hierarchy.NewAddSurveyToSourceCommand())
	registerCommand(hierarchy.NewRemoveSurveyFromSourceCommand())
	registerCommand(serve.NewServeCommand())
}

// Main takes a command and a list of arguments and attempts to run the matching
// subcommand, using those arguments.
func Main(cfg *config.Config, command string, args []string) error {
	log = logging.Logger()
	registerCommands()

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

// Help handles showing help to the user. It accepts a slice of strings. If this
// is empty, it prints the master help for Vault. If not, it assumes that the
// first argument is the topic one would want help on. If this is `commands`,
// it provides a list of commands with their names. Otherwise, it attempts to
// provide help for the provided command.
func Help(args []string) error {
	log = logging.Logger()
	// If we have no topic, print the master help.
	if len(args) == 0 {
		log.Debug("printing masterHelp")
		fmt.Print(masterHelp)
		return nil
	}
	topic := args[0]

	// If the topic was 'commands', list the commands and their names.
	if topic == "commands" {
		log.Debug("listing commands")
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
	log.Debugf("listing help for %s", topic)
	fmt.Print(cmd.Help)
	return nil
}
