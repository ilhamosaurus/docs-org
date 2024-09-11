package handler

import (
	"database/sql"
	"errors"
	"strconv"
	"time"

	"go-templ/infra/models"
	"go-templ/infra/service"
	"go-templ/pkg/util"
	"go-templ/pkg/views/components"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CreateDocument(c echo.Context) error {
	user := util.CurrentUser(c)

	code := c.FormValue("code")
	title := c.FormValue("title")
	tag := c.FormValue("tags")
	issuedAtStr := c.FormValue("issued_at")
	dueDateStr := c.FormValue("due_date")
	description := c.FormValue("description")
	issuedAt, err := time.Parse("2006-01-02", issuedAtStr)
	if err != nil {
		return util.Render(c, 400, components.Toast(&components.ToastProps{Error: err, Message: nil}))
	}
	var dueDate *time.Time
	switch {
	case dueDateStr == "":
		dueDate = nil
	default:
		parsedDueDate, err := time.Parse("2006-01-02", dueDateStr)
		if err != nil {
			return util.Render(c, 400, components.Toast(&components.ToastProps{Error: err, Message: nil}))
		}
		dueDate = &parsedDueDate
	}
	document := models.CreateDocumentValidator{
		Code:        code,
		Title:       title,
		Tags:        &tag,
		IssuedAt:    issuedAt,
		DueDate:     dueDate,
		Description: &description,
	}

	if err := c.Validate(document); err != nil {
		var httpErr *echo.HTTPError
		if errors.As(err, &httpErr) {
			return util.Render(c, httpErr.Code, components.Toast(&components.ToastProps{Error: httpErr.Message, Message: nil}))
		}
	}

	dueDateNullTime := sql.NullTime{
		Time:  time.Time{},
		Valid: false,
	}

	if dueDate != nil {
		dueDateNullTime.Time = *dueDate
		dueDateNullTime.Valid = true
	}

	_, err = service.CreateDocument(models.Document{
		Code:        code,
		Title:       title,
		Tags:        &tag,
		IssuedAt:    issuedAt,
		DueDate:     dueDateNullTime,
		Description: &description,
		UserID:      user.ID,
	})
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return util.Render(c, 409, components.Toast(&components.ToastProps{Error: "Document already exists", Message: nil}))
		}

		return util.Render(c, 500, components.Toast(&components.ToastProps{Error: err, Message: nil}))
	}

	return c.Redirect(302, "/dashboard")
}

func GetDocuments(c echo.Context) error {
	user := util.CurrentUser(c)
	limitStr := c.QueryParam("limit")
	offsetStr := c.QueryParam("offset")
	var limit int
	var offset int
	switch {
	case limitStr == "" && offsetStr == "":
		limit = 10
		offset = 1
	default:
		limit, _ = strconv.Atoi(limitStr)
		offset, _ = strconv.Atoi(offsetStr)
	}

	documents, err := service.GetDocuments(user.ID, limit, offset)
	if err != nil {
		return util.Render(c, 500, components.Toast(&components.ToastProps{Error: err, Message: nil}))
	}

	return c.JSON(200, documents)
}

func EditDocument(c echo.Context) error {
	user := util.CurrentUser(c)
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}
	code := c.FormValue("code")
	title := c.FormValue("title")
	tag := c.FormValue("tags")
	issuedAtStr := c.FormValue("issued_at")
	dueDateStr := c.FormValue("due_date")
	description := c.FormValue("description")
	issuedAt, err := time.Parse("2006-01-02", issuedAtStr)
	if err != nil {
		return util.Render(c, 400, components.Toast(&components.ToastProps{Error: err, Message: nil}))
	}
	var dueDate *time.Time
	switch {
	case dueDateStr == "":
		dueDate = nil
	default:
		parsedDueDate, err := time.Parse("2006-01-02", dueDateStr)
		if err != nil {
			return util.Render(c, 400, components.Toast(&components.ToastProps{Error: err, Message: nil}))
		}
		dueDate = &parsedDueDate
	}

	document := models.CreateDocumentValidator{
		Code:        code,
		Title:       title,
		Tags:        &tag,
		IssuedAt:    issuedAt,
		DueDate:     dueDate,
		Description: &description,
	}

	if err = c.Validate(document); err != nil {
		var httpErr *echo.HTTPError
		if errors.As(err, &httpErr) {
			return util.Render(c, httpErr.Code, components.Toast(&components.ToastProps{Error: httpErr.Message, Message: nil}))
		}
	}

	_, err = service.GetDocumentByID(id)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1146 {
			return c.JSON(404, echo.Map{
				"error": "Document not found",
			})
		}

		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	updatedDocument, err := service.UpdateDocument(models.Document{
		ID:          id,
		Code:        code,
		Title:       title,
		Tags:        &tag,
		IssuedAt:    issuedAt,
		DueDate:     sql.NullTime{Time: *dueDate, Valid: dueDate != nil},
		Description: &description,
		UserID:      user.ID,
	})
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return util.Render(c, 409, components.Toast(&components.ToastProps{Error: "Document already exists", Message: nil}))
		}

		return util.Render(c, 500, components.Toast(&components.ToastProps{Error: err, Message: nil}))
	}

	return c.JSON(200, updatedDocument)
}

func DeleteDocument(c echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	_, err = service.GetDocumentByID(id)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1146 {
			return util.Render(c, 404, components.Toast(&components.ToastProps{Error: "Document not found", Message: nil}))
		}

		return util.Render(c, 500, components.Toast(&components.ToastProps{Error: err, Message: nil}))
	}

	err = service.DeleteDocument(id)
	if err != nil {
		return util.Render(c, 500, components.Toast(&components.ToastProps{Error: err, Message: nil}))
	}
	successMsg := "Document deleted successfully"
	return util.Render(c, 200, components.Toast(&components.ToastProps{Error: nil, Message: &successMsg}))
}
