// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package add

import (
	"github.com/adjspecies/vault/cmd/vault/commands/command"
	"github.com/adjspecies/vault/config"
)

// RespondentCommand allows adding a respondent to the vault.
type RespondentCommand command.BaseCommand

// Init initializes the command
func (cmd RespondentCommand) Init(cfg *config.Config, args []string) error {
	cmd.Config = cfg
	return nil
}

// Run runs the command
func (cmd RespondentCommand) Run() error {
	return nil
}

// NewRespondentCommand returns a RegisteredCommand wrapping a
// RespondentCommand.
func NewRespondentCommand() *command.RegisteredCommand {
	return &command.RegisteredCommand{
		Name:    "Add a respondent",
		Command: "add-respondent",
		Help:    ``,
		Entry:   &RespondentCommand{},
	}
}
