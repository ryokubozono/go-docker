package entity

import (
    "github.com/jinzhu/gorm"
)

type SampleTable struct {
    gorm.Model
    Name string `json:"name"`
}
