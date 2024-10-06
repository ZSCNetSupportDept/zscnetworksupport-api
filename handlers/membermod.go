// membermod.go 是网维成员管理相关的代码
package handler

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Member struct {
	Usingdb *gorm.DB
}

// GetAll返回当前的所有成员
func (self Member) GetAll(i echo.Context) error {
	return nil
}

// GetTodayAssign返回当天的值班信息
func (self Member) GetTodayAssign(i echo.Context) error {
	return nil
}
