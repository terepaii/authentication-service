package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User Mongo DB Model
type User struct {
	ID primitive.ObjectID `bson:"_id"`
	//First_Name *string            `json:"first_name" validate:"required,min=2,max=100"`
	//Last_Name  *string            `json:"last_name" validate:"required,min=2,max=100"`
	Password      *string   `json:"password" validate:"required,min=6"`
	User_name     *string   `json:"user_name" validate:"required,min=5"`
	Token         *string   `json:"token"`
	Refresh_token *string   `json:"refresh_token"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
	User_id       string    `json:"user_id"`
}
