package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken is a function to generate a token based on the user id and role
func GenerateToken(id, role, secret string) (tokenString string, err error) {
	claims := jwt.MapClaims{
		"id":   id,
		"role": role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err = token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken is a function to validate the token and extract the claims
func ValidateToken(tokenString, secret string) (id string, role string, err error) {
	tokens, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return []byte(secret), nil
	})

	if err != nil {
		return
	}

	claims, ok := tokens.Claims.(jwt.MapClaims)
	if ok && tokens.Valid {
		id = fmt.Sprintf("%v", claims["id"])
		role = fmt.Sprintf("%v", claims["role"])
		return
	}

	err = fmt.Errorf("unable to extract claims")
	return
}
