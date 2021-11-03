package server

import (
	"github.com/gin-gonic/gin"

    "github.com/ryokubozono/go-docker/handler"
)

// Init is initialize server
func Init() {
	r := router()
	r.Run(":8080")
}

func router() *gin.Engine {
	r := gin.Default()

    r.GET("/", handler.RootGet())

    r.GET("/ping", handler.DbPing())
	
	return r
}