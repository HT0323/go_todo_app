package entity

import "time"

type UserID int64

type User struct {
	ID       UserID
	Name     string    `json:"id" db:"id"`
	Password string    `json:"name" db:"name"`
	Role     string    `json:"role" db:"role"`
	Created  time.Time `json:"created" db:"created"`
	Modified time.Time `json:"modified" db:"modified"`
}
