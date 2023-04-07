package helper

import "github.com/dgrijalva/jwt-go"

const (
	secret = "@pr1j@l9h1y@sset1@w@n"
)

func GenerateToken(id uint, email string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}

	parsedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := parsedToken.SignedString([]byte(secret))

	return signedToken
}
