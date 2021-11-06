package entity

import (
    "github.com/jinzhu/gorm"
)

type User struct {
    gorm.Model
    Username string `form:"username" binding:"required" gorm:"unique;not null"`
    Password string `form:"password" binding:"required" gorm:"not null"`
}