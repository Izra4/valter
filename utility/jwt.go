package utility

import (
	"Valter/db/sqlc"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type Userlaims struct {
	ID uint32
	jwt.RegisteredClaims
}

func NewUserClaims(id uint32, exp time.Duration) Userlaims {
	return Userlaims{
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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, NewUserClaims(data.ID, exp))
	tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}
