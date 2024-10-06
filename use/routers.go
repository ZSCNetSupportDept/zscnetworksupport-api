package use

import (
	"github.com/ZSCNetSupportDept/zscnetworksupport-api/databases"
	//"github.com/ZSCNetSupportDept/zscnetworksupport-api/handlers"
	"github.com/labstack/echo/v4"
	//"gorm.io/gorm"
	"github.com/ZSCNetSupportDept/zscnetworksupport-api/handlers"
)

func ConfRouterList(app *echo.Echo) {

	zx, gd, cs := handler.InitHandlers(database.Usingdb) // initialize the handler
	_, _, _ = zx, gd, cs                                 // you will not get an error if a handler is not used

	//static pages
	app.File("/", "static/WebsiteIndex.html") //index page

	//test API
	app.GET("/test/", cs.GetTestIndex)

	//Recruitment API
	app.POST("/recruitment/FirstRequest", zx.FirstRequest)
	app.POST("/recruitment/ChangeRequest", zx.ChangeRequest)
	app.GET("/recruitment/ViewRequest", zx.ViewRequestUser)
	app.POST("/recruitment/CancelRequest", zx.CancelUserRequest)

	app.GET("/admin/recruitment/ViewAllRequest", zx.ViewRequestAdmin)
	app.POST("/admin/recruitment/CancelRequest", zx.CancelRequestAdmin)

}
