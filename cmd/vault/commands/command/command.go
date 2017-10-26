// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package command

type Command interface {
	Init([]string) error
	Run() error
}

type RegisteredCommand struct {
	Name    string
	Command string
	Help    string
	Entry   Command
}
