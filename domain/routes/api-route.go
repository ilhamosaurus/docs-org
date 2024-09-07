package routes

import (
	"go-templ/infra/handler"

	"github.com/labstack/echo/v4"
)

func ApiRoutes(e *echo.Echo) {
	api := e.Group("/api")

	api.POST("/login", handler.Login)
	api.POST("/register", handler.Register)
	api.POST("/logout", handler.Logout)

	DocRoutes(api)
}
