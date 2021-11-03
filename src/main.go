package main

import (
    "github.com/gin-gonic/gin"

    "github.com/ryokubozono/go-docker/handler"
    "github.com/ryokubozono/go-docker/db"

)

func main() {
    db.Init()

	defer db.Close()

    engine:= gin.Default()
    engine.GET("/", handler.RootGet())

    engine.GET("/ping", handler.DbPing())

    engine.Run(":8080")
    db.Close()
}

