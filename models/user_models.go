package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type signIn struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Email    string             `json:"email,omitempty" validate:"required"`
	Password string             `json:"password,omitempty" validate:"required"`
}
type User struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Email    string             `json:"email,omitempty" validate:"required"`
	Password string             `json:"password,omitempty" validate:"required"`
	Name     string             `json:"name,omitempty" validate:"required"`
	Age      string             `json:"age,omitempty" validate:"required"`
	Orgs     string             `json:"orgs,omitempty" validate:"required"`
	About    string             `json:"about,omitempty" validate:"required"`
	Gender   string             `json:"gender,omitempty" validate:"required"`
}

type AllContents struct {
	ContentId primitive.ObjectID `json:"contentid,omitempty"`
	User_id   string             `json:"user_id,omitempty"`
	Title     string             `json:"title,omitempty" validate:"required"`
	Tags      string             `json:"tag,omitempty" validate:"required"`
	Type      string             `json:"type,omitempty" validate:"required"`
	Photos    string             `json:"photos,omitempty" validate:"required"`
	Content   string             `json:"content,omitempty" validate:"required"`
}
