package handler

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type Test struct {
	Usingdb *gorm.DB
}

func (self *Test) GetTestIndex(i echo.Context) error {
	return i.String(http.StatusOK, "Received a HTTP GET request!")
}
