// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package main

import (
	"os"

	"github.com/adjspecies/vault/cmd/vault/commands"
	"go.uber.org/zap"
)

var log *zap.SugaredLogger

func main() {
	var command string
	if len(os.Args) < 2 {
		command = "help"
	} else {
		command = os.Args[0]
	}

	commands.Main(command, os.Args[2:])
}
