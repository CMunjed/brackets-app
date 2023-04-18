package model

import (
	//"github.com/google/uuid"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique" json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	UserID   string `gorm:"unique" json:"userid"`
	Admin    bool   `json:"admin"`
}

type GoogleUser struct {
	//Not a gorm model bc this won't be saved in DB
	//gorm.Model
	Aud    string `json:"aud"`
	Azp    string `json:"azp"`
	Email  string `json:"email"`
	Emailv bool   `json:"email_verified"`
	Exp    int    `json:"exp"`
	Gname  string `json:"given_name"`
	//Fname	 string `json:"family_name"`
	Iat    int    `json:"iat"`
	Iss    string `json:"iss"`
	Jti    string `json:"jti"`
	Name   string `json:"name"`
	Nbf    int    `json:"nbf"`
	Imgurl string `json:"picture"`
	Id     string `json:"sub"`
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
	Name         string        `json:"name"`
	BracketID    string        `gorm:"unique" json:"bracketid"`
	UserID       string        `json:"userid"`
	Size         int           `json:"size"`
	Matches      int           `json:"matches"`
	Type         int           `json:"type"`
	Public       bool          `json:"public"`
	Edit         bool          `json:"edit"`
	AllowedUsers []AllowedUser `json:"allowedusers"`
	Teams        []Team        `json:"teams"`
}

type AllowedUser struct {
	gorm.Model
	BracketID   string `json:"bracketid"`
	AllowedUser string `json:"alloweduser"`
}

type Session struct {
	gorm.Model
	UserID string    `json:"userid"`
	Token  string    `gorm:"primaryKey" json:"token"`
	Expiry time.Time `json:"expiry"`
}

/*func (s Session) isExpired() bool {
	return s.Expiry.Before(time.Now())
}*/

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{}, &Team{}, &Bracket{}, &Session{}, &AllowedUser{})
	return db
}
