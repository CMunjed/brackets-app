package handler

import (
	"encoding/json"
	"errors"
	"math"
	"net/http"

	"example.com/api/app/model"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// getBracketOr404 gets a bracket instance if exists, or respond the 404 error otherwise
func getBracketOr404(db *gorm.DB, bracketID string, w http.ResponseWriter, r *http.Request) *model.Bracket {
	bracket := model.Bracket{}
	if err := db.Preload("Teams").Preload("AllowedUsers").First(&bracket, model.Bracket{BracketID: bracketID}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &bracket
}

// findBracket gets a bracket from the database based on its ID, or it returns nil
func findBracket(db *gorm.DB, w http.ResponseWriter, r *http.Request) *model.Bracket {
	vars := mux.Vars(r)

	bracketID := vars["bracketid"]
	bracket := getBracketOr404(db, bracketID, w, r)

	return bracket
}

// saveBracket saves the bracket to a database and returns the saved bracket, or it responds with an Internal Server Error
func saveBracket(db *gorm.DB, w http.ResponseWriter, b *model.Bracket) {
	if err := db.Save(b).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

// canView checks if a user can view a bracket
func canView(db *gorm.DB, r *http.Request, b *model.Bracket) error {
	c := r.Cookies()[0]
	token := c.Value
	session := &model.Session{}
	if err := db.First(&session, model.Session{Token: token}).Error; err != nil {
		return err
	}

	if b.Public || b.UserID == session.UserID {
		return nil
	}

	for _, user := range b.AllowedUsers {
		if user.AllowedUser == session.UserID {
			return nil
		}
	}

	return errors.New("unauthorized")
}

// canEdit checks if a user can edit a bracket
func canEdit(db *gorm.DB, r *http.Request, b *model.Bracket) error {
	c := r.Cookies()[0]
	token := c.Value
	session := &model.Session{}
	if err := db.First(&session, model.Session{Token: token}).Error; err != nil {
		return err
	}

	if b.UserID == session.UserID {
		return nil
	}

	if b.Edit {
		for _, user := range b.AllowedUsers {
			if user.AllowedUser == session.UserID {
				return nil
			}
		}
	}

	return errors.New("unauthorized")
}

// GetBracket Returns a bracket in JSON formatting
func GetBracket(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	bracket := findBracket(db, w, r)
	if bracket != nil {
		err := canView(db, r, bracket)
		if err != nil {
			respondError(w, http.StatusUnauthorized, err.Error())
			return
		} else {
			respondJSON(w, http.StatusOK, bracket)
		}
	}
}

// GetUserBrackets returns all the brackets tied to a user
func GetUserBrackets(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	//Permissions
	//If the user is the owner of the bracket, they can view it

	userBrackets := []model.Bracket{}
	vars := mux.Vars(r)

	db.Preload("Teams").Preload("AllowedUsers").Find(&userBrackets, model.Bracket{UserID: vars["userid"]})
	respondJSON(w, http.StatusOK, userBrackets)
}

// GetAllBrackets returns all of the brackets in the database
func GetAllBrackets(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	brackets := []model.Bracket{}
	db.Preload("Teams").Preload("AllowedUsers").Find(&brackets)
	respondJSON(w, http.StatusOK, brackets)
}

// CreateBracket takes in a json and uses it to create a bracket
func CreateBracket(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	c := r.Cookies()[0]
	token := c.Value
	session := &model.Session{}
	if err := db.First(&session, model.Session{Token: token}).Error; err != nil {
		respondError(w, http.StatusUnauthorized, err.Error())
		return
	}

	bracket := model.Bracket{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&bracket); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if bracket.UserID != session.UserID {
		respondError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	bracket.Matches = int(math.Ceil(math.Log2(float64(bracket.Size))))
	bracket.BracketID = uuid.New().String()
	for i := 0; i < len(bracket.Teams); i++ {
		bracket.Teams[i].BracketID = bracket.BracketID
		bracket.Teams[i].Index = i
	}

	saveBracket(db, w, &bracket)
	respondJSON(w, http.StatusOK, &bracket)
}

// UpdateBracket updates the parameters of a bracket
func UpdateBracket(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	bracket := findBracket(db, w, r)

	if bracket != nil {
		err := canEdit(db, r, bracket)
		if err != nil {
			respondError(w, http.StatusUnauthorized, err.Error())
			return
		} else {
			decoder := json.NewDecoder(r.Body)
			if err := decoder.Decode(&bracket); err != nil {
				respondError(w, http.StatusBadRequest, err.Error())
				return
			}
			defer r.Body.Close()

			saveBracket(db, w, bracket)
			respondJSON(w, http.StatusOK, &bracket)
		}
	}
}

// DeleteBracket deletes a bracket from the database
func DeleteBracket(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	bracketID := vars["bracketid"]
	bracket := getBracketOr404(db, bracketID, w, r)
	if bracket == nil {
		return
	}

	err := canEdit(db, r, bracket)
	if err != nil {
		respondError(w, http.StatusUnauthorized, err.Error())
		return
	} else {
		if err := db.Delete(&bracket).Error; err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondJSON(w, http.StatusNoContent, nil)
	}
}
