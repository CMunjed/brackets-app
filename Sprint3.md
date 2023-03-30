# Sprint 3.md

## Video Link
[Link to Video](https://youtu.be/EvKcMkR6w4k)

## Work Done in Sprint 3
Front End:
During this sprint, much of the actual bracket creation was completed on the front end with a working system in place to create brackets with a user-entered number of teams. The user is also able to rename teams, change the title of the bracket, and click on the selected team to advance to the next stage. Work still needs to be done to integrate this with the back end and polish features, but much of the progress needed for this project was completed in this stage. Additionally, unit testing done through Cypress has provided to be a more versatile tool for what we're testing than using Karma, so the unit tests in this sprint were done entirely through Cypress.

Back End:
During this sprint, we added Google sign-up and sign-in and completely overhauled the existing user system to seamlessly integrate it with Google authentication. This included changing the means of identifying users in the code from a username-based approach to a mixed UUID and email-based approach. We also updated the routing system, brackets, unit tests, and everything else that worked with the username-based approach accordingly. Lastly, we created a better unit testing system using Go's included functionality rather than Postman or CURL commands, created a wiki (https://github.com/RetroSpaceMan123/brackets-app/wiki) as an easier method of sharing information between the front-end and back-end, and also began working on cookies.

## Front-End Unit Tests
create and run web application
ensure existence of and click on sign in component
ensure existence of and click on add teams component, and subsequent creation and clicking of edit team selection
ensure existence of and click on github star button

looking to add E2E tests to enter certain number of teams and click through created bracket, but text entry has been a pain so far


## Back-end Unit Tests

### Google Signup and Signin Tests

```
curl -X POST -H 'Content-Type:application/json' -d '{"iss":"test","nbf":123456789,"aud":"test","sub":"test","email":"test@email.com","email_verified":true,"azp":"test","name":"test","picture":"test","given_name":"test","iat":123456789,"exp":123456789,"jti":"test"}' localhost:3000/users/googlesignup

curl localhost:3000/users

curl -X PUT -H 'Content-Type:application/json' -d '{"iss":"test","nbf":123456789,"aud":"test","sub":"test","email":"test@email.com","email_verified":true,"azp":"test","name":"test","picture":"test","given_name":"test","iat":123456789,"exp":123456789,"jti":"test"}' localhost:3000/users/googlesignin
```

### Golang Users Unit Tests
```
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
}

func TestUpdateEmail(t *testing.T) {
	app, w := setup()

	user := login(t, app, w)
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
	app.Router.ServeHTTP(w, r)

	response := decodeUser(w, t)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, user.UserID, response.UserID)
	assert.Equal(t, user_update.Email, response.Email)
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
	url := "/users/" + dummy_user.UserID

	r, err = http.NewRequest("DELETE", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	app.Router.ServeHTTP(w, r)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestGetAllUsers(t *testing.T) {
	app, w := setup()

	r, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	app.Router.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}
```

### Golang Brackets Unit Tests
```
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

	app, w := setup()

	test_bracket_id := getTestBracketID(t, app, w)
	address := "/users/testuser/" + test_bracket_id
	r, err := http.NewRequest("GET", address, nil)
	if err != nil {
		t.Fatal(err)
	}

	w = httptest.NewRecorder()
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
```

### Golang Teams Unit Tests
```
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

	url := "/users/testuser/" + getTestBracketID(t, app, w) + "/teams/0"
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	w = httptest.NewRecorder()
	app.Router.ServeHTTP(w, r)

	assert.Equal(t, w.Code, http.StatusOK)

	response := decodeTeam(w, t)

	assert.Equal(t, response, test_bracket_teams[0])
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
	assert.Equal(t, check.Round, 0)
	assert.Equal(t, check.Position, 8)
	assert.Equal(t, check.Eliminated, true)
}

func TestUpdateTeam(t *testing.T) {
	app, w := setup()

	newTeam := model.Team{
		Name:       "Lamda",
		Index:      7,
		Round:      0,
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
	assert.Equal(t, check.Round, 0)
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

	assert.Equal(t, w.Code, http.StatusOK)
}
```


## Updated Back-end Documentation

https://github.com/RetroSpaceMan123/brackets-app/wiki
