package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	password := "super-secret-password-here"
	hash, _ := HashPassword(password) // Ignoring error for sake of simplicity

	fmt.Println("Password: ", password)
	fmt.Println("Hash: ", hash)

	match := CheckPasswordHash(password, hash)
	fmt.Println("Match: ", match)

	fmt.Println("End of Main")
}
