package server

import (
	"github.com/gin-gonic/gin"

    "github.com/ryokubozono/go-docker/controller"
)

// Init is initialize server
func Init() {
	r := router()
	r.Run(":8080")
}

func router() *gin.Engine {
	r := gin.Default()
	ctrl := root.Controller{}
    r.GET("/", ctrl.RootGet)

    r.GET("/ping", ctrl.DbPing)
	
	return r
}