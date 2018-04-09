package faultinfo

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/nasa9084/syg"
)

// Server represents faultinfo API server.
type Server struct {
	httpServer *http.Server
	closed     chan struct{}
}

// New returns a new faultinfo server.
func New(addr string) *Server {
	router := mux.NewRouter()
	bindroutes(router)
	s := &Server{
		httpServer: &http.Server{
			Addr:    addr,
			Handler: router,
		},
		closed: make(chan struct{}),
	}
	return s
}

// Run listen and serve.
func (s *Server) Run() error {
	cancel := syg.Listen(s.Shutdown, os.Interrupt)
	defer cancel()

	log.Printf("server is listening on: %s", s.httpServer.Addr)
	err := s.httpServer.ListenAndServe()
	<-s.closed
	return err
}

// Shutdown gracefully shuts down.
func (s *Server) Shutdown(os.Signal) {
	defer close(s.closed)
	log.Print("server shutdown")
	s.httpServer.Shutdown(context.Background())
}
