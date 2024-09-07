package handler

import (
	"errors"
	"net/http"
	"time"

	"go-templ/infra/models"
	"go-templ/infra/service"
	"go-templ/pkg/util"
	"go-templ/pkg/views/components"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	var user models.LoginRequest

	email := c.FormValue("email")
	password := c.FormValue("password")
	user.Email = email
	user.Password = password

	if err := c.Validate(&user); err != nil {
		var httpErr *echo.HTTPError
		if errors.As(err, &httpErr) {
			return util.Render(c, httpErr.Code, components.Toast(&components.ToastProps{Error: httpErr.Message, Message: nil}))
		}
	}

	token, err := service.Login(user.Email, user.Password)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1045 || err.Error() == "invalid password" {
			return util.Render(c, 400, components.Toast(&components.ToastProps{Error: "Invalid Credentials", Message: nil}))
		}

		return util.Render(c, 500, components.Toast(&components.ToastProps{Error: err, Message: nil}))
	}
	cookie := new(http.Cookie)
	cookie.Name = "Authorization"
	cookie.Value = *token
	cookie.Expires = time.Now().Add(time.Minute * 30)
	cookie.Path = "/"
	cookie.Secure = false
	cookie.SameSite = http.SameSiteLaxMode
	cookie.HttpOnly = true
	c.SetCookie(cookie)
	return c.Redirect(http.StatusFound, "/dashboard")
}

func Register(c echo.Context) error {
	var user models.RegisterRequest

	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")
	user.Name = name
	user.Email = email
	user.Password = password

	if err := c.Validate(&user); err != nil {
		var httpErr *echo.HTTPError
		if errors.As(err, &httpErr) {
			return util.Render(c, httpErr.Code, components.Toast(&components.ToastProps{Error: httpErr.Message, Message: nil}))
		}
	}

	err := service.CreateUser(user)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return util.Render(c, 409, components.Toast(&components.ToastProps{Error: "Email already exists", Message: nil}))
		}
		return util.Render(c, 500, components.Toast(&components.ToastProps{Error: err, Message: nil}))
	}
	message := "User created successfully"
	return util.Render(c, 201, components.Toast(&components.ToastProps{Error: nil, Message: &message}))
}

func Logout(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "Authorization"
	cookie.Value = ""
	cookie.Expires = time.Unix(0, 0)
	cookie.Path = "/"
	cookie.Secure = false
	cookie.SameSite = http.SameSiteLaxMode
	cookie.HttpOnly = true
	c.SetCookie(cookie)
	return c.Redirect(http.StatusFound, "/")
}
