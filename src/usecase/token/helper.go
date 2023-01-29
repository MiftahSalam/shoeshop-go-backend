package token

import (
	"shoeshop-backend/src/interfaces/http/context"
	"strings"
)

const JWT_USER_KEY = "userId"

func getToken(ctx *context.ApplicationContext) string {
	header := ctx.Context.Request().Header.Get("Authorization")
	if header == "" {
		ctx.Logger.Warning("getToken Error: no authorization")
		return ""
	}

	if strings.HasPrefix(header, "Bearer") {
		header = header[7:]
	}

	return header
}
