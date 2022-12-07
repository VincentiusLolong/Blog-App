package models

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JWTClaim struct {
	Email string `json:"email,omitempty" validate:"required"`
	Name  string `json:"name,omitempty" validate:"required"`
	jwt.StandardClaims
}

type JWTRefreshClaim struct {
	jwt.StandardClaims
}

type GetDataToken struct {
	Id    primitive.ObjectID `json:"id,omitempty"`
	Email string             `json:"email,omitempty" validate:"required"`
	Name  string             `json:"password,omitempty" validate:"required"`
}
