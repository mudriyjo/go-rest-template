package main

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

func (s *Server) mountHandler() {
	s.Router.Use(middleware.Logger)
	s.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world"))
	})
}
func main() {
	server := CreateNewServer()
	server.mountHandler()
	http.ListenAndServe(":8080", server.Router)
}
