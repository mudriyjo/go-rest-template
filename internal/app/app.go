package App

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	Router *chi.Mux
}

func CreateNewServer() *Server {
	return &Server{
		chi.NewRouter(),
	}
}

func (s *Server) MountHandler() {
	s.Router.Use(middleware.Logger)
	s.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world"))
	})
}
