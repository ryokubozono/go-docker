package root

import (
	"github.com/gin-gonic/gin"
	// "net/http"
	"github.com/ryokubozono/go-docker/pkg/crypto"

	"github.com/ryokubozono/go-docker/entity"
    "github.com/ryokubozono/go-docker/db"
	"log"
)

type SampleTable entity.SampleTable
type User entity.User

type Service struct{}

func (s Service) FirstSampleTable(c *gin.Context) (SampleTable, error){

	db := db.GetDB()
	var sample SampleTable

	db.First(&sample, 1)

	return sample, nil
}

func (s Service) CreateUser(username string, password string) error {
	db := db.GetDB()

	passwordEncrypt, _ := crypto.PasswordEncrypt(password)
	err := db.Create(&User{Username: username, Password: passwordEncrypt}).Error
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func (s Service) Login(username string, password string) error {
	dbPassword := getUser(username).Password
	log.Print(dbPassword)
	if err := crypto.CompareHashAndPassword(dbPassword, password); err != nil {
		return err
	}
	return nil
}


func getUser(username string) User {
	db := db.GetDB()
	var user User
	db.First(&user, "username = ?", username)
	return user
}
