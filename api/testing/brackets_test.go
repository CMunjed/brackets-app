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
	test_bracket = model.Bracket{
		Name:         "Test_Bracket",
		UserID:       "testuser",
		Size:         16,
		Type:         0,
		Public:       true,
		Edit:         false,
		AllowedUsers: []model.AllowedUser{},
		Teams:        test_bracket_teams,
	}
)

func setup() (*app.App, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	app := &app.App{}
	app.Initialize()
	return app, w
}

func getTestBracketID(t *testing.T, a *app.App, w *httptest.ResponseRecorder) string {
	r, err := http.NewRequest("GET", "/brackets", nil)
	if err != nil {
		t.Fatal(err)
	}

	a.Router.ServeHTTP(w, r)

	var response []model.Bracket
	body, err := io.ReadAll(w.Result().Body)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal(err)
	}

	return response[0].BracketID
}

func decodeBracket(w *httptest.ResponseRecorder, t *testing.T) model.Bracket {
	var response model.Bracket
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

func TestGetAllBrackets(t *testing.T) {
	r, err := http.NewRequest("GET", "/brackets", nil)
	if err != nil {
		t.Fatal(err)
	}

	app, w := setup()

	app.Router.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateBracket(t *testing.T) {
	app, w := setup()
	user := login(t, app, w)
	test_bracket.UserID = user.Email

	jsonData, err := json.Marshal(test_bracket)
	if err != nil {
		t.Fatal(err)
	}

	requestBody := bytes.NewBuffer(jsonData)
	r, err := http.NewRequest("POST", "/users/testuser/brackets", requestBody)
	if err != nil {
		t.Fatal(err)
	}

	c := w.Result().Cookies()
	w = httptest.NewRecorder()
	r.AddCookie(c[0])

	app.Router.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetBracket(t *testing.T) {

	app, w := setup()
	test_bracket_id := getTestBracketID(t, app, w)

	w = httptest.NewRecorder()
	user := login(t, app, w)
	address := "/users/" + user.Username + "/" + test_bracket_id
	r, err := http.NewRequest("GET", address, nil)
	if err != nil {
		t.Fatal(err)
	}

	c := w.Result().Cookies()
	w = httptest.NewRecorder()
	r.AddCookie(c[0])
	app.Router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	response := decodeBracket(w, t)

	assert.Equal(t, response.BracketID, test_bracket_id)
}

func TestUpdateBracket(t *testing.T) {
	type test_change struct {
		Type int
	}

	change := test_change{Type: 2}

	jsonData, err := json.Marshal(change)
	if err != nil {
		t.Fatal(err)
	}

	requestBody := bytes.NewBuffer(jsonData)

	app, w := setup()
	test_bracket_id := getTestBracketID(t, app, w)
	url := "/users/testuser/"
	url += test_bracket_id
	r, err := http.NewRequest("PUT", url, requestBody)
	if err != nil {
		t.Fatal(err)
	}

	w = httptest.NewRecorder()
	login(t, app, w)
	c := w.Result().Cookies()
	w = httptest.NewRecorder()
	r.AddCookie(c[0])
	app.Router.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

	response := decodeBracket(w, t)

	assert.Equal(t, 2, response.Type)
}

func TestDeleteBracket(t *testing.T) {
	app, w := setup()
	user := login(t, app, w)

	dummy_bracket := model.Bracket{
		Name:   "YOU CAN'T SEE THIS!!!",
		Size:   0,
		Type:   0,
		Public: false,
		Edit:   false,
		UserID: user.Email,

		Teams: []model.Team{},
	}

	jsonData, err := json.Marshal(dummy_bracket)
	if err != nil {
		t.Fatal(err)
	}

	requestBody := bytes.NewBuffer(jsonData)
	r, err := http.NewRequest("POST", "/users/testuser/brackets", requestBody)
	if err != nil {
		t.Fatal(err)
	}

	c := w.Result().Cookies()
	w = httptest.NewRecorder()
	r.AddCookie(c[0])

	app.Router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)

	response := decodeBracket(w, t)

	url := "/users/testuser/" + response.BracketID
	r, err = http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	w = httptest.NewRecorder()
	r.AddCookie(c[0])

	app.Router.ServeHTTP(w, r)

	database_response := decodeBracket(w, t)
	assert.Equal(t, response.BracketID, database_response.BracketID)

	r, err = http.NewRequest("DELETE", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	w = httptest.NewRecorder()
	r.AddCookie(c[0])
	app.Router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusNoContent, w.Code)

	r, err = http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	w = httptest.NewRecorder()
	r.AddCookie(c[0])
	app.Router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetUserBracket(t *testing.T) {
	r, err := http.NewRequest("GET", "/users/testuser/brackets", nil)
	if err != nil {
		t.Fatal(err)
	}

	app, w := setup()

	app.Router.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}
