package handler

import (
	"fmt"
	"net/http"

	"github.com/ZSCNetSupportDept/zscnetworksupport-api/databases"
	"github.com/labstack/echo/v4"
)

func (self Test) TestSQLite(c echo.Context) error {
	co, err := database.ConnectSQLite()
	if err != nil {
		fmt.Println(fmt.Errorf("error In TestSQLite(): %v", err))
		return c.String(http.StatusInternalServerError, "error in database connection")
	}
	co.AutoMigrate(&database.Member{}, &database.User{}, &database.Ticket{})
	test := database.Member{ID: 10001, Name: "测试", Access: 10, Phone: 11111111111}
	result := co.Create(&test)

	if result.Error != nil {
		fmt.Println(fmt.Errorf("error In TestSQLite(): %v", result.Error))
		defer co.Delete(&test)
		return c.String(http.StatusInternalServerError, "error in database operation")
	}
	defer co.Delete(&test)
	var receive database.Member
	co.First(&receive, 10001)
	return c.JSON(http.StatusOK, receive)
	//return 0
}
