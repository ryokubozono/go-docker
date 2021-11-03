package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
    "github.com/ryokubozono/go-docker/db"


)

type TestTable struct {
    ID   int
    Name string
}

func RootGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
            "message": "hello world",
        })
	}
}

func DbPing() gin.HandlerFunc{
	return func(c *gin.Context){
		db := db.GetDB()

        var user TestTable

        db.First(&user, 1)

        c.JSON(200, gin.H{
            "hello": user.Name,
        })
	}
}