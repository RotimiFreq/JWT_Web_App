package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserData struct {
	Id        primitive.ObjectID `json:"id" validate:"required"`
	UserId    string             `json:"userid"`
	Firstname string             `json:"firstname" validate:"required"`
	Lastname  string             `json:"lastname" validate:"required"`
	Email     string             `json:"email" validate:"email"`
	Password  []byte             `json:"password" validate:"min=2 max=50"`
}
