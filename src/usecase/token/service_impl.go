package token

import (
	"shoeshop-backend/src/interfaces/http/context"

	"github.com/dgrijalva/jwt-go"
)

func (s *service) Generate(ctx *context.ApplicationContext, userId string) (result string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = userId
	tokenString, err := token.SignedString([]byte(s.cfg.Options.JwtSecret))
	if err != nil {
		ctx.Error(err)
		return "", err
	}

	return tokenString, nil
}
