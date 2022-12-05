package secure

import (
	"fmt"
	"time"

	"fiber-mongo-api/configs"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(configs.AllEnv("JWT_SECRET_KEY"))

func GenerateJWT(payload interface{}, ttl time.Duration) (string, error) {
	claims := make(jwt.MapClaims)
	claims["sub"] = 1
	claims["dat"] = payload
	claims["nbf"] = time.Now().Unix()
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * ttl).Unix()
	t, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

func ValidateToken(signedToken string) (map[string]interface{}, error) {
	tokenValidate, err := jwt.Parse(signedToken, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}
	claims, ok := tokenValidate.Claims.(jwt.MapClaims)
	if !ok || !tokenValidate.Valid {
		return nil, fmt.Errorf("validate: invalid")
	}

	return claims["dat"].(map[string]interface{}), nil
}
