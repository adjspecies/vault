// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package command

import (
	"github.com/adjspecies/vault/config"
)

type Command interface {
	Init(*config.Config, []string) error
	Run() error
}

type RegisteredCommand struct {
	Name    string
	Command string
	Help    string
	Entry   Command
}
