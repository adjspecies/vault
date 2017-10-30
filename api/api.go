// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/adjspecies/vault/api/v1"
	"github.com/adjspecies/vault/logging"
)

var log *zap.SugaredLogger

// Register adds handlers for the various API versions.
func Register(r *mux.Router) error {
	log = logging.Logger()
	r.HandleFunc("/", Status)
	log.Debug("registering v1 endpoints")
	v1.Register(r.PathPrefix("/api/v1").Subrouter())
	return nil
}

// Status shows the status of the service and all of its connections.
func Status(w http.ResponseWriter, r *http.Request) {

}
