package web

import (
	"net/http"

	"github.com/JoelInman-Dev/tickwars/pkg/utils"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	v1Routes := r.PathPrefix("/v1").Subrouter()

	v1Routes.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		utils.PlayerLogger.Info("Hello, first route working", "Yee", "Haw!")
	})
}
