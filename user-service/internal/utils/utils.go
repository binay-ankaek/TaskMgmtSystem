package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

var jwtKey = []byte(os.Getenv("SECRET_KEY"))

func HashPassword(password string) (string, error) {
	// Hash the password using a secure hash function
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err

	}
	return string(hashedPassword), nil
}
func ComparePassword(hashedPassword, password string) error {
	// Compare the hashed password with the provided password
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

}

func GenerateJWT(UserID string) (string, error) {
	//expiration time
	expirationTime := time.Now().Add(24 * time.Hour)
	//create token
	claim := &Claims{
		// User ID
		UserID: UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			// Set the expiration time
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	//create token variable declare
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	//signed jwt token
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
