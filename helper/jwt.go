package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pebruwantoro/technical_test_dealls/config"
	"github.com/pebruwantoro/technical_test_dealls/repository"
)

type Claims struct {
	UUID      string `json:"uuid"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	IsPremium bool   `json:"is_premium"`
	jwt.RegisteredClaims
}

func GenerateJWT(user repository.User) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &Claims{
		UUID:      user.UUID,
		Email:     user.Email,
		Username:  user.Username,
		IsPremium: user.IsPremium,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.JWT_ISSUER,
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(config.JWT_SECRET_KEY))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
