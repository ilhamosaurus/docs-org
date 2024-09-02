package domain

import (
	"go-templ/infra/service"
	"go-templ/pkg/util"
	"go-templ/pkg/views"
	"go-templ/pkg/views/components"

	"github.com/labstack/echo/v4"
)

func GetHome(c echo.Context) error {
	auth, _ := c.Cookie("Authorization")
	if auth != nil {
		return c.Redirect(302, "/dashboard")
	}
	return util.Render(c, 200, views.Index(&components.ToastProps{Error: nil, Message: nil}))
}

func LoginPage(c echo.Context) error {
	return util.Render(c, 200, components.LoginForm())
}

func RegisterPage(c echo.Context) error {
	return util.Render(c, 200, components.RegisterForm())
}

func Dashboard(c echo.Context) error {
	user := util.CurrentUser(c)
	documents, err := service.GetDocuments(user.ID, 10, 1)
	if err != nil {
		return util.Render(c, 500, components.Toast(&components.ToastProps{Error: err, Message: nil}))
	}
	return util.Render(c, 200, views.Dashboard(*documents, &components.ToastProps{Error: nil, Message: nil}))
}

func CreateDocument(c echo.Context) error {
	return util.Render(c, 200, components.DocsForm())
}
