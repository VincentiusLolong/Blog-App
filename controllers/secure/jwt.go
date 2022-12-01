package secure

import (
	"errors"
	"time"

	"fiber-mongo-api/configs"
	"fiber-mongo-api/models"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(configs.AllEnv("JWT_SECRET_KEY"))

func GenerateJWT(email string, name string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &models.JWTClaim{
		Email: email,
		Name:  name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateToken(signedToken string) (set map[string]string, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&models.JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	credential, ok := token.Claims.(*models.JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if credential.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	set = map[string]string{
		"Email": credential.Email,
		"Name":  credential.Name}

	return
}
