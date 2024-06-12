package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	Server "github.com/mudriyjo/go-rest-template/src/go_rest_template/server"
)

func executeRequest(r *http.Request, s *Server.Server) *httptest.ResponseRecorder {
	response := httptest.NewRecorder()
	s.Router.ServeHTTP(response, r)

	return response
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestHelloWorld(t *testing.T) {
	// Create a New Server Struct
	s := Server.CreateNewServer()
	// Mount Handlers
	s.MountHandler()

	// Create a New Request
	req, _ := http.NewRequest("GET", "/", nil)

	// Execute Request
	response := executeRequest(req, s)

	// Check the response code
	checkResponseCode(t, http.StatusOK, response.Code)

	// We can use testify/require to assert values, as it is more convenient
	require.Equal(t, "Hello, world", response.Body.String())
}
