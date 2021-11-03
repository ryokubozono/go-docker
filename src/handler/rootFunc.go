package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
    "github.com/ryokubozono/go-docker/models"


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

        test_table := TestTable{}

        models.DB.First(&test_table, 1)

        c.JSON(200, gin.H{
            "hello": test_table.Name,
        })
	}
}