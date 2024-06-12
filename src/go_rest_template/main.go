package main

import (
	"net/http"

	Server "github.com/mudriyjo/go-rest-template/src/go_rest_template/server"
)

func main() {
	server := Server.CreateNewServer()
	server.MountHandler()
	http.ListenAndServe(":8080", server.Router)
}
