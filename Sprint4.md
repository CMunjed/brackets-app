# Sprint 4.md

## Video Link
[Link to Video](https://youtu.be/gAWW92HbVMQ)

## Work Done in Sprint 4
Front End:
During this sprint, features of the bracket were added and polished up, integration with the back end was done, and the footer was rewrote to adjust position based on how much content is on screen and keep its size static. Additionally, work was done to add a double elimination system of bracket generation and it is functional, although it isn't quite where we would like it to be at this point. UI elements were added and modified to better support the new integration with the backend to make a functioning login system. Finally, unit testing through cypress was successful in testing new funcionality, such as the sign in button and double elimination slider.

Back End:
During this sprint, we implemented user sessions and user permisions. We also redid our initialization function enable CORS permissions, which allowed us to integrate with the frontend. Users sessions work through cookies. Once a user logs in, they recieve a cookie that represents their current session. When interacting with brackets, this cookie will be passed to the backend to validate what user they are. User permissions also work through cookies. Brackets now store a whitelist of users which can view or edit the bracket depending on what settings the creator of the bracket sets. Only the creator of the bracket can delete the bracket. An admin flag has also been added, which was intended to block regular users from calling certain API calls, however, due to time constraints, this functionally has not been implemented at this time.

## Front-End Unit Tests

open web application on local host
compounding test, click on add teams button and edit teams dropdown
ensure existence of and test click on google sign in button
ensure existence of and test click on other sign in button
ensure existence of and test click on github button
test slide functionality
test slide functionality through multiple uses

## Back-end Unit Tests

### users_test.go
```
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
	test_user2 = model.User{
		Email:    "testemail2@test.com",
		Username: "testuser2",
		Password: "testpassword",
		Admin:    true,
	}
	test_user3 = model.User{
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

func signup(t *testing.T, a *app.App, w *httptest.ResponseRecorder, user model.User) {
	jsonData, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}

	requestBody := bytes.NewBuffer(jsonData)

	r, err := http.NewRequest("POST", "/users/signup", requestBody)
	if err != nil {
		t.Fatal(err)
	}
	a.Router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func loginUser(t *testing.T, a *app.App, w *httptest.ResponseRecorder, user model.User) model.User {
	jsonData, err := json.Marshal(user)
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

func login2(t *testing.T, a *app.App, w *httptest.ResponseRecorder) model.User {
	jsonData, err := json.Marshal(test_user2)
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
func login3(t *testing.T, a *app.App, w *httptest.ResponseRecorder) model.User {
	jsonData, err := json.Marshal(test_user3)
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

/*func TestUpdatePassword(t *testing.T) {
	app, w := setup()

	user := login3(t, app, w)

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

	test_user3.Password = user_update.Password
	test_user3.UserID = user.UserID

	jsonData, err = json.Marshal(test_user3)
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

	test_user3.Password = user_update.Password
	test_user3.UserID = user.UserID

	jsonData, err = json.Marshal(test_user3)
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

	user := login3(t, app, w)
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
}*/

func TestUpdateUser(t *testing.T) {
	app, w := setup()

	signup(t, app, w, test_user3)
	w = httptest.NewRecorder()

	user := login3(t, app, w)
	c := w.Result().Cookies()
	w = httptest.NewRecorder()
	type update struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	user_update := update{
		Email:    "electronic@arts.com",
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
	r.AddCookie(c[0])
	app.Router.ServeHTTP(w, r)

	response := decodeUser(w, t)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, user.UserID, response.UserID)
	assert.Equal(t, user_update.Email, response.Email)
	assert.Equal(t, user_update.Password, response.Password)
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

/*func TestGoogleSignup(t *testing.T) {
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
}*/

func TestGoogleSignIn(t *testing.T) {
	app, w := setup()

	jsonData, err := json.Marshal(googlesignin_user)
	if err != nil {
		t.Fatal(err)
	}

	requestBody := bytes.NewBuffer(jsonData)

	r, err := http.NewRequest("POST", "/users/googlesignin", requestBody)
	if err != nil {
		t.Fatal(err)
	}
	app.Router.ServeHTTP(w, r)

	assert.Equal(t, http.StatusCreated, w.Code)

	r, err = http.NewRequest("POST", "/users/googlesignin", requestBody)
	if err != nil {
		t.Fatal(err)
	}
	app.Router.ServeHTTP(w, r)

	assert.Equal(t, http.StatusCreated, w.Code)
}

```

### sessions_tests.go
```
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

```
### brackets_test.go
```
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
	//login(t, app, w)
	user := login(t, app, w)
	test_bracket.UserID = user.UserID

	jsonData, err := json.Marshal(test_bracket)
	if err != nil {
		t.Fatal(err)
	}

	requestBody := bytes.NewBuffer(jsonData)
	r, err := http.NewRequest("POST", "/users/"+user.UserID+"/brackets", requestBody)
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
	address := "/users/" + user.UserID + "/" + test_bracket_id
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
		UserID: user.UserID,

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

func TestUnauthorizedBracketEdit(t *testing.T) {
	//Sign Up a new user
	app, w := setup()
	user := model.User{
		Username: "scriptkid",
		Password: "123456",
		Email:    "hackerman123@live.com",
	}

	signup(t, app, w, user)

	// Login as new user
	w = httptest.NewRecorder()
	loginUser(t, app, w, user)

	// Try and edit the test bracket
	c := w.Result().Cookies()
	w = httptest.NewRecorder()
	test_bracket_id := getTestBracketID(t, app, w)
	url := "/users/testuser/"
	url += test_bracket_id
	r, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	w = httptest.NewRecorder()
	r.AddCookie(c[0])
	app.Router.ServeHTTP(w, r)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

```
### teams_test.go
```
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
	login(t, app, w)
	c := w.Result().Cookies()
	w = httptest.NewRecorder()
	r.AddCookie(c[0])
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
	login(t, app, w)
	c := w.Result().Cookies()
	w = httptest.NewRecorder()
	r.AddCookie(c[0])
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
	login(t, app, w)
	c := w.Result().Cookies()
	w = httptest.NewRecorder()
	r.AddCookie(c[0])
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
	login(t, app, w)
	c := w.Result().Cookies()
	w = httptest.NewRecorder()
	r.AddCookie(c[0])
	app.Router.ServeHTTP(w, r)

	assert.Equal(t, w.Code, http.StatusNoContent)
}

```

## Updated Back-end Documentation

https://github.com/RetroSpaceMan123/brackets-app/wiki
