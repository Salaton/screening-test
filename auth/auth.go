package auth

import (
	"log"

	"github.com/dgrijalva/jwt-go"
)

var SecretKey = []byte("SecretKeyThatShouldBeInTheEnvFile")

// CreateNewToken function that generates a token from a username
func CreateNewToken(username string) (string, error) {

	// Create a new token Having a map of out claims
	// Statements about out user..
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
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
		username := claims["username"].(string)
		// password := claims["password"].(string)
		return username, nil
	}
	return "", err

}
