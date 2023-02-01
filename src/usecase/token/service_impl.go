package token

import (
	"shoeshop-backend/src/interfaces/http/context"
	"shoeshop-backend/src/shared/constant"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func (s *service) CheckAuth(ctx *context.ApplicationContext) (userId string, err error) {
	tokenString := getToken(ctx)
	if tokenString == "" {
		return "", constant.ErrorInvalidToken
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.Options.JwtSecret), nil
	})
	if err != nil {
		ctx.Logger.Error("CheckAuth Error: failed to parse token with error %s", err.Error())
		return "", constant.ErrorInvalidToken
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if userId, ok := claims[JWT_USER_KEY].(string); ok {
			return userId, nil
		}
	}

	return "", constant.ErrorInvalidTokenData
}

func (s *service) Generate(ctx *context.ApplicationContext, userId string) (result string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims[JWT_USER_KEY] = userId
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()
	tokenString, err := token.SignedString([]byte(s.cfg.Options.JwtSecret))
	if err != nil {
		ctx.Logger.Error("Generate Token: failed to generate token with error %s", err.Error())
		return "", err
	}

	return tokenString, nil
}
