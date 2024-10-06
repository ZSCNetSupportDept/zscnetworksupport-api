package handler

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// 所有招新handler作为该对象的方法实现
type Recruit struct {
	usingdb *gorm.DB
}

/* 用户API */

// FirstRequest初次提交报名请求到数据库
func (self Recruit) FirstRequest(i echo.Context) error

// ChangeRequest使用户修改自己的报名请求
func (self Recruit) ChangeRequest(i echo.Context) error

// ViewRequestUser使用户查看自己的报名请求
func (self Recruit) ViewRequestUser(i echo.Context) error

// CancelUserRequest取消用户自己的报名请求
func (self Recruit) CancelUserRequest(i echo.Context) error

/* 管理员API */

// ViewRequestAdmin使管理员查看全部报名
func (self Recruit) ViewRequestAdmin(i echo.Context) error

// CancelRequestAdmin使管理员取消某个报名
func (self Recruit) CancelRequestAdmin(i echo.Context) error

// AdmitVolunteer是管理员将指定的报名人员加入网维实习成员
func (self Recruit) AdmitVolunteer(i echo.Context) error
