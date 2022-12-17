package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Login struct {
	Email    string `json:"email,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
}
type User struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Email    string             `json:"email,omitempty" validate:"required"`
	Password string             `json:"password,omitempty" validate:"required"`
	Name     string             `json:"name,omitempty" validate:"required"`
	Age      string             `json:"age,omitempty" validate:"required"`
	Orgs     string             `json:"orgs,omitempty"`
	About    string             `json:"about,omitempty"`
	Gender   string             `json:"gender,omitempty" validate:"required"`
}

type UserPorfile struct {
	Name   string `json:"name,omitempty" validate:"required"`
	Age    string `json:"age,omitempty" validate:"required"`
	Orgs   string `json:"orgs,omitempty"`
	About  string `json:"about,omitempty"`
	Gender string `json:"gender,omitempty" validate:"required"`
}

// type ChangePass struct{

// }
type AllContents struct {
	Id      primitive.ObjectID `json:"id,omitempty"`
	User_id string             `json:"user_id,omitempty"`
	Title   string             `json:"title,omitempty" validate:"required"`
	Tags    string             `json:"tag,omitempty" validate:"required"`
	Type    string             `json:"type,omitempty" validate:"required"`
	Photos  string             `json:"photos,omitempty" validate:"required"`
	Content string             `json:"content,omitempty" validate:"required"`
}

type Comments struct {
	Id primitive.ObjectID `json:"id,omitempty"`
}
