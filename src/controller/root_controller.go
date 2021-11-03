package root

import (
	"net/http"

	"github.com/gin-gonic/gin"
    "github.com/ryokubozono/go-docker/db"
    "github.com/ryokubozono/go-docker/entity"

)

type TestTable entity.TestTable

type Controller struct{}

func (pc Controller) RootGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
            "message": "hello world",
        })
	}
}

func (pc Controller) DbPing() gin.HandlerFunc{
	return func(c *gin.Context){
		db := db.GetDB()

        var user TestTable

        db.First(&user, 1)

        c.JSON(200, gin.H{
            "hello": user.Name,
        })
	}
}