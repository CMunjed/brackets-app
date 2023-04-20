package testing

import (
	"bytes"
	"encoding/json"

	//"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	//"example.com/api/app/model"
	"github.com/go-playground/assert"
)

func TestWelcome(t *testing.T) {

	//Sign up to create user, sign in to create a user session, call welcome route
	//	to test the user session to see if it's working
	//Unsure how to check cookies using unit tests

	app, w := setup()

	//Convert test user to json
	jsonData, err := json.Marshal(test_user2)
	if err != nil {
		t.Fatal(err)
	}

	//requestBody := bytes.NewBuffer(jsonData)

	//Sign up - Not necessary, test user should already be in database
	/*r, err := http.NewRequest("POST", "/users/signup", requestBody)
	if err != nil {
		t.Fatal(err)
	}

	app.Router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusCreated, w.Code)*/

	//Sign in
	//w = httptest.NewRecorder()

	requestBody := bytes.NewBuffer(jsonData)
	r, err := http.NewRequest("POST", "/users/signup", requestBody)
	if err != nil {
		t.Fatal(err)
	}
	app.Router.ServeHTTP(w, r)

	w = httptest.NewRecorder()
	signedInUser := login2(t, app, w)

	//Get user (make sure user exists)
	w1 := httptest.NewRecorder()
	url := "/users/" + signedInUser.UserID

	r, err = http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	app.Router.ServeHTTP(w1, r)

	response := decodeUser(w1, t)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, test_user2.Username, response.Username)

	//Welcome
	w1 = httptest.NewRecorder()
	signInCookie := w.Result().Cookies()[0]
	requestBody = bytes.NewBuffer(jsonData)
	r, err = http.NewRequest("PUT", "/welcome", requestBody)
	if err != nil {
		t.Fatal(err)
	}
	r.AddCookie(signInCookie)
	app.Router.ServeHTTP(w1, r)

	//fmt.Println(w1.Body)

	assert.Equal(t, http.StatusOK, w1.Code)
}
