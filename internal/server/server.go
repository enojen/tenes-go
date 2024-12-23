package server

import (
	"log"
	"net/http"
	"time"
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

	srv := &http.Server{
		Addr:         port,
		Handler:      s.router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	return srv.ListenAndServe()
}

func (s *Server) routes() {
	s.router.HandleFunc("/health", s.handleHealth())
}

func (s *Server) handleHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte("OK")); err != nil {
			log.Printf("Error writing response: %v", err)
		}
	}
}
