package routes

import (
	"log"
	"net/http"
	"os"

	"go-templ/infra/handler"
	"go-templ/infra/types"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func DocRoutes(e *echo.Group) {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	docs := e.Group("/document")
	secret := os.Getenv("SECRET")
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

	docs.Use(echojwt.WithConfig(config))
	docs.POST("", handler.CreateDocument)
}
