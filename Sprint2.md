# Sprint 2.md

## Video Link
[Link to Video](https://youtu.be/003O3YYKI4Y)

## Work Done in Sprint 2
Front-End:
During this sprint, we implemented many features necessary for user authentication on the front end, which will be used for bracket users. This was done to establish communication between the front-end and back-end and begin the process of integration. Adding functionality for cypress testing and user testing using Karma was done, despite many integration issues due to cypress, angular, and included packages such as the mat slide toggle having complications with each other.

Back-End:
During this sprint, we were able to figure out how to implement user authentication and basic bracket functionality. In regard to user authentication, we are able to handle sign-ups and logins, and store user profiles which will eventually be linked to the brackets they have created. In regard to brackets, we we added functionality to create, edit, store, and delete brackets, along with the teams within the brackets. The brackets will also have user IDs stored with them, allowing us to link users to their brackets in the future.

## Front-End Unit Tests
create sign in component
create and run web application
authenticate the title of the application
ensure proper positioning of elements in application

## Back-end Unit Tests
[![Run in Postman](https://run.pstmn.io/button.svg)](https://god.gw.postman.com/run-collection/26133312-0e510e7f-2de9-454d-90f8-0db0a157ad4e?action=collection%2Ffork&collection-url=entityId%3D26133312-0e510e7f-2de9-454d-90f8-0db0a157ad4e%26entityType%3Dcollection%26workspaceId%3D9510da33-894b-400b-9e39-abd3e0122696)

## Back-end Documentation

### Users
 ```
 type User struct {
	gorm.Model
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
	UUID     string `json:"uuid"`
    }
```
* User Struct
  * Username
    * Name Given by the user to identify themselves when logging in
  * Password
    * Hashed password used by the user to login
  * UUID
    * Unique 128 character ID given to a user to use for internal identification
* User Functions
  * GetAllUsers
    * Gets all of the users in the database
    * Route: "/users"
  * SignUp
    * Handles sign-up events, and creates a new user in the database
    * Route: "/users/signup"
  * SignIn
    * Handles sign-in events, and responds with an User object on succesful login
    * Route: "/users/signin"
  * GetUser
    * Returns an User based on its username from the database as a JSON string, or returns a 404 error
    * Route: "/users/{username}"
  * UpdateUser
    * Updates the User object in the database that has the same username, and returns the updated User object
    * Route: "/users/{username}"
  * DeleteUser
    * Deletes the user based on the specified username
    * Route: "/users/{username}"

### Brackets and Teams
```
type Team struct {
	gorm.Model
	Name       string `json:"name"`
	BracketID  string `json:"bracketid"`
	Index      int    `json:"index"`
	Round      int    `json:"round"`
	Position   int    `json:"position"`
	Eliminated bool   `json:"eliminated"`
}

type Bracket struct {
	gorm.Model
	Name      string `json:"name"`
	BracketID string `gorm:"unique" json:"bracketid"`
	UserID    string `json:"userid"`
	Size      int    `json:"size"`
	Matches   int    `json:"matches"`
	Teams     []Team `json:"teams"`
}
```
* Team Struct
  * Name
    * Name of the team
  * BracketID
    * UUID of the Bracket 
  * Index
    * Index of the team in the Bracket's Array
  * Round
    * The current round the team is in
  * Position
    * The position the team currently is in the round
  * Eliminated
    * Flags if the team is eliminated from the tournament
* Bracket Struct
  * Name
    * Name of the Bracket
  * BracketID
    * UUID of the Bracket 
  * UserID
    * UUID of the owner of the Bracket
  * Size
    * The ammount of teams
  * Matches
    * The ammount of matches the bracket will have
  * Teams
    * Teams competing in the bracket
* Bracket Functions
  * GetAllBrackets
    * Gets all of the brackets in the database
    * Route: "/brackets"
  * GetUserBracket
    * Gets all brackets the specified user owns
    * Route: "/users/{userid}/brackets"
  * CreateBracket
    * Creates a bracket that's tied to a user id
    * Route "/brackets"
  * GetBracket
    * Returns a bracket based on its bracket ID from the database as a JSON string, or returns a 404 error
    * Route: "/brackets/{bracketid}"
  * UpdateBracket
    * Updates the Bracket object in the database that has the same bracket ID, and returns the updated User object
    * Route: "/brackets/{bracketid}"
  * DeleteBracket
    * Deletes the bracket based on the specified bracket ID
    * Route: "/brackets/{bracketid}"
* Team Functions
  * GetAllTeams
    * Gets all of the teams in the specified bracket
    * Route: "/brackets/{bracketid}/teams"
  * AddTeam
    * Adds a team to the specified bracket if there's enough space, else returns a bad request error
    * Route: "/brackets/{bracketid}/teams"
  * GetTeam
    * Returns a team from a specified bracket based on its index as a JSON string, or returns a 404 error
    * Route: "/brackets/{bracketid}/teams/{index}"
  * UpdateTeam
    * Updates the team from the specified bracket based on its index, and returns the updated team
    * Route: "/brackets/{bracketid}/teams/{index}"
  * DeleteBracket
    * Deletes the team from the specified bracket based on its index
    * Route: "/brackets/{bracketid}/teams/{index}"