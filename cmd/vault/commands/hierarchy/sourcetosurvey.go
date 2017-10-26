// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package hierarchy

import (
	"github.com/adjspecies/vault/cmd/vault/commands/command"
	"github.com/adjspecies/vault/config"
)

// AddSurveyToSourceCommand controls adding one source to another in the
// hierarchy. This only adds a child to a parent; all other relationships are
// sussed out by Vault.
type AddSurveyToSourceCommand struct {
	cfg *config.Config
}

// Init initializes the command.
func (cmd AddSurveyToSourceCommand) Init(cfg *config.Config, args []string) error {
	cmd.cfg = cfg
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
		Name:    "Add a survey to source",
		Command: "add-survey-to-source",
		Help:    ``,
		Entry:   AddSurveyToSourceCommand{},
	}
}

// RemoveSurveyFromSourceCommand controls removing one source from another in
// the hierarchy. This only adds a child to a parent; all other relationships
// are sussed out by Vault.
type RemoveSurveyFromSourceCommand struct {
	cfg *config.Config
}

// Init initializes the command.
func (cmd RemoveSurveyFromSourceCommand) Init(cfg *config.Config, args []string) error {
	cmd.cfg = cfg
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
		Name:    "Remove a survey from a dource",
		Command: "remove-survey-from-source",
		Help:    ``,
		Entry:   RemoveSurveyFromSourceCommand{},
	}
}
