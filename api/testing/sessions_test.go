package testing

/*import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	//"example.com/api/app/model"
	"github.com/go-playground/assert"
)

func testWelcome(t *testing.T) {

	//Sign up to create user, sign in to create a user session, call welcome route
	//	to test the user session to see if it's working
	//Unsure how to check cookies using unit tests

	app, w := setup()

	//Convert test user to json
	jsonData, err := json.Marshal(test_user)
	if err != nil {
		t.Fatal(err)
	}

	requestBody := bytes.NewBuffer(jsonData)

	//Sign up
	r, err := http.NewRequest("POST", "/users/signup", requestBody)
	if err != nil {
		t.Fatal(err)
	}

	app.Router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusCreated, w.Code)

	//Sign in
	r, err = http.NewRequest("PUT", "/users/signin", requestBody)
	if err != nil {
		t.Fatal(err)
	}
	app.Router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)

	//Get user (make sure user exists)
	url := "/users/" + test_user.Username

	r, err = http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	app.Router.ServeHTTP(w, r)

	response := decodeUser(w, t)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, test_user.Username, response.Username)

	//Welcome
	r, err = http.NewRequest("PUT", "/welcome", requestBody)
	if err != nil {
		t.Fatal(err)
	}
	app.Router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}
*/
