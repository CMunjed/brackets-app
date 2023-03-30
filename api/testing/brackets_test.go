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
	Test_bracket_id = "59d8ef98-b83b-4a49-a069-cbd686d89736"

	test_bracket = model.Bracket{
		Name:   "Test_Bracket",
		UserID: "testuser",
		Size:   16,
		Type:   0,
		Teams:  test_bracket_teams,
	}
)

func setup() (*app.App, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	app := &app.App{}
	app.Initialize()
	return app, w
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
	jsonData, err := json.Marshal(test_bracket)
	if err != nil {
		t.Fatal(err)
	}

	requestBody := bytes.NewBuffer(jsonData)
	r, err := http.NewRequest("POST", "/users/testuser/brackets", requestBody)
	if err != nil {
		t.Fatal(err)
	}

	app, w := setup()

	app.Router.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetBracket(t *testing.T) {
	address := "/users/testuser/" + Test_bracket_id
	r, err := http.NewRequest("GET", address, nil)
	if err != nil {
		t.Fatal(err)
	}

	app, w := setup()

	app.Router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	response := decodeBracket(w, t)

	assert.Equal(t, response.BracketID, Test_bracket_id)
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
	url := "/users/testuser/"
	url += Test_bracket_id
	r, err := http.NewRequest("PUT", url, requestBody)
	if err != nil {
		t.Fatal(err)
	}

	app, w := setup()

	app.Router.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

	response := decodeBracket(w, t)

	assert.Equal(t, 2, response.Type)
}

func TestDeleteBracket(t *testing.T) {
	dummy_bracket := model.Bracket{
		Name:   "YOU CAN'T SEE THIS!!!",
		Size:   0,
		Type:   0,
		UserID: "testuser",
		Teams:  []model.Team{},
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

	app, w := setup()

	app.Router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)

	response := decodeBracket(w, t)

	url := "/users/testuser/" + response.BracketID
	r, err = http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	w = httptest.NewRecorder()

	app.Router.ServeHTTP(w, r)

	database_response := decodeBracket(w, t)
	assert.Equal(t, response.BracketID, database_response.BracketID)

	r, err = http.NewRequest("DELETE", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	w = httptest.NewRecorder()
	app.Router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusNoContent, w.Code)

	r, err = http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	w = httptest.NewRecorder()
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
