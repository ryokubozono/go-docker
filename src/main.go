package main

import (
    "github.com/gin-gonic/gin"

    "github.com/ryokubozono/go-docker/handler"
    "github.com/ryokubozono/go-docker/models"

)

func main() {
    

    models.Setup()
	defer models.CloseDB()

    engine:= gin.Default()
    engine.GET("/", handler.RootGet())

    engine.GET("/ping", handler.DbPing())

    engine.Run(":8080")
}

