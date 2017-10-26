// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package main

import (
	"fmt"
	"os"

	"github.com/adjspecies/vault/cmd/vault/commands"
	"github.com/adjspecies/vault/config"
	"github.com/adjspecies/vault/logging"
)

var specifyConfig = `%s

Vault requires a YAML configuration file. This may live anywhere, and is
specified through an environment variable, VAULT_CONFIG.

To run a vault command, you may set the environment variable globally, or run:

	VAULT_CONFIG=path/to/config.yaml vault <command>

	# Or, for fish:
	env VAULT_CONFIG=path/to/config.yaml vault <command>
`

func main() {
	var command string
	if len(os.Args) < 2 {
		command = "help"
	} else {
		command = os.Args[1]
	}
	args := []string{}
	if len(os.Args) > 2 {
		args = os.Args[2:]
	}

	cfg, err := config.Read(os.Getenv("VAULT_CONFIG"))
	if err != nil {
		fmt.Fprintf(os.Stderr, specifyConfig, err)
		os.Exit(1)
	}
	if errSetup := logging.Setup(cfg.LogLevel); errSetup != nil {
		fmt.Fprintln(os.Stderr, "could not set up logging")
		os.Exit(2)
	}
	log := logging.Logger()
	defer log.Sync()
	log.Debug("logging set up")

	commands.Main(cfg, command, args)
}
