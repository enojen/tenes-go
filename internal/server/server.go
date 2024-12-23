package server

import (
	"log"
	"net/http"
)

type Server struct {
	router *http.ServeMux
}

func NewServer() *Server {
	return &Server{
		router: http.NewServeMux(),
	}
}

func (s *Server) Start(port string) error {
	s.routes()
	log.Printf("Server starting on %s", port)
	return http.ListenAndServe(port, s.router)
}

func (s *Server) routes() {
	s.router.HandleFunc("/health", s.handleHealth())
}

func (s *Server) handleHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}
}
