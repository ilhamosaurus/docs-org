package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-templ/domain/routes"
	"go-templ/pkg/database"
	"go-templ/pkg/util"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "Method = ${method} | URL = \"${uri}\"| Status = ${status} | Latency = ${latency_human}\n",
	}))
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	app.Use(middleware.Recover())
	routes.AppRoute(app)
	app.Validator = util.NewCustomValidator()
	app.Static("/static", "assets")
	routes.ApiRoutes(app)

	database.Connect()
	PORT := os.Getenv("PORT")

	killSig := make(chan os.Signal, 1)

	signal.Notify(killSig, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := app.Start(PORT); err != nil && err != http.ErrServerClosed {
			app.Logger.Fatal(err)
		}
	}()

	<-killSig
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown Failed:", err)
		os.Exit(1)
	}

	log.Println("Server Exited Properly")
}
