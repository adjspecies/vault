// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package add

import (
	"github.com/adjspecies/vault/cmd/vault/commands/command"
	"github.com/adjspecies/vault/config"
)

// SourceCommand allows adding a source to the vault.
type SourceCommand command.BaseCommand

// Init initializes the command.
func (cmd SourceCommand) Init(cfg *config.Config, args []string) error {
	cmd.Config = cfg
	return nil
}

// Run runs the command
func (cmd SourceCommand) Run() error {
	return nil
}

// NewSourceCommand returns a RegisteredCommand wrapping a SourceCommand.
func NewSourceCommand() *command.RegisteredCommand {
	return &command.RegisteredCommand{
		Name:    "Add source",
		Command: "add-source",
		Help:    ``,
		Entry:   &SourceCommand{},
	}
}
