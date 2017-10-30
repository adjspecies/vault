// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package commands

import (
	"github.com/adjspecies/vault/cmd/vault/commands/command"
)

func GetCommandList() []string {
	return commandList
}

func GetRegisteredCommands() map[string]*command.RegisteredCommand {
	return registeredCommands
}

func RegisterCommand(cmd *command.RegisteredCommand) {
	registerCommand(cmd)
}

func RegisterCommands() {
	registerCommands()
}
