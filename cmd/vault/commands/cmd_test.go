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
	logging.ObserveLogging(zapcore.DebugLevel)
	defaultConfig := &config.Config{}

	Convey("It should register commands", t, func() {
		// Call help as a default command which will succeed
		err := commands.Main(defaultConfig, "help", []string{})
		So(err, ShouldBeNil)

		// Check the command registry.
		So(len(commands.GetRegisteredCommands()), ShouldEqual, 13)
		So(commands.GetCommandList(), ShouldResemble, []string{
			"add-source",
			"add-summary",
			"add-survey",
			"add-response",
			"add-responses",
			"add-respondent",
			"add-source-to-source",
			"remove-source-from-source",
			"add-summary-to-source",
			"remove-summary-from-source",
			"add-survey-to-source",
			"remove-survey-from-source",
			"serve",
		})
	})

	Convey("It should run commands", t, func() {
		Convey("It should run without args", func() {
			cmd := createCommand()
			err := commands.Main(defaultConfig, "test-command", []string{})
			So(err, ShouldBeNil)
			So(cmd.InitCalled, ShouldEqual, 1)
			So(cmd.InitArgs, ShouldResemble, [][]string{{}})
			So(cmd.RunCalled, ShouldEqual, 1)
		})
		Convey("It should run with args", func() {
			cmd := createCommand()
			args := []string{"the", "doctor"}
			err := commands.Main(defaultConfig, "test-command", args)
			So(err, ShouldBeNil)
			So(cmd.InitCalled, ShouldEqual, 1)
			So(cmd.InitArgs[0], ShouldResemble, args)
			So(cmd.RunCalled, ShouldEqual, 1)
		})
		Convey("It should return an error from Run", func() {
			cmd := createCommand()
			cmd.RunError = "bad-wolf"
			err := commands.Main(defaultConfig, "test-command", []string{})
			So(err, ShouldBeError)
			So(err.Error(), ShouldEqual, "Run: bad-wolf")
		})
		Convey("It should return an error from Init", func() {
			cmd := createCommand()
			cmd.InitError = "bad-wolf"
			err := commands.Main(defaultConfig, "test-command", []string{})
			So(err, ShouldBeError)
			So(err.Error(), ShouldEqual, "Init: bad-wolf")
		})
	})
}

func TestHelp(t *testing.T) {
	createCommand()
	Convey("It should display help for a test command", t, func() {
		logging.ObserveLogging(zapcore.DebugLevel)
		err := commands.Help([]string{"test-command"})
		So(err, ShouldBeNil)
		So(logging.ObserverLogs().All()[0].Message, ShouldEqual, "listing help for test-command")
	})
	Convey("It should error if the command doesn't exist", t, func() {
		logging.ObserveLogging(zapcore.DebugLevel)
		err := commands.Help([]string{"bad-wolf"})
		So(err, ShouldBeError)
		So(err.Error(), ShouldEqual, "could not find command bad-wolf")
		So(logging.ObserverLogs().All()[0].Message, ShouldEqual, "could not find command bad-wolf")
	})
	Convey("It should list commands and their names", t, func() {
		logging.ObserveLogging(zapcore.DebugLevel)
		err := commands.Help([]string{"commands"})
		So(err, ShouldBeNil)
		So(logging.ObserverLogs().All()[0].Message, ShouldEqual, "listing commands")
	})
	Convey("It should display the master help", t, func() {
		logging.ObserveLogging(zapcore.DebugLevel)
		err := commands.Help([]string{})
		So(err, ShouldBeNil)
		So(logging.ObserverLogs().All()[0].Message, ShouldEqual, "printing masterHelp")
	})
}

func createCommand() *TestCommand {
	cmd := &TestCommand{}
	testCommand := &command.RegisteredCommand{
		Name:    "command for testing",
		Command: "test-command",
		Help:    "test command help",
		Entry:   cmd,
	}
	commands.RegisterCommand(testCommand)
	return cmd
}

type TestCommand struct {
	command.BaseCommand
	InitCalled int
	InitArgs   [][]string
	InitError  string
	RunCalled  int
	RunError   string
}

func (cmd *TestCommand) Init(cfg *config.Config, args []string) error {
	cmd.Config = cfg
	cmd.InitCalled++
	cmd.InitArgs = append(cmd.InitArgs, args)
	if cmd.InitError != "" {
		return fmt.Errorf("Init: %s", cmd.InitError)
	}
	return nil
}

func (cmd *TestCommand) Run() error {
	cmd.RunCalled++
	if cmd.RunError != "" {
		return fmt.Errorf("Run: %s", cmd.RunError)
	}
	return nil
}
