package main

import (
    "github.com/gin-gonic/gin"
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/ryokubozono/go-docker/handler"
)

type User struct {
    ID   int
    Name string
}

func main() {
    engine:= gin.Default()
    engine.GET("/", handler.RootGet())

    engine.GET("/ping", func(c *gin.Context){
        db, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/test_db")
        defer db.Close()
        if err != nil {
            fmt.Println(err.Error())
        }

        rows, err := db.Query("SELECT * FROM test_table")
        defer rows.Close()
        if err != nil {
            fmt.Println(err)
        }

        user := User{}
        for rows.Next() {
            err = rows.Scan(&user.ID, &user.Name)
            if err != nil {
                fmt.Println(err)
            }
            fmt.Println(user)
        }

        c.JSON(200, gin.H{
            "hello": user.Name,
        })
    })
    engine.Run(":8080")
}
