package auth

import (
	"log"

	"github.com/dgrijalva/jwt-go"
)

var SecretKey = []byte("SecretKeyThatShouldBeInTheEnvFile")

// var SecretKey = os.Getenv("HashingSecretKey")

// CreateNewToken function that generates a token from a email
func CreateNewToken(email string) (string, error) {

	// Create a new token Having a map of out claims
	// Statements about our customer..
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
	})

	newTokenString, err := newToken.SignedString(SecretKey)

	if err != nil {
		log.Fatal("An error occured in creating the string: ", err)
		return "", nil
	}

	return newTokenString, nil
}

// ParseToken function to validate the given tokenString. A token has to be present
// Want to have an idea of who sent the token
func ParseToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email := claims["email"].(string)
		return email, nil
	}
	return "", err

}
