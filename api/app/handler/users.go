package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"example.com/api/app/model"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GetAllUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	users := []model.User{}
	db.Find(&users)
	respondJSON(w, http.StatusOK, users)
}

func SignUp(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	// Parse request into User instance
	user := model.User{}
	err := json.NewDecoder(r.Body).Decode(&user) // These lines changed slightly from tutorial
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		//w.WriteHeader(http.StatusBadRequest)
		respondError(w, http.StatusBadRequest, err.Error()) // The above line replaced with our respondError function created in common.go
		return
	}

	// Salt and hash the password using the bcrypt algorithm
	// The second argument is the cost of hashing, which we arbitrarily set as 8 (this value can be more or less, depending on the computing power you wish to utilize)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error()) // The above line replaced with our respondError function created in common.go
		return
	}

	user.Password = string(hashedPassword) // Is changing the user object's password to the hashed version the best way to pass the data to the db?
	user.UserID = uuid.New().String()

	// These lines changed from tutorial, insert new user into database
	if err := db.Save(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Respond with JSON
	respondJSON(w, http.StatusCreated, user)
}

func SignIn(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	// Parse request into User instance
	user := model.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		//w.WriteHeader(http.StatusBadRequest)
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	//Get first user from database matching username (changed to email)
	//storedUser := getUserOr404(db, username, w, r)
	storedUser := getUserFromEmailOr404(db, user.Email, w, r)
	if storedUser == nil {
		//respondError(w, http.StatusInternalServerError, err.Error())
		//Error called in getUserOr404 helper function I think
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		// If the two passwords don't match, return a 401 status
		respondError(w, http.StatusUnauthorized, err.Error())
	}

	user.UserID = storedUser.UserID

	GenerateUserSession(db, user.UserID, w)

	respondJSON(w, http.StatusOK, user)
}

func GetUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userid := vars["userid"]
	user := getUserOr404(db, userid, w, r)
	if user == nil {
		return
	}
	respondJSON(w, http.StatusOK, user)
}

func getUserOr404(db *gorm.DB, userid string, w http.ResponseWriter, r *http.Request) *model.User {
	user := model.User{}
	if err := db.First(&user, model.User{UserID: userid}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &user
}

func UpdateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userid := vars["userid"]
	user := getUserOr404(db, userid, w, r)
	if user == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	user.Password = string(hashedPassword)

	if err := db.Save(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, user)
}

func DeleteUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userid := vars["userid"]
	user := getUserOr404(db, userid, w, r)
	if user == nil {
		return
	}
	if err := db.Delete(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func GoogleSignUp(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)

	gdata := model.GoogleUser{}

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&gdata)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	//fmt.Println(gdata.Token)

	user := &model.User{
		Email:    gdata.Token.Email,
		Username: gdata.Token.Email[:strings.Index(gdata.Token.Email, "@")],
		Password: uuid.New().String(), //Generates a random UUID as a password, since the user will never log into this account without google
		//UserID:   uuid.New().String(),
	}

	/*
		email := gdata.Email
		username := gdata.Email[:strings.Index(gdata.Email, "@")-1]
		password := uuid.New().String()
	*/

	jsonData, err := json.Marshal(user)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
	}

	requestBody := bytes.NewBuffer(jsonData)

	newRequest, err := http.NewRequest("POST", "/users/signup", requestBody)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
	}

	SignUp(db, w, newRequest)
}

func GoogleSignIn(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	//Another consideration is just storing the google data as a user
	gdata := model.GoogleUser{}

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&gdata)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	//username := gdata.Email[:strings.Index(gdata.Email, "@")]
	email := "" //gdata.Email
	storedUser := getUserFromEmailOr404(db, email, w, r)
	if storedUser == nil {
		//respondError(w, http.StatusInternalServerError, err.Error())
		//Error called in getUserOr404 helper function I think
		return
	}

	GenerateUserSession(db, email, w)

	respondJSON(w, http.StatusOK, storedUser)
}

func getUserFromEmailOr404(db *gorm.DB, email string, w http.ResponseWriter, r *http.Request) *model.User {
	user := model.User{}
	if err := db.First(&user, model.User{Email: email}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &user
}

func GenerateUserSession(db *gorm.DB, userid string, w http.ResponseWriter) {
	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Second)

	userSession := &model.Session{
		UserID: userid,
		Expiry: expiresAt,
		Token:  sessionToken,
	}

	if err := db.Save(&userSession).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: expiresAt,
	})

	//respondJSON(w, http.StatusOK, userSession)
}

func RefreshSession(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			//w.WriteHeader(http.StatusUnauthorized)
			respondError(w, http.StatusUnauthorized, err.Error())
			return
		}
		//w.WriteHeader(http.StatusBadRequest)
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	sessionToken := c.Value

	userSession := &model.Session{}

	if err := db.First(&userSession, model.Session{Token: sessionToken}).Error; err != nil {
		respondError(w, http.StatusUnauthorized, err.Error())
		return
	}
	if isExpired(userSession) {
		if err := db.Delete(&userSession).Error; err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondError(w, http.StatusUnauthorized, "Unauthorized")
	}

	GenerateUserSession(db, userSession.UserID, w)

	if err := db.Delete(&userSession).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	//Any way to check the user session after calling the GenerateUserSession function?
	//Should this respondJSON?
	respondJSON(w, http.StatusOK, nil)
}

func isExpired(s *model.Session) bool {
	return s.Expiry.Before(time.Now())
}

func Logout(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			respondError(w, http.StatusUnauthorized, err.Error())
			return
		}
		// For any other type of error, return a bad request status
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	sessionToken := c.Value

	// Remove the user's session
	//delete(sessions, sessionToken)
	if err := db.Delete(model.Session{Token: sessionToken}).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	//Unsure if the above works, EXPERIMENTAL

	// We need to let the client know that the cookie is expired
	// In the response, we set the session token to an empty
	// value and set its expiry as the current time
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   "",
		Expires: time.Now(),
	})

	//Unsure how to respondJSON
	respondJSON(w, http.StatusOK, nil)
}

func Welcome(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	// We can obtain the session token from the requests cookies, which come with every request
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			respondError(w, http.StatusUnauthorized, err.Error())
			return
		}
		// For any other type of error, return a bad request status
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	sessionToken := c.Value

	// We then get the session from our session map
	userSession := &model.Session{}

	if err := db.First(&userSession, model.Session{Token: sessionToken}).Error; err != nil {
		respondError(w, http.StatusUnauthorized, err.Error())
		return
	}
	// If the session is present, but has expired, we can delete the session, and return
	// an unauthorized status
	if isExpired(userSession) {
		if err := db.Delete(&userSession).Error; err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondError(w, http.StatusUnauthorized, "Unauthorized")
	}

	// If the session is valid, return the welcome message to the user
	w.Write([]byte(fmt.Sprintf("Welcome %s!", userSession.UserID)))

	//Unsure how to respondJSON
	//RefreshSession(db, w, r)
	respondJSON(w, http.StatusOK, userSession)
}
