package main

import (

    "github.com/ryokubozono/go-docker/db"
    "github.com/ryokubozono/go-docker/server"

)

func main() {
    db.Init()
    server.Init()

	defer db.Close()

}

