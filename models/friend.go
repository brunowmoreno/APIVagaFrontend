package models

import "time"

type Friend struct {
    ID        uint       `gorm:"primary_key;AUTO_INCREMENT" json:"id" form:"id"`
    Name      string     `json:"name" form:"name"`
    Avatar    string     `json:"avatar" form:"avatar"`
    CreatedAt *time.Time `json:"created_at" form:"created_at"`
    UpdatedAt *time.Time `json:"updated_at" form:"updated_at"`
}
