// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package add

import (
	"github.com/adjspecies/vault/cmd/vault/commands/command"
	"github.com/adjspecies/vault/config"
)

type AddSurveyCommand command.BaseCommand

// Init initializes the command
func (cmd AddSurveyCommand) Init(cfg *config.Config, args []string) error {
	cmd.Config = cfg
	return nil
}

// Run runs the command
func (cmd AddSurveyCommand) Run() error {
	return nil
}

func NewAddSurveyCommand() *command.RegisteredCommand {
	return &command.RegisteredCommand{
		Name:    "Add a survey",
		Command: "add-survey",
		Help:    ``,
		Entry:   AddSurveyCommand{},
	}
}
