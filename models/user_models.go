package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id         primitive.ObjectID `json:"id,omitempty"`
	Name       string             `json:"name,omitempty" validate:"required"`
	Age        string             `json:"age,omitempty" validate:"required"`
	Orgs       string             `json:"orgs,omitempty" validate:"required"`
	About      string             `json:"about,omitempty" validate:"required"`
	Allcontent *AllContents       `json:"allcontents,omitempty" validate:"required"`
}

type AllContents struct {
	ContentId primitive.ObjectID `json:"contentid,omitempty"`
	Title     string             `json:"title,omitempty" validate:"required"`
	Tags      string             `json:"tag,omitempty" validate:"required"`
	Type      string             `json:"type,omitempty" validate:"required"`
	Photos    string             `json:"photos,omitempty" validate:"required"`
	Content   string             `json:"content,omitempty" validate:"required"`
}
