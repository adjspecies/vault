// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

// Package server manages creating the HTTP server for Vault.
package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/adjspecies/vault/api"
)

// NewServer generates a new http.Handler to serve Vault. It constructs the
// Gorilla Mux router by letting the API package register routes that it knows
// about.
func NewServer() (http.Handler, error) {
	r := mux.NewRouter()
	if err := api.Register(r); err != nil {
		return nil, err
	}
	return r, nil
}
