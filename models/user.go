package models

import "time"

type User struct {
    ID        uint       `gorm:"primary_key;AUTO_INCREMENT" json:"id" form:"id"`
    Email      string     `json:"email" form:"email"`
    FirstName  string     `json:"first_name" form:"first_name"`
    LastName  string     `json:"last_name" form:"last_name"`
    CreatedAt *time.Time `json:"created_at" form:"created_at"`
    UpdatedAt *time.Time `json:"updated_at" form:"updated_at"`
}
