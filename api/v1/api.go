// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package v1

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Register adds handlers for endpoints in the v1 API.
func Register(r *mux.Router) error {
	r.HandleFunc("/", Overview)
	return nil
}

// Overview provides a brief glimpse at the data available in the vault.
func Overview(w http.ResponseWriter, r *http.Request) {

}
