// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package hierarchy

import (
	"github.com/adjspecies/vault/cmd/vault/commands/command"
)

// AddSourceToSourceCommand controls adding one source to another in the
// hierarchy. This only adds a child to a parent; all other relationships are
// sussed out by Vault.
type AddSourceToSourceCommand struct{}

// Init initializes the command.
func (cmd AddSourceToSourceCommand) Init(args []string) error {
	return nil
}

// Run runs the command.
func (cmd AddSourceToSourceCommand) Run() error {
	return nil
}

// NewAddSourceToSourceCommand returns a RegisteredCommand containing all of
// the info required by the subcommand system.
func NewAddSourceToSourceCommand() *command.RegisteredCommand {
	return &command.RegisteredCommand{
		Name:    "Add source to source",
		Command: "add-source-to-source",
		Help:    ``,
		Entry:   AddSourceToSourceCommand{},
	}
}

// RemoveSourceFromSourceCommand controls removing one source from another in
// the hierarchy. This only adds a child to a parent; all other relationships
// are sussed out by Vault.
type RemoveSourceFromSourceCommand struct{}

// Init initializes the command.
func (cmd RemoveSourceFromSourceCommand) Init(args []string) error {
	return nil
}

// Run runs the command.
func (cmd RemoveSourceFromSourceCommand) Run() error {
	return nil
}

// NewRemoveSourceFromSourceCommand returns a RegisteredCommand containing all
// of the info required by the subcommand system.
func NewRemoveSourceFromSourceCommand() *command.RegisteredCommand {
	return &command.RegisteredCommand{
		Name:    "Add source to source",
		Command: "add-source-to-source",
		Help:    ``,
		Entry:   RemoveSourceFromSourceCommand{},
	}
}
