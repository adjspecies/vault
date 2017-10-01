// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package api

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/adjspecies/vault/api/v1"
)

func Register(r *mux.Router) error {
	r.HandleFunc("/", Status)
	v1.Register(r.PathPrefix("/api/v1").Subrouter())
	return nil
}

func Status(w http.ResponseWriter, r *http.Request) {

}
