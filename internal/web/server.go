package web

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
}

func NewServer() *Server {
	r := mux.NewRouter()
	RegisterRoutes(r)
	return &Server{
		Router: r,
	}
}

func (s *Server) Run(addr, port string) {
	if addr != "" {
		log.Printf("Server is running at %s:%s\n", addr, port)
		serverAddr := addr + ":" + port
		if err := http.ListenAndServe(serverAddr, s.Router); err != nil {
			log.Fatalf("Server Failed: %s", err)
		}
	}
}
