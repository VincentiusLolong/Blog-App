package token

import (
	"fiber-mongo-api/models"

	"golang.org/x/crypto/bcrypt"
)

type UsetHash models.User

func (u *UsetHash) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}
func (u *UsetHash) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
