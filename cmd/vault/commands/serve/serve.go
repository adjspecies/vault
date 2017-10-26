// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package serve

import (
	"fmt"
	"net/http"

	"github.com/adjspecies/vault"
	"github.com/adjspecies/vault/cmd/vault/commands/command"
	"github.com/adjspecies/vault/config"
	"github.com/adjspecies/vault/logging"
	errgo "gopkg.in/errgo.v1"
)

type ServeCommand struct {
	cfg *config.Config
}

func (cmd *ServeCommand) Init(cfg *config.Config, args []string) error {
	cmd.cfg = cfg
	return nil
}

func (cmd *ServeCommand) Run() error {
	log := logging.Logger()
	handler, err := vault.NewServer()
	if err != nil {
		return errgo.Notef(err, "Could not create server")
	}
	log.Infow("starting server", "host", cmd.cfg.Host, "port", cmd.cfg.Port)
	return http.ListenAndServe(fmt.Sprintf("%s:%d", cmd.cfg.Host, cmd.cfg.Port), handler)
}

func NewServeCommand() *command.RegisteredCommand {
	return &command.RegisteredCommand{
		Name:    "start the Vault server",
		Command: "serve",
		Help:    ``,
		Entry:   &ServeCommand{},
	}
}
