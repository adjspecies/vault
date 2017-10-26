// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package add

import (
	"github.com/adjspecies/vault/cmd/vault/commands/command"
)

type AddSourceCommand struct{}

func (cmd AddSourceCommand) Init(args []string) error {
	return nil
}

func (cmd AddSourceCommand) Run() error {
	return nil
}

func NewAddSourceCommand() *command.RegisteredCommand {
	return &command.RegisteredCommand{
		Name:    "Add source to source",
		Command: "add-source-to-source",
		Help:    ``,
		Entry:   AddSourceCommand{},
	}
}
