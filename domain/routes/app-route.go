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

func AppRoute(e *echo.Echo) {
	e.GET("/", app.GetHome)
	e.GET("/login", app.LoginPage)
	e.GET("/register", app.RegisterPage)
	DashboardRoutes(e)

	doc := e.Group("/document")
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
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

	doc.Use(echojwt.WithConfig(config))
	doc.GET("", app.CreateDocument)
}
