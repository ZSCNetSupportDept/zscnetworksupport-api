package use

import (
	"github.com/ZSCNetSupportDept/zscnetworksupport-api/databases"
	"github.com/ZSCNetSupportDept/zscnetworksupport-api/handlers"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ConfRouterList(app *echo.Echo) {

	// initialize the handler
	zx, gd := initHandlers(database.Usingdb)

	//static pages
	app.File("/", "static/WebsiteIndex.html") //index page

	//Recruitment API
	app.POST("/recruitment/FirstRequest", zx.FirstRequest)
	app.POST("/recruitment/ChangeRequest", zx.ChangeRequest)
	app.GET("/recruitment/ViewRequest", zx.ViewRequestUser)
	app.POST("/recruitment/CancelRequest", zx.CancelUserRequest)

	app.GET("/admin/recruitment/ViewAllRequest", zx.ViewRequestAdmin)
	app.POST("/admin/recruitment/CancelRequest", zx.CancelRequestAdmin)

}

func initHandlers(usingdb *gorm.DB) (*handler.Recruit, *handler.Ticket) {
	// initialize the handler symbol
	recruitment, err := handler.NewRecruitHandler(database.Usingdb)
	if err != nil {
		panic("error in initializing recruitment handler symbol")
	}
	tickets, err := handler.NewTicketHandler(database.Usingdb)
	if err != nil {
		panic("error in initializing tickets handler symbol")
	}
	return recruitment, tickets
}
