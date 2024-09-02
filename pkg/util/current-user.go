package util

import (
	"go-templ/infra/types"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func CurrentUser(c echo.Context) *types.JwtCustomClaims {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*types.JwtCustomClaims)
	return claims
}
