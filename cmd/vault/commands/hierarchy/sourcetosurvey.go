// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package hierarchy

import (
	"github.com/adjspecies/vault/cmd/vault/commands/command"
)

// AddSurveyToSourceCommand controls adding one source to another in the
// hierarchy. This only adds a child to a parent; all other relationships are
// sussed out by Vault.
type AddSurveyToSourceCommand struct{}

// Init initializes the command.
func (cmd AddSurveyToSourceCommand) Init(args []string) error {
	return nil
}

// Run runs the command.
func (cmd AddSurveyToSourceCommand) Run() error {
	return nil
}

// NewAddSurveyToSourceCommand returns a RegisteredCommand containing all of
// the info required by the subcommand system.
func NewAddSurveyToSourceCommand() *command.RegisteredCommand {
	return &command.RegisteredCommand{
		Name:    "Add source to source",
		Command: "add-source-to-source",
		Help:    ``,
		Entry:   AddSurveyToSourceCommand{},
	}
}

// RemoveSurveyFromSourceCommand controls removing one source from another in
// the hierarchy. This only adds a child to a parent; all other relationships
// are sussed out by Vault.
type RemoveSurveyFromSourceCommand struct{}

// Init initializes the command.
func (cmd RemoveSurveyFromSourceCommand) Init(args []string) error {
	return nil
}

// Run runs the command.
func (cmd RemoveSurveyFromSourceCommand) Run() error {
	return nil
}

// NewRemoveSurveyFromSourceCommand returns a RegisteredCommand containing all
// of the info required by the subcommand system.
func NewRemoveSurveyFromSourceCommand() *command.RegisteredCommand {
	return &command.RegisteredCommand{
		Name:    "Add source to source",
		Command: "add-source-to-source",
		Help:    ``,
		Entry:   RemoveSurveyFromSourceCommand{},
	}
}
