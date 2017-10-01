// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"gopkg.in/errgo.v1"

	"github.com/adjspecies/vault"
	"github.com/adjspecies/vault/config"
	"github.com/adjspecies/vault/logging"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s <config path>\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
		os.Exit(2)
	}
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
	}
	if err := serve(flag.Arg(0)); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func serve(configPath string) error {
	conf, err := config.Read(configPath)
	if err != nil {
		return errgo.Notef(err, "cannot load configuration file")
	}
	if err := logging.Setup(conf.LogLevel); err != nil {
		return errgo.Notef(err, "unable to set up logging")
	}
	log := logging.Logger()
	defer log.Sync()
	handler, err := vault.NewServer()
	if err != nil {
		return errgo.Notef(err, "Could not create server")
	}
	return http.ListenAndServe(fmt.Sprintf("%s:%d", conf.Host, conf.Port), handler)
}
