package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"../databases"
	"../models"
	"../services"
)

var (
	characterSt   databases.CharacterStorage = databases.NewSQLiteStorage()
	characterSrv  services.CharacterService  = services.NewCharacterService(characterSt)
	characterCtrl CharacterController        = NewCharacterController(characterSrv)
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/characters", characterCtrl.GetCharacters).Methods("GET")
	router.HandleFunc("/api/characters", characterCtrl.CreateCharacter).Methods("POST")
	return router
}

func TestGetCharacters(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/characters", nil)
	response := httptest.NewRecorder()

	Router().ServeHTTP(response, request)
	t.Logf("!---> RESPONSE: %v", response)
	assert.Equal(t, 200, response.Code, "The expected status code must be 200.")
}

func TestCreateCharacter(t *testing.T) {
	character := models.Character{
		FirstName: "John",
		LastName:  "Doe",
		Nickname:  "johnny",
	}
	jsonCharacter, _ := json.Marshal(&character)

	request, _ := http.NewRequest("POST", "/api/characters", bytes.NewBuffer(jsonCharacter))
	response := httptest.NewRecorder()

	Router().ServeHTTP(response, request)
	assert.Equal(t, 201, response.Code, "The expected status code must be 201.")
}
