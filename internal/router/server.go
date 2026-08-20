package router

import (
	"net/http"
	"os"
)

type Server struct {
	httpServer *http.Server
}

func NewServer() *Server {
	return &Server{
		httpServer: &http.Server{
			Addr: os.Getenv("BACKEND_ADDR"),
			Handler: CreateRoutes(),
		},
	}
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}
