package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	//Email   string `gorm:"unique" json:"email"`
	Username   string `gorm:"unique" json:"username"`
	Password   string `json:"password"`
	//UUID   string `gorm:"unique" json:"uuid"`
	//UUID   int `gorm:"unique" json:"uuid"`
	//Status bool   `json:"status"`
}

/*type Employee struct {
	gorm.Model
	Name   string `gorm:"unique" json:"name"`
	City   string `json:"city"`
	Age    int    `json:"age"`
	Status bool   `json:"status"`
}*/

/*func (e *Employee) Disable() {
	e.Status = false
}

func (p *Employee) Enable() {
	p.Status = true
}*/

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{})
	return db
}
