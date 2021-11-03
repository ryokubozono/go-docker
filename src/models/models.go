package models

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)

var DB *gorm.DB

type User struct {
    ID   int
    Name string
}

// Setup initializes the database instance
func Setup() {
	var err error
	DB, err = gorm.Open("mysql", "root:password@tcp(mysql:3306)/test_db")

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
	DB.SingularTable(true)
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer DB.Close()
}
