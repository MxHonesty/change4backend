package auth

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id       string `bson:"_id"`
	UserName string `bson:"userName"`
	PassWord string `bson:"password"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err != nil
}
