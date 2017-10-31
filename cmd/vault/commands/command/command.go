// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

// Package command manages the basics of subcommands.
package command

import (
	"github.com/adjspecies/vault/config"
)

// BaseCommand represents a Vault subcommand. It contains the current config.
// Subcommands should be based off of this and implement Command.
type BaseCommand struct {
	Config *config.Config
}

// Command represents a subcommand. It has an `Init` method and a `Run` method.
type Command interface {
	Init(*config.Config, []string) error
	Run() error
}

// RegisteredCommand represents a command, plus additional data such as a short
// name, the text used to call the subcommand, and a help string.
type RegisteredCommand struct {
	Name    string
	Command string
	Help    string
	Entry   Command
}
