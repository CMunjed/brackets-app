package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"example.com/api/app/model"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// AddTeam adds a team to a bracket and saves said bracket
func AddTeam(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	bracket := findBracket(db, w, r)
	if bracket != nil {
		edit, err := canEdit(db, r, bracket)
		if err != nil {
			respondError(w, http.StatusUnauthorized, err.Error())
			return
		}

		if edit {
			if len(bracket.Teams) < bracket.Size {
				team := model.Team{}

				decoder := json.NewDecoder(r.Body)
				if err := decoder.Decode(&team); err != nil {
					respondError(w, http.StatusBadRequest, err.Error())
					return
				}
				defer r.Body.Close()

				team.BracketID = bracket.BracketID
				team.Position = len(bracket.Teams) + 1
				team.Index = len(bracket.Teams)
				team.Round = 1
				team.Eliminated = false
				bracket.Teams = append(bracket.Teams, team)
				saveBracket(db, w, bracket)
				respondJSON(w, http.StatusCreated, &bracket)
			}
		}
	}
}

// GetTeam returns a specific team from a bracket
func GetTeam(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bracket := findBracket(db, w, r)
	if bracket != nil {
		view, err := canView(db, r, bracket)
		if err != nil {
			respondError(w, http.StatusUnauthorized, err.Error())
			return
		}

		if view {
			indexString := vars["index"]
			index, err := strconv.Atoi(indexString)
			if err != nil {
				respondError(w, http.StatusBadRequest, err.Error())
				return
			}
			team := bracket.Teams[index]
			respondJSON(w, http.StatusOK, team)
		}
	}
}

// GetAllTeams returns all the teams in the bracket
func GetAllTeams(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	bracket := findBracket(db, w, r)
	respondJSON(w, http.StatusOK, bracket.Teams)
}

// UpdateTeam edits a team within a bracket
func UpdateTeam(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bracket := findBracket(db, w, r)
	if bracket != nil {
		indexString := vars["index"]
		index, err := strconv.Atoi(indexString)
		if err != nil {
			respondError(w, http.StatusBadRequest, err.Error())
			return
		}
		team := bracket.Teams[index]

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&team); err != nil {
			respondError(w, http.StatusBadRequest, err.Error())
			return
		}
		defer r.Body.Close()

		bracket.Teams[index] = team

		saveBracket(db, w, bracket)
		respondJSON(w, http.StatusOK, &bracket)
	}
}

// DeleteTeam removes a team from a bracket
func DeleteTeam(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bracket := findBracket(db, w, r)
	if bracket != nil {
		indexString := vars["index"]
		index, err := strconv.Atoi(indexString)
		if err != nil {
			respondError(w, http.StatusBadRequest, err.Error())
			return
		}
		firstPartition := bracket.Teams[:index]
		if index+1 < len(bracket.Teams) {
			secondPartition := bracket.Teams[index+1 : len(bracket.Teams)]

			for i := 0; i < len(secondPartition); i++ {
				secondPartition[i].Index--
			}

			firstPartition = append(firstPartition, secondPartition...)
		}

		bracket.Teams = firstPartition
		saveBracket(db, w, bracket)
		respondJSON(w, http.StatusNoContent, &bracket)
	}
}
