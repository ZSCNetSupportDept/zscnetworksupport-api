package handler

import (
	"github.com/ZSCNetSupportDept/zscnetworksupport-api/databases"
	"github.com/labstack/echo/v4"
	"net/http"
)

// FirstRequest:该API负责提交用户的初次报名请求到数据库，前端需要使用POST发送一段json信息，看文档
func FirstRequest(i echo.Context) error {
	buffer := database.Volunteer{}
	err := i.Bind(&buffer)
	if err != nil {
		return i.String(http.StatusBadRequest, "Bad Request")
	}
	result := database.Usingdb.Create(&buffer)
	if result.Error != nil {
		return i.String(http.StatusInternalServerError, "server error")
	}
	return i.String(http.StatusOK, "OK")
}
