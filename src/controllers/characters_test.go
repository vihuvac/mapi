package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/characters", GetCharacters).Methods("GET")
	return router
}

func TestGetCharacters(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/characters", nil)
	response := httptest.NewRecorder()

	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "The expected status code must be 200.")
}
