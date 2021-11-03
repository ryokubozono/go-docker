package server

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
    "github.com/ryokubozono/go-docker/controller"
    "github.com/ryokubozono/go-docker/docs"

)

// Init is initialize server
func Init() {
	r := router()
	r.Run(":8080")
}

func router() *gin.Engine {
	docs.SwaggerInfo.BasePath = "/"
	r := gin.Default()
	ctrl := root.Controller{}
    r.GET("/", ctrl.RootGet)

    r.GET("/ping", ctrl.DbPing)
	
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}