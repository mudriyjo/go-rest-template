package main

import (
	"net/http"

	App "github.com/mudriyjo/go-rest-template/internal/app"
)

func main() {
	server := App.CreateNewServer()
	server.MountHandler()
	http.ListenAndServe(":8080", server.Router)
}
