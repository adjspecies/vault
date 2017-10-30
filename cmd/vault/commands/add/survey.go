// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package add

import (
	"github.com/adjspecies/vault/cmd/vault/commands/command"
	"github.com/adjspecies/vault/config"
)

// SurveyCommand allows adding a survey to the vault.
type SurveyCommand command.BaseCommand

// Init initializes the command
func (cmd SurveyCommand) Init(cfg *config.Config, args []string) error {
	cmd.Config = cfg
	return nil
}

// Run runs the command
func (cmd SurveyCommand) Run() error {
	return nil
}

// NewSurveyCommand returns a RegisteredCommand wrapping a SurveyCommand.
func NewSurveyCommand() *command.RegisteredCommand {
	return &command.RegisteredCommand{
		Name:    "Add a survey",
		Command: "add-survey",
		Help:    ``,
		Entry:   &SurveyCommand{},
	}
}
