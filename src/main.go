package main

import (
    "github.com/ryokubozono/go-docker/db"
    "github.com/ryokubozono/go-docker/server"
)

// @title go-docker API
// @version 1.0
// @description An example of gin
// @license.name MIT
// @host 192.168.2.117:3000
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
    db.Init()
    server.Init()

    defer db.Close()

}
