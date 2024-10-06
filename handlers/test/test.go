package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (self Test) Test(c echo.Context) error {
	return c.String(http.StatusFound, "Recieved your test POST")
}

type Test struct {
	Usingdb *gorm.DB
}
