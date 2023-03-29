package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
	//UUID   int `gorm:"unique" json:"uuid"`
	//Status bool   `json:"status"`
}

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
	Type      int    `json:"type"`
	Teams     []Team `json:"teams"`
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{}, &Team{}, &Bracket{})
	return db
}
