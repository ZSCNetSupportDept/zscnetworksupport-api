package use

import (
	"github.com/ZSCNetSupportDept/zscnetworksupport-api/handlers/test"
	"github.com/labstack/echo/v4"
)

func ConfRouterList(app *echo.Echo) {
	app.File("/", "static/WebsiteIndex.html") //index page

	//test API
	app.POST("/test", handler.Test)
	app.GET("/test", handler.TestGET)
}
