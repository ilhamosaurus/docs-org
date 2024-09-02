package routes

import (
	"log"
	"net/http"
	"os"

	app "go-templ/domain"
	"go-templ/infra/types"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func DashboardRoutes(e *echo.Echo) {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	secret := os.Getenv("SECRET")
	d := e.Group("/dashboard")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return &types.JwtCustomClaims{}
		},
		SigningKey:  []byte(secret),
		TokenLookup: "cookie:Authorization",
		ErrorHandler: func(c echo.Context, err error) error {
			return c.Redirect(http.StatusFound, "/")
		},
	}
	d.Use(echojwt.WithConfig(config))
	d.GET("", app.Dashboard)
}
