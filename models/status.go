package models

import "time"

type Status struct {
    ID        uint       `gorm:"primary_key;AUTO_INCREMENT" json:"id" form:"id"`
    User      uint     `json:"user" form:"user"`
    Text  string     `json:"text" form:"text"`
    CreatedAt *time.Time `json:"created_at" form:"created_at"`
    UpdatedAt *time.Time `json:"updated_at" form:"updated_at"`
}
