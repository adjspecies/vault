// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package v1

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/adjspecies/vault/logging"
)

var log *zap.SugaredLogger

// Register adds handlers for endpoints in the v1 API.
func Register(r *mux.Router) error {
	log = logging.Logger()
	log.Debug("registering endpoint overview")
	r.HandleFunc("/", Overview)
	return nil
}

// Overview provides a brief glimpse at the data available in the vault.
func Overview(w http.ResponseWriter, r *http.Request) {

}
