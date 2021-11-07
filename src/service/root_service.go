package service

import (
	"github.com/gin-gonic/gin"
	"github.com/ryokubozono/go-docker/db"
	"github.com/ryokubozono/go-docker/entity"
	"github.com/ryokubozono/go-docker/pkg/crypto"
)

type SampleTable entity.SampleTable
type User entity.User

type RootService struct{}

func (s RootService) FirstSampleTable(c *gin.Context) (SampleTable, error) {

	db := db.GetDB()
	var sample SampleTable

	db.First(&sample, 1)

	return sample, nil
}

func (s RootService) CreateUser(username string, password string) error {
	db := db.GetDB()

	passwordEncrypt, _ := crypto.PasswordEncrypt(password)
	err := db.Create(&User{Username: username, Password: passwordEncrypt}).Error
	if err != nil {
		return err
	}

	return nil
}

func (s RootService) Login(username string, password string) error {
	dbPassword := getUser(username).Password

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
