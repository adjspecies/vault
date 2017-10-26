// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package add

import (
	"github.com/adjspecies/vault/cmd/vault/commands/command"
)

type AddSurveyCommand struct{}

func (cmd AddSurveyCommand) Init(args []string) error {
	return nil
}

func (cmd AddSurveyCommand) Run() error {
	return nil
}

func NewAddSurveyCommand() *command.RegisteredCommand {
	return &command.RegisteredCommand{
		Name:    "Add source to source",
		Command: "add-source-to-source",
		Help:    ``,
		Entry:   AddSurveyCommand{},
	}
}
