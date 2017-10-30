// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package add

import (
	"github.com/adjspecies/vault/cmd/vault/commands/command"
	"github.com/adjspecies/vault/config"
)

// SummaryCommand allows adding a summary to the vault.
type SummaryCommand command.BaseCommand

// Init initializes the command
func (cmd SummaryCommand) Init(cfg *config.Config, args []string) error {
	cmd.Config = cfg
	return nil
}

// Run runs the command
func (cmd SummaryCommand) Run() error {
	return nil
}

// NewSummaryCommand returns a RegisteredCommand wrapping a SummaryCommand.
func NewSummaryCommand() *command.RegisteredCommand {
	return &command.RegisteredCommand{
		Name:    "Add a summary",
		Command: "add-summary",
		Help:    ``,
		Entry:   &SummaryCommand{},
	}
}
