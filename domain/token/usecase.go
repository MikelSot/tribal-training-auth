package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/MikelSot/tribal-training-auth/model"
)

type Token struct {
	ExpiresAt int

	SignKey string
}

func New(expiresAt int, signKey string) Token {
	return Token{expiresAt, signKey}
}

func (t Token) Generate(m model.User) (string, error) {
	payload := jwt.MapClaims{
		"email": m.Email,
		"ia":    time.Now().Unix(),
		"exp":   time.Now().Add(time.Hour * time.Duration(t.ExpiresAt)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	key, err := token.SignedString([]byte(t.SignKey))
	if err != nil {
		return "", fmt.Errorf("error generating token: %w", err)
	}

	return key, nil
}
