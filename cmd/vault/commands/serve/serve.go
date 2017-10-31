// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package serve

import (
	"fmt"
	"net/http"

	"github.com/adjspecies/vault/cmd/vault/commands/command"
	"github.com/adjspecies/vault/config"
	"github.com/adjspecies/vault/logging"
	"github.com/adjspecies/vault/server"
	errgo "gopkg.in/errgo.v1"
)

type ServeCommand command.BaseCommand

// Init initializes the command
func (cmd *ServeCommand) Init(cfg *config.Config, args []string) error {
	cmd.Config = cfg
	return nil
}

// Run starts the server.
func (cmd *ServeCommand) Run() error {
	log := logging.Logger()
	handler, err := server.NewServer()
	if err != nil {
		return errgo.Notef(err, "Could not create server")
	}
	log.Infow("starting server", "host", cmd.Config.Host, "port", cmd.Config.Port)
	return http.ListenAndServe(fmt.Sprintf("%s:%d", cmd.Config.Host, cmd.Config.Port), handler)
}

// NewServeCommand creates a new registered command for starting the server.
func NewServeCommand() *command.RegisteredCommand {
	return &command.RegisteredCommand{
		Name:    "start the Vault server",
		Command: "serve",
		Help:    ``,
		Entry:   &ServeCommand{},
	}
}
