package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	UserID string
	jwt.RegisteredClaims
}

func generateToken(Id string) (string, error) {
	secret, isPresent := os.LookupEnv("JWT_SECRET")
	if !isPresent {
		return "", fmt.Errorf("JWT_SECRET is not present")
	}

	claims := &JwtClaims{
		UserID: Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func validate(tokenStr string) (bool, error) {
	claims := &JwtClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		secret, isPresent := os.LookupEnv("JWT_SECRET")
		if !isPresent || secret == "" {
			return nil, fmt.Errorf("JWT_SECRET is not present")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return false, err
	}
	if _, ok := token.Claims.(*JwtClaims); !ok {
		return false, fmt.Errorf("invalid token")
	}

	return token.Valid, nil
}
