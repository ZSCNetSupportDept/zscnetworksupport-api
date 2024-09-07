package use

import (
	"github.com/ZSCNetSupportDept/zscnetworksupport-api/handlers"
	"github.com/labstack/echo/v4"
)

func ConfRouterList(app *echo.Echo) {
	app.File("/", "static/WebsiteIndex.html")
	app.POST("/test", handler.Test)
	app.GET("/test", handler.TestGET)
}
