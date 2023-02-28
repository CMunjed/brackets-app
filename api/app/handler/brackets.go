package handler

import (
	"math"
	"net/http"
	"strconv"

	"example.com/api/app/model"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// getBracketOr404 gets a bracket instance if exists, or respond the 404 error otherwise
func getBracketOr404(db *gorm.DB, bracketID string, w http.ResponseWriter, r *http.Request) *model.Bracket {
	bracket := model.Bracket{}
	if err := db.First(&bracket, model.Bracket{BracketID: bracketID}).Error; err != nil {
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

// GetBracket Returns a bracket in JSON formatting
func GetBracket(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	bracket := findBracket(db, w, r)
	if bracket != nil {
		respondJSON(w, http.StatusOK, bracket)
	}
}

func CreateBracket(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bracket := model.Bracket{}

	name := vars["name"]
	userID := vars["userID"]
	sizeString := vars["size"]
	size, err := strconv.Atoi(sizeString)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	bracket.Name = name
	bracket.UserID = userID
	bracket.Size = size
	bracket.Matches = int(math.Ceil(math.Log2(float64(size))))
	bracket.BracketID = uuid.New().String()

	if err := db.Save(&bracket).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, bracket)
}

/*
func getUserBrackets()
search the database for all the brackets with the same userID
return the selected brackets
*/
