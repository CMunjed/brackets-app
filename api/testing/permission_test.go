package testing

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/api/app/model"
	"github.com/go-playground/assert"
)

func TestUnauthenticatedUser(t *testing.T) {
	// Setup the server
	app, w := setup()

	// Create a new user
	new_user := model.User{
		Email:    "newuser@hotmail.com",
		Username: "newuser",
		Password: "newuser",
	}

	jsonBody, err := json.Marshal(new_user)
	if err != nil {
		t.Fatal(err)
	}

	requestBody := bytes.NewBuffer(jsonBody)

	r, err := http.NewRequest("POST", "/users/signup", requestBody)
	if err != nil {
		t.Fatal(err)
	}

	app.Router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusCreated, w.Code)

	//Login with new user to get token
	w = httptest.NewRecorder()

	jsonBody, err = json.Marshal(new_user)
	if err != nil {
		t.Fatal(err)
	}

	requestBody = bytes.NewBuffer(jsonBody)

	r, err = http.NewRequest("PUT", "/users/signin", requestBody)
	if err != nil {
		t.Fatal(err)
	}

	app.Router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)

	// Create a bracket linked to the user
	response := decodeUser(w, t)

	new_bracket := model.Bracket{
		Name:         "New Bracket",
		UserID:       response.UserID,
		Size:         4,
		Type:         0,
		Public:       false,
		Edit:         false,
		AllowedUsers: []model.AllowedUser{},
		Teams:        []model.Team{},
	}

	jsonBody, err = json.Marshal(new_bracket)
	if err != nil {
		t.Fatal(err)
	}

	requestBody = bytes.NewBuffer(jsonBody)

	r, err = http.NewRequest("POST", "/users/"+response.UserID+"/brackets", requestBody)
	if err != nil {
		t.Fatal(err)
	}

	c := w.Result().Cookies()
	w = httptest.NewRecorder()
	r.AddCookie(c[0])

	app.Router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)

	// Get the bracket ID
	bracketResponse := decodeBracket(w, t)

	// Sign up the unautorized user
	w = httptest.NewRecorder()
	unauthorizedUser := model.User{
		Email:    "badguy@villian.com",
		Username: "badguy",
		Password: "badguy",
	}

	jsonBody, err = json.Marshal(unauthorizedUser)
	if err != nil {
		t.Fatal(err)
	}

	requestBody = bytes.NewBuffer(jsonBody)

	r, err = http.NewRequest("POST", "/users/signup", requestBody)
	if err != nil {
		t.Fatal(err)
	}

	app.Router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusCreated, w.Code)

	//Login with unauthorized user to get token
	w = httptest.NewRecorder()

	jsonBody, err = json.Marshal(unauthorizedUser)
	if err != nil {
		t.Fatal(err)
	}

	requestBody = bytes.NewBuffer(jsonBody)

	r, err = http.NewRequest("PUT", "/users/signin", requestBody)
	if err != nil {
		t.Fatal(err)
	}

	app.Router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)

	// Attempt to get the bracket
	r, err = http.NewRequest("GET", "/users/testuser/brackets/"+bracketResponse.BracketID, nil)
	if err != nil {
		t.Fatal(err)
	}

	c = w.Result().Cookies()
	w = httptest.NewRecorder()
	r.AddCookie(c[0])
	app.Router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
