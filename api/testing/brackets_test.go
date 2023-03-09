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
	test_bracket_id = "77416fe0-c7f9-417a-af9e-8680064d2aa1"

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

	test_bracket = model.Bracket{
		Name:   "Test_Bracket",
		UserID: "12345",
		Size:   16,
		Teams:  test_bracket_teams,
	}
)

func setup() *app.App {
	app := &app.App{}
	app.Initialize()
	return app
}

func TestGetAllBrackets(t *testing.T) {
	r, err := http.NewRequest("GET", "/brackets", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	app := setup()

	app.GetAllBrackets(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateBracket(t *testing.T) {
	jsonData, err := json.Marshal(test_bracket)
	if err != nil {
		t.Fatal(err)
	}

	requestBody := bytes.NewBuffer(jsonData)
	r, err := http.NewRequest("POST", "/brackets", requestBody)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	app := setup()

	app.CreateBracket(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetBracket(t *testing.T) {
	address := "/brackets/"
	address += test_bracket_id
	r, err := http.NewRequest("GET", address, nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	app := setup()

	app.GetBracket(w, r)
	var response model.Bracket
	body, err := io.ReadAll(w.Result().Body)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, response.ID, test_bracket.ID)
}
