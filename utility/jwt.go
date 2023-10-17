package utility

import (
	"Valter/db/sqlc"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type UserClaims struct {
	ID uint `json:"id"`
	jwt.RegisteredClaims
}

func NewUserClaims(id uint, exp time.Duration) UserClaims {
	return UserClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp)),
		},
	}
}

func Token(data sqlc.User) (string, error) {
	expStr := os.Getenv("EXP")
	exp, err := time.ParseDuration(expStr)
	if err != nil {
		exp = time.Hour * 2
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, NewUserClaims(uint(data.ID), exp))
	tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func DecodeToken(signedToken string, claims jwt.Claims, key string) (string, error) {
	token, err := jwt.ParseWithClaims(signedToken, claims, func(tk *jwt.Token) (interface{}, error) {
		_, ok := tk.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return "", errors.New("wrong signing method")
		}
		return []byte(key), nil
	})
	if err != nil {
		// parse failed
		return "", fmt.Errorf("token has been tampered with")
	}

	if !token.Valid {
		// token is not valid
		return "", fmt.Errorf("invalid token")
	}

	return signedToken, nil
}
