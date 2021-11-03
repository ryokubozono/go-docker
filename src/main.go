package main

import (
    "github.com/gin-gonic/gin"

    "github.com/ryokubozono/go-docker/handler"
)

func main() {
    engine:= gin.Default()
    engine.GET("/", handler.RootGet())

    engine.GET("/ping", handler.DbPing())

    engine.Run(":8080")
}

