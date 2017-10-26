// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package serve

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/adjspecies/vault"
	"github.com/adjspecies/vault/cmd/vault/commands/command"
	"github.com/adjspecies/vault/config"
	"github.com/adjspecies/vault/logging"
	errgo "gopkg.in/errgo.v1"
)

type ServeCommand struct{}

func (cmd *ServeCommand) Init(args []string) error {
	return nil
}

func (cmd *ServeCommand) Run() error {
	if err := serve(flag.Arg(0)); err != nil {
		return err
	}
	return nil
}

func NewServeCommand() *command.RegisteredCommand {
	return &command.RegisteredCommand{
		Name:    "Vault server",
		Command: "serve",
		Help:    ``,
		Entry:   &ServeCommand{},
	}
}

func serve(configPath string) error {
	conf, err := config.Read(configPath)
	if err != nil {
		return errgo.Notef(err, "cannot load configuration file")
	}
	if errSetup := logging.Setup(conf.LogLevel); errSetup != nil {
		return errgo.Notef(err, "unable to set up logging")
	}
	log := logging.Logger()
	defer log.Sync()
	handler, err := vault.NewServer()
	if err != nil {
		return errgo.Notef(err, "Could not create server")
	}
	log.Infow("starting server", "host", conf.Host, "port", conf.Port)
	return http.ListenAndServe(fmt.Sprintf("%s:%d", conf.Host, conf.Port), handler)
}
