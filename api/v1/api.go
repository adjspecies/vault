package v1

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Register(r *mux.Router) error {
	r.HandleFunc("/", Overview)
	return nil
}

func Overview(w http.ResponseWriter, r *http.Request) {

}
