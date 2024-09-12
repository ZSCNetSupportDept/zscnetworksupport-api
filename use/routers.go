package use

import (
	"github.com/ZSCNetSupportDept/zscnetworksupport-api/handlers/recruitment"
	"github.com/ZSCNetSupportDept/zscnetworksupport-api/handlers/test"
	"github.com/labstack/echo/v4"
)

func ConfRouterList(app *echo.Echo) {
	app.File("/", "static/WebsiteIndex.html") //index page

	//test API
	app.POST("/test", handler.Test)
	app.GET("/test", handler.TestGET)
	app.GET("/testdb", handler.TestSQLite)

	//recruitment businesses
	app.POST("/recruitment/first_request", handler.FirstRequest)
}
