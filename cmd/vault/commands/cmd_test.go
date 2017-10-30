// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package commands_test

import (
	"fmt"
	"testing"

	"go.uber.org/zap/zapcore"

	"github.com/adjspecies/vault/cmd/vault/commands"
	"github.com/adjspecies/vault/cmd/vault/commands/command"
	"github.com/adjspecies/vault/config"
	"github.com/adjspecies/vault/logging"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(t *testing.T) {
	err := logging.Setup("development", zapcore.Level(-1))
	if err != nil {
		panic(err)
	}
	defaultConfig := &config.Config{}

	Convey("It should register commands", t, func() {
		// Call help as a default command which will succeed
		err := commands.Main(defaultConfig, "help", []string{})
		So(err, ShouldBeNil)

		// Check the command registry.
		So(len(commands.GetRegisteredCommands()), ShouldEqual, 7)
		So(commands.GetCommandList(), ShouldResemble, []string{
			"add-source",
			"add-survey",
			"add-source-to-source",
			"remove-source-from-source",
			"add-survey-to-source",
			"remove-survey-from-source",
			"serve",
		})
	})

	SkipConvey("It should run commands", t, func() {
		cmd := &TestCommand{}
		testCommand := &command.RegisteredCommand{
			Name:    "command for testing",
			Command: "test-command",
			Help:    "test command help",
			Entry:   cmd,
		}
		commands.RegisterCommand(testCommand)
		Convey("It should run without args", func() {
			commands.Main(defaultConfig, "test-command", []string{})
			So(cmd.InitCalled, ShouldEqual, 1)
			So(cmd.InitArgs, ShouldResemble, [][]string{})
			So(cmd.RunCalled, ShouldEqual, 1)
		})
		Convey("It should run with args", func() {
			commands.Main(defaultConfig, "test-command", []string{})
			So(cmd.InitCalled, ShouldEqual, 1)
			So(cmd.InitArgs, ShouldResemble, [][]string{})
			So(cmd.RunCalled, ShouldEqual, 1)
		})
	})
}

type TestCommand struct {
	command.BaseCommand
	InitCalled int
	InitArgs   [][]string
	InitError  string
	RunCalled  int
	RunError   string
}

func (cmd TestCommand) Init(cfg *config.Config, args []string) error {
	cmd.Config = cfg
	cmd.InitCalled++
	cmd.InitArgs = append(cmd.InitArgs, args)
	if cmd.InitError != "" {
		return fmt.Errorf(cmd.InitError)
	}
	return nil
}

func (cmd TestCommand) Run() error {
	cmd.RunCalled++
	if cmd.RunError != "" {
		return fmt.Errorf(cmd.InitError)
	}
	return nil
}
