package models

import (
	"time"

	guuid "github.com/google/uuid"
)

// User Mongo DB Model
type User struct {
	Id            guuid.UUID `bson:"_id"`
	Password      *string    `json:"password" validate:"required,min=6"`
	User_name     *string    `json:"user_name" validate:"required,min=5"`
	Token         *string    `json:"token"`
	Refresh_token *string    `json:"refresh_token"`
	Created_at    time.Time  `json:"created_at"`
	Updated_at    time.Time  `json:"updated_at"`
	User_id       string     `json:"user_id"`
}
