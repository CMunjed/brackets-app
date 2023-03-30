package testing

import (
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
		{
			Name:       "Team8",
			Index:      8,
			Round:      0,
			Position:   8,
			Eliminated: false,
		},
		{
			Name:       "Team9",
			Index:      8,
			Round:      0,
			Position:   9,
			Eliminated: false,
		},
		{
			Name:       "Team10",
			Index:      9,
			Round:      0,
			Position:   10,
			Eliminated: false,
		},
		{
			Name:       "Team11",
			Index:      10,
			Round:      0,
			Position:   11,
			Eliminated: false,
		},
		{
			Name:       "Team12",
			Index:      11,
			Round:      0,
			Position:   12,
			Eliminated: false,
		},
		{
			Name:       "Team13",
			Index:      12,
			Round:      0,
			Position:   13,
			Eliminated: false,
		},
		{
			Name:       "Team14",
			Index:      13,
			Round:      0,
			Position:   14,
			Eliminated: false,
		},
		{
			Name:       "Team15",
			Index:      14,
			Round:      0,
			Position:   15,
			Eliminated: false,
		},
		{
			Name:       "Team16",
			Index:      15,
			Round:      0,
			Position:   16,
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

	url := "/users/testuser/" + Test_bracket_id + "/teams"
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}

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

	for i := 0; i < len(response); i++ {
		assert.Equal(t, response[i].Name, test_bracket_teams[i].Name)
		assert.Equal(t, response[i].Index, test_bracket_teams[i].Index)
		assert.Equal(t, response[i].Round, test_bracket_teams[i].Round)
		assert.Equal(t, response[i].Position, test_bracket_teams[i].Position)
		assert.Equal(t, response[i].Eliminated, test_bracket_teams[i].Eliminated)
	}
}

func TestGetTeam(t *testing.T) {
	app, w := setup()

	url := "/users/testuser/" + Test_bracket_id + "/teams/0"
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	app.Router.ServeHTTP(w, r)

	assert.Equal(t, w.Code, http.StatusOK)

	response := decodeTeam(w, t)

	assert.Equal(t, response, test_bracket_teams[0])
}
