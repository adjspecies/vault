package main

import (
    "path/filepath"
    "flag"
    "fmt"
    "os"

    "gopkg.in/errgo.v1"

    "github.com/adjspecies/vault"
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
    params := vault.Params{}
    _, err := vault.NewServer(params)
    if err != nil {
        return errgo.Notef(err, "Could not create server")
    }
    return nil
}
