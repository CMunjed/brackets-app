package testing

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/api/app/model"
	"github.com/go-playground/assert"
)

var (
	test_bracket_teams = []model.Team{
		{
			Name:       "Team1",
			Index:      0,
			Round:      0,
			Position:   1,
			Eliminated: false,
		},
		{
			Name:       "Team2",
			Index:      1,
			Round:      0,
			Position:   2,
			Eliminated: false,
		},
		{
			Name:       "Team3",
			Index:      2,
			Round:      0,
			Position:   3,
			Eliminated: false,
		},
		{
			Name:       "Team4",
			Index:      3,
			Round:      0,
			Position:   4,
			Eliminated: false,
		},
		{
			Name:       "Team5",
			Index:      4,
			Round:      0,
			Position:   5,
			Eliminated: false,
		},
		{
			Name:       "Team6",
			Index:      5,
			Round:      0,
			Position:   6,
			Eliminated: false,
		},
		{
			Name:       "Team7",
			Index:      6,
			Round:      0,
			Position:   7,
			Eliminated: false,
		},
	}
)

func decodeTeam(w *httptest.ResponseRecorder, t *testing.T) model.Team {
	var response model.Team
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

func TestGetAllTeams(t *testing.T) {
	app, w := setup()

	url := "/users/testuser/" + getTestBracketID(t, app, w) + "/teams"
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	w = httptest.NewRecorder()
	app.Router.ServeHTTP(w, r)

	assert.Equal(t, w.Code, http.StatusOK)

	var response []model.Team
	body, err := io.ReadAll(w.Result().Body)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(test_bracket_teams); i++ {
		assert.Equal(t, response[i].Name, test_bracket_teams[i].Name)
		assert.Equal(t, response[i].Index, test_bracket_teams[i].Index)
		assert.Equal(t, response[i].Round, test_bracket_teams[i].Round)
		assert.Equal(t, response[i].Position, test_bracket_teams[i].Position)
		assert.Equal(t, response[i].Eliminated, test_bracket_teams[i].Eliminated)
	}
}

func TestGetTeam(t *testing.T) {
	app, w := setup()

	url := "/users/testuser/" + getTestBracketID(t, app, w) + "/teams/0"
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	w = httptest.NewRecorder()
	app.Router.ServeHTTP(w, r)

	assert.Equal(t, w.Code, http.StatusOK)

	response := decodeTeam(w, t)

	assert.Equal(t, response.Name, test_bracket_teams[0].Name)
	assert.Equal(t, response.Index, test_bracket_teams[0].Index)
	assert.Equal(t, response.Round, test_bracket_teams[0].Round)
	assert.Equal(t, response.Position, test_bracket_teams[0].Position)
	assert.Equal(t, response.Eliminated, test_bracket_teams[0].Eliminated)
}

func TestCreateTeam(t *testing.T) {
	app, w := setup()

	newTeam := model.Team{
		Name: "Team8",
	}

	jsonData, err := json.Marshal(newTeam)
	if err != nil {
		t.Fatal(err)
	}

	requestBody := bytes.NewBuffer(jsonData)

	url := "/users/testuser/" + getTestBracketID(t, app, w) + "/teams"
	r, err := http.NewRequest("POST", url, requestBody)
	if err != nil {
		t.Fatal(err)
	}

	w = httptest.NewRecorder()
	app.Router.ServeHTTP(w, r)

	assert.Equal(t, w.Code, http.StatusCreated)

	response := decodeBracket(w, t)
	check := response.Teams[7]

	assert.Equal(t, check.Name, "Team8")
	assert.Equal(t, check.Index, 7)
	assert.Equal(t, check.Round, 1)
	assert.Equal(t, check.Position, 8)
	assert.Equal(t, check.Eliminated, false)
}

func TestUpdateTeam(t *testing.T) {
	app, w := setup()

	newTeam := model.Team{
		Name:       "Lamda",
		Index:      7,
		Round:      1,
		Position:   8,
		Eliminated: true,
	}

	jsonData, err := json.Marshal(newTeam)
	if err != nil {
		t.Fatal(err)
	}

	requestBody := bytes.NewBuffer(jsonData)

	url := "/users/testuser/" + getTestBracketID(t, app, w) + "/teams/7"
	r, err := http.NewRequest("PUT", url, requestBody)
	if err != nil {
		t.Fatal(err)
	}

	w = httptest.NewRecorder()
	app.Router.ServeHTTP(w, r)

	assert.Equal(t, w.Code, http.StatusOK)

	response := decodeBracket(w, t)
	check := response.Teams[7]

	assert.Equal(t, check.Name, "Lamda")
	assert.Equal(t, check.Index, 7)
	assert.Equal(t, check.Round, 1)
	assert.Equal(t, check.Position, 8)
	assert.Equal(t, check.Eliminated, true)
}

func TestDeleteTeam(t *testing.T) {
	app, w := setup()

	url := "/users/testuser/" + getTestBracketID(t, app, w) + "/teams/7"
	r, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	w = httptest.NewRecorder()
	app.Router.ServeHTTP(w, r)

	assert.Equal(t, w.Code, http.StatusNoContent)
}
