// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package vault

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/adjspecies/vault/api"
)

func NewServer() (http.Handler, error) {
	r := mux.NewRouter()
	if err := api.Register(r); err != nil {
		return nil, err
	}
	return r, nil
}
