package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ryokubozono/go-docker/entity"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	db, err = gorm.Open("mysql", "root:password@tcp(mysql:3306)/gin_db?parseTime=true")

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	autoMigration()
}

// GetDB is called in models
func GetDB() *gorm.DB {
	return db
}

// Close is closing db
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func autoMigration() {
	db.AutoMigrate(&entity.SampleTable{}, &entity.User{})
}
