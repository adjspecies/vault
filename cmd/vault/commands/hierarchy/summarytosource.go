// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package hierarchy

import (
	"github.com/adjspecies/vault/cmd/vault/commands/command"
	"github.com/adjspecies/vault/config"
)

// AddSummaryToSourceCommand controls adding one source to another in the
// hierarchy. This only adds a child to a parent; all other relationships are
// sussed out by Vault.
type AddSummaryToSourceCommand command.BaseCommand

// Init initializes the command.
func (cmd AddSummaryToSourceCommand) Init(cfg *config.Config, args []string) error {
	cmd.Config = cfg
	return nil
}

// Run runs the command.
func (cmd AddSummaryToSourceCommand) Run() error {
	return nil
}

// NewAddSummaryToSourceCommand returns a RegisteredCommand containing all of
// the info required by the subcommand system.
func NewAddSummaryToSourceCommand() *command.RegisteredCommand {
	return &command.RegisteredCommand{
		Name:    "Add summary to a source",
		Command: "add-summary-to-source",
		Help:    ``,
		Entry:   &AddSummaryToSourceCommand{},
	}
}

// RemoveSummaryFromSourceCommand controls removing one source from another in
// the hierarchy. This only adds a child to a parent; all other relationships
// are sussed out by Vault.
type RemoveSummaryFromSourceCommand command.BaseCommand

// Init initializes the command.
func (cmd RemoveSummaryFromSourceCommand) Init(cfg *config.Config, args []string) error {
	cmd.Config = cfg
	return nil
}

// Run runs the command.
func (cmd RemoveSummaryFromSourceCommand) Run() error {
	return nil
}

// NewRemoveSummaryFromSourceCommand returns a RegisteredCommand containing all
// of the info required by the subcommand system.
func NewRemoveSummaryFromSourceCommand() *command.RegisteredCommand {
	return &command.RegisteredCommand{
		Name:    "Remove a summary from a source",
		Command: "remove-summary-from-source",
		Help:    ``,
		Entry:   &RemoveSummaryFromSourceCommand{},
	}
}
