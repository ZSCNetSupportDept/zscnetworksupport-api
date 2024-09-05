package use

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ConfMWList(app *echo.Echo) {
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
}
