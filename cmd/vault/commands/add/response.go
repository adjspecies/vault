// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package add

import (
	"github.com/adjspecies/vault/cmd/vault/commands/command"
	"github.com/adjspecies/vault/config"
)

// ResponseCommand allows adding a response to the vault.
type ResponseCommand command.BaseCommand

// Init initializes the command
func (cmd ResponseCommand) Init(cfg *config.Config, args []string) error {
	cmd.Config = cfg
	return nil
}

// Run runs the command
func (cmd ResponseCommand) Run() error {
	return nil
}

// NewResponseCommand returns a RegisteredCommand wrapping a ResponseCommand.
func NewResponseCommand() *command.RegisteredCommand {
	return &command.RegisteredCommand{
		Name:    "Add a response",
		Command: "add-response",
		Help:    ``,
		Entry:   &ResponseCommand{},
	}
}

// ResponsesCommand allows adding multiple responses to the vault.
type ResponsesCommand command.BaseCommand

// Init initializes the command
func (cmd ResponsesCommand) Init(cfg *config.Config, args []string) error {
	cmd.Config = cfg
	return nil
}

// Run runs the command
func (cmd ResponsesCommand) Run() error {
	return nil
}

// NewResponsesCommand returns a RegisteredCommand wrapping a ResponsesCommand.
func NewResponsesCommand() *command.RegisteredCommand {
	return &command.RegisteredCommand{
		Name:    "Add responses",
		Command: "add-responses",
		Help:    ``,
		Entry:   &ResponsesCommand{},
	}
}
