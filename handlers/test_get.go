package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func TestGET(c echo.Context) error {
	return c.String(http.StatusOK, "Recieved your test GET")
}
