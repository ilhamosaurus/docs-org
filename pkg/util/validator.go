package util

import (
	"fmt"
	"reflect"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		errMsgs := make([]string, 0)
		for _, e := range err.(validator.ValidationErrors) {
			errMsgs = append(errMsgs, fmt.Sprintf("%s: %s %s", e.Field(), e.Tag(), e.Param()))
		}
		return echo.NewHTTPError(400, errMsgs)
	}

	return nil
}

func NewCustomValidator() *CustomValidator {
	v := validator.New()
	v.RegisterCustomTypeFunc(func(field reflect.Value) interface{} {
		if val, ok := field.Interface().(time.Time); ok {
			return val.Format(time.RFC3339)
		}
		return nil
	}, time.Time{})
	return &CustomValidator{validator: v}
}
