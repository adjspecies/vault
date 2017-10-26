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

	serve					    Start the Vault server
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

var specifyConfig = `Could not load config file.

Vault requires a YAML configuration file. This may live anywhere, and is
specified through an environment variable, VAULT_CONFIG.

To run a vault command, you may set the environment variable globally, or run:

	VAULT_CONFIG=path/to/config.yaml vault <command>

	# Or, for fish:
	env VAULT_CONFIG=path/to/config.yaml vault <command>`

var registeredCommands map[string]*command.RegisteredCommand

func registerCommand(cmd *command.RegisteredCommand) {
	registeredCommands[cmd.Command] = cmd
}

// RegisterCommands adds subcommands to the registry of available commands.
func RegisterCommands() {
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
func Main(command string, args []string) {
	RegisterCommands()

	cfg, err := config.Read(os.Getenv("VAULT_CONFIG"))
	if err != nil {
		fmt.Fprintln(os.Stderr, specifyConfig)
		os.Exit(1)
	}
	if errSetup := logging.Setup(cfg.LogLevel); errSetup != nil {
		fmt.Fprintln(os.Stderr, "could not set up logging")
		os.Exit(2)
	}
	log = logging.Logger()
	defer log.Sync()

	// Setup logging
	level, err := logging.LevelFromString("debug")
	if err != nil {
		fmt.Fprint(os.Stderr, "unable to set log level: debug")
		os.Exit(3)
	}
	if errSetup := logging.Setup(level); errSetup != nil {
		fmt.Fprint(os.Stderr, "unable to set up logging")
		os.Exit(4)
	}
	log = logging.Logger()
	defer log.Sync()
	log.Info("logging set up")

	// Run help if needed
	if command == "help" {
		help(args[0])
	}

	// Attempt to find the requested command.
	cmd, ok := registeredCommands[command]
	if !ok {
		log.Fatalf("could not find command %s", command)
		os.Exit(5)
	}
	// Initialize and run the command.
	cmd.Entry.Init(args)
	cmd.Entry.Run()
}

func help(topic string) {
	// If we have no topic, print the master help.
	if topic == "" {
		fmt.Print(masterHelp)
		os.Exit(0)
	}

	// If the topic was 'commands', list the commands and their names.
	if topic == "commands" {
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		for command, entry := range registeredCommands {
			fmt.Fprintf(w, "%s\t%s", command, entry.Name)
		}
		w.Flush()
		os.Exit(0)
	}

	// Otherwise, attempt to print the help for a given command.
	cmd, ok := registeredCommands[topic]
	if !ok {
		log.Fatalf("could not find command %s", topic)
		os.Exit(5)
	}
	fmt.Print(cmd.Help)
	os.Exit(0)
}
