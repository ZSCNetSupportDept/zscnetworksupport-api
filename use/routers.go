package use

import (
	"github.com/labstack/echo/v4"
)

func ConfRouterList(app *echo.Echo) {
	app.File("/", "static/WebsiteIndex.html")
}
