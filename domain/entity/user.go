package entity

import (
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
)

// FooEntity struct definition
type User struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	PasswordHash string `json:"password_hash"`
}

// NewFooEntity initialize MyEntity
func CreateUserEntity(name, password string) (*User, error) {
	id := uuid.New()
	if name == "" {
		return nil, fmt.Errorf("name is required")
	}

	return &User{
		ID:           id.String(),
		Name:         name,
		PasswordHash: makePasswordHash([]byte(password)),
	}, nil
}

func makePasswordHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
