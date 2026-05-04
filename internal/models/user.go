package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey" form:"-"`
	Username  string    `json:"username" form:"username"`
	Email     string    `json:"email" gorm:"unique" form:"email"`
	Password  string    `json:"password" form:"password"`
	CreatedAt time.Time `json:"created_at" form:"-"`
	UpdatedAt time.Time `json:"updated_at" form:"-"`
}