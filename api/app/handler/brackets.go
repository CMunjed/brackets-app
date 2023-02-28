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

// saveBracket saves the bracket to a database and returns the saved bracket, or it responds with an Internal Server Error
func saveBracket(db *gorm.DB, w http.ResponseWriter, b *model.Bracket) {
	if err := db.Save(b).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, *b)
}

// GetBracket Returns a bracket in JSON formatting
func GetBracket(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	bracket := findBracket(db, w, r)
	if bracket != nil {
		respondJSON(w, http.StatusOK, bracket)
	}
}

// GetUserBrackets returns all the brackets tied to a user
func GetUserBrackets(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	userBrackets := []model.Bracket{}
	vars := mux.Vars(r)

	db.Find(&userBrackets, vars["userid"])
	respondJSON(w, http.StatusOK, userBrackets)
}

// GetAllBrackets returns all of the brackets in the database
func GetAllBrackets(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	brackets := []model.Bracket{}
	db.Find(&brackets)
	respondJSON(w, http.StatusOK, brackets)
}

// CreateBracket takes in a json and uses it to create a bracket
func CreateBracket(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bracket := model.Bracket{}

	name := vars["name"]
	userID := vars["userid"]
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

	saveBracket(db, w, &bracket)
}

// AddTeam adds a team to a bracket and saves said bracket
func AddTeam(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bracket := findBracket(db, w, r)
	if bracket != nil && len(bracket.Teams) < bracket.Size {
		team := model.Team{}

		team.Name = vars["teamName"]
		team.Position = len(bracket.Teams)
		team.Index = len(bracket.Teams) - 1
		team.Round = 1
		team.Eliminated = false
		bracket.Teams = append(bracket.Teams, team)
		saveBracket(db, w, bracket)
	}
}

// EditTeam edits a team within a bracket
/*func EditTeam(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bracket := findBracket(db, w, r)
	if bracket != nil {
		index := vars["index"]

	}
}*/
