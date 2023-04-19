package testing

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/api/app"
	"example.com/api/app/model"
	"github.com/go-playground/assert"
)

var (
	test_user = model.User{
		Email:    "testemail@test.com",
		Username: "testuser",
		Password: "testpassword",
		Admin:    true,
	}

	dummy_user = model.User{
		Email:    "cantsee@me.net",
		Username: "cantseeme",
		Password: "cantseeme",
		Admin:    false,
	}

	googlesignin_user = model.GoogleUser{
		Token: model.Token{
			Iss:    "test",
			Nbf:    123456789,
			Aud:    "test",
			Id:     "test",
			Email:  "test@email.com",
			Emailv: true,
			Azp:    "test",
			Name:   "test",
			Imgurl: "test",
			Gname:  "test",
			Iat:    123456789,
			Exp:    123456789,
			Jti:    "test",
		},
	}
)

func decodeUser(w *httptest.ResponseRecorder, t *testing.T) model.User {
	var response model.User
	body, err := io.ReadAll(w.Result().Body)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal(err)
	}

	return response
}

func login(t *testing.T, a *app.App, w *httptest.ResponseRecorder) model.User {
	jsonData, err := json.Marshal(test_user)
	if err != nil {
		t.Fatal(err)
	}

	requestBody := bytes.NewBuffer(jsonData)

	r, err := http.NewRequest("PUT", "/users/signin", requestBody)
	if err != nil {
		t.Fatal(err)
	}
	a.Router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)

	return decodeUser(w, t)
}

func TestSignUp(t *testing.T) {
	app, w := setup()

	jsonData, err := json.Marshal(test_user)
	if err != nil {
		t.Fatal(err)
	}

	requestBody := bytes.NewBuffer(jsonData)

	r, err := http.NewRequest("POST", "/users/signup", requestBody)
	if err != nil {
		t.Fatal(err)
	}
	app.Router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestSignIn(t *testing.T) {
	app, w := setup()

	response := login(t, app, w)
	assert.Equal(t, test_user.Username, response.Username)
}

func TestGetUser(t *testing.T) {
	app, w := setup()

	user := login(t, app, w)

	url := "/users/" + user.UserID

	w = httptest.NewRecorder()

	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	app.Router.ServeHTTP(w, r)

	response := decodeUser(w, t)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, user.UserID, response.UserID)
}

func TestUpdatePassword(t *testing.T) {
	app, w := setup()

	user := login(t, app, w)

	w = httptest.NewRecorder()
	type update struct {
		Password string `json:"password"`
	}
	user_update := update{
		Password: "newpassword",
	}

	jsonData, err := json.Marshal(user_update)
	if err != nil {
		t.Fatal(err)
	}

	requestBody := bytes.NewBuffer(jsonData)

	url := "/users/" + user.UserID

	r, err := http.NewRequest("PUT", url, requestBody)
	if err != nil {
		t.Fatal(err)
	}
	app.Router.ServeHTTP(w, r)

	response := decodeUser(w, t)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, user.UserID, response.UserID)

	url = "/users/signin"

	test_user.Password = user_update.Password
	test_user.UserID = user.UserID

	jsonData, err = json.Marshal(test_user)
	if err != nil {
		t.Fatal(err)
	}

	requestBody = bytes.NewBuffer(jsonData)

	r, err = http.NewRequest("PUT", url, requestBody)
	if err != nil {
		t.Fatal(err)
	}

	w = httptest.NewRecorder()
	app.Router.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

	//Change back
	w = httptest.NewRecorder()
	user_update = update{
		Password: "testpassword",
	}

	jsonData, err = json.Marshal(user_update)
	if err != nil {
		t.Fatal(err)
	}

	requestBody = bytes.NewBuffer(jsonData)

	url = "/users/" + user.UserID

	r, err = http.NewRequest("PUT", url, requestBody)
	if err != nil {
		t.Fatal(err)
	}
	app.Router.ServeHTTP(w, r)

	response = decodeUser(w, t)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, user.UserID, response.UserID)

	url = "/users/signin"

	test_user.Password = user_update.Password
	test_user.UserID = user.UserID

	jsonData, err = json.Marshal(test_user)
	if err != nil {
		t.Fatal(err)
	}

	requestBody = bytes.NewBuffer(jsonData)

	r, err = http.NewRequest("PUT", url, requestBody)
	if err != nil {
		t.Fatal(err)
	}

	w = httptest.NewRecorder()
	app.Router.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateEmail(t *testing.T) {
	app, w := setup()

	user := login(t, app, w)
	c := w.Result().Cookies()
	w = httptest.NewRecorder()
	type update struct {
		Email string `json:"email"`
	}
	user_update := update{
		Email: "newemail@example.com",
	}

	jsonData, err := json.Marshal(user_update)
	if err != nil {
		t.Fatal(err)
	}

	requestBody := bytes.NewBuffer(jsonData)
	url := "/users/" + user.UserID

	r, err := http.NewRequest("PUT", url, requestBody)
	if err != nil {
		t.Fatal(err)
	}
	r.AddCookie(c[0])
	app.Router.ServeHTTP(w, r)

	response := decodeUser(w, t)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, user.UserID, response.UserID)
	assert.Equal(t, user_update.Email, response.Email)

	//Change back
	w = httptest.NewRecorder()

	user_update = update{
		Email: "testemail@test.com",
	}

	jsonData, err = json.Marshal(user_update)
	if err != nil {
		t.Fatal(err)
	}

	requestBody = bytes.NewBuffer(jsonData)
	url = "/users/" + user.UserID

	r, err = http.NewRequest("PUT", url, requestBody)
	if err != nil {
		t.Fatal(err)
	}
	r.AddCookie(c[0])
	app.Router.ServeHTTP(w, r)

	response = decodeUser(w, t)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, user.UserID, response.UserID)
	assert.Equal(t, user_update.Email, response.Email)
	//test_user.Email = "newemail@example.com"
}

func TestDeleteUser(t *testing.T) {
	app, w := setup()

	//Implace the dummy data first
	jsonData, err := json.Marshal(dummy_user)
	if err != nil {
		t.Fatal(err)
	}

	requestBody := bytes.NewBuffer(jsonData)

	r, err := http.NewRequest("POST", "/users/signup", requestBody)
	if err != nil {
		t.Fatal(err)
	}
	app.Router.ServeHTTP(w, r)

	assert.Equal(t, http.StatusCreated, w.Code)

	//Delete the dummy user
	dummy_user = decodeUser(w, t)

	w = httptest.NewRecorder()
	//login(t, app, w)
	//c := w.Result().Cookies()
	//w = httptest.NewRecorder()
	url := "/users/" + dummy_user.UserID

	r, err = http.NewRequest("DELETE", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	//r.AddCookie(c[0])
	app.Router.ServeHTTP(w, r)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestGetAllUsers(t *testing.T) {
	app, w := setup()
	//login(t, app, w)
	//c := w.Result().Cookies()
	//w = httptest.NewRecorder()

	r, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	//r.AddCookie(c[0])
	app.Router.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGoogleSignup(t *testing.T) {
	app, w := setup()

	jsonData, err := json.Marshal(googlesignin_user)
	if err != nil {
		t.Fatal(err)
	}

	requestBody := bytes.NewBuffer(jsonData)

	r, err := http.NewRequest("POST", "/users/googlesignup", requestBody)
	if err != nil {
		t.Fatal(err)
	}
	app.Router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGoogleSignIn(t *testing.T) {
	app, w := setup()

	jsonData, err := json.Marshal(googlesignin_user)
	if err != nil {
		t.Fatal(err)
	}

	requestBody := bytes.NewBuffer(jsonData)

	r, err := http.NewRequest("PUT", "/users/googlesignin", requestBody)
	if err != nil {
		t.Fatal(err)
	}
	app.Router.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}
