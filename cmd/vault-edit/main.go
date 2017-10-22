// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/adjspecies/vault/logging"
	"go.uber.org/zap"
)

var log *zap.SugaredLogger

func main() {
	// Set up flags and parse them
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [options] <action> <input file>\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
		os.Exit(2)
	}
	var logLevel = flag.String("log-level", "info", "level of logging desired")
	flag.Parse()
	if flag.NArg() != 2 {
		flag.Usage()
	}

	// Set up logging.
	level, err := logging.LevelFromString(*logLevel)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to set log level: %s", *logLevel)
		os.Exit(3)
	}
	if errSetup := logging.Setup(level); errSetup != nil {
		fmt.Fprint(os.Stderr, "unable to set up logging")
		os.Exit(4)
	}
	log = logging.Logger()
	defer log.Sync()
	log.Info("logging set up")

	// Parse actions.
	action := flag.Arg(0)
	log.Debugf("command given: %s", action)
	switch action {
	case "import-survey":
		importSurvey(flag.Arg(1))
	case "import-source":
		importSource(flag.Arg(1))
	default:
		log.Fatalf("did not understand action %s", action)
	}
}

func importSurvey(filename string) {
	log.Debug("beginning import-survey")
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("could not read file %s", filename)
	}
	if contents == nil {
		log.Fatalf("file %s was empty", filename)
	}
}

func importSource(filename string) {
	log.Debug("beginning import-source")
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("could not read file %s", filename)
	}
	if contents == nil {
		log.Fatalf("file %s was empty", filename)
	}
}
