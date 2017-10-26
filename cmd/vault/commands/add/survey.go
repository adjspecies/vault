// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package add

import (
	"github.com/adjspecies/vault/cmd/vault/commands/command"
	"github.com/adjspecies/vault/config"
)

type AddSurveyCommand struct {
	cfg *config.Config
}

func (cmd AddSurveyCommand) Init(cfg *config.Config, args []string) error {
	cmd.cfg = cfg
	return nil
}

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
