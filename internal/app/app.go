package App

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mudriyjo/go-rest-template/internal/config"
)

type Server struct {
	Config config.Config
	Router *chi.Mux
}

func CreateNewServer() *Server {
	return &Server{
		Router: chi.NewRouter(),
		Config: config.GetConfig(),
	}
}

func (s *Server) MountHandler() {
	fmt.Println(s.Config)
	s.Router.Use(middleware.Logger)
	s.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world"))
	})
}
