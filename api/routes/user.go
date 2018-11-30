package routes

import (
	"time"

	jwtLib "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type userResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func createAuthorizationToken(id uint, email, secret string) (string, error) {
	// Create the token
	token := jwtLib.New(jwtLib.GetSigningMethod("HS256"))

	// Set claims
	token.Claims = jwtLib.MapClaims{
		"id":    id,
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, nil
}
