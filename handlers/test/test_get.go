package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (self Test) Get(c echo.Context) error {
	return c.String(http.StatusOK, "Recieved your test GET")
}
