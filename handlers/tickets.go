package handler

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// 所有工单handler作为该对象的方法实现
type Ticket struct {
	usingdb *gorm.DB
}

/* 用户API */

// SubmitTicket 使用户提交一个工单
func (self Ticket) SubmitTicket(i echo.Context) error { return nil }

// CancelTicket 使用户取消自己所报修的工单
func (self Ticket) CancelTicket(i echo.Context) error { return nil }

// Bind 要求用户提供自己的学号，与报修信息，绑定数据库中已导入的学号（用户提供姓名和学号，API将这些信息与数据库中存储的信息进行比对，若成功则存储进用户提供的宿舍号，楼栋，网络账号和手机号，运营商等信息）
func (self Ticket) Bind(i echo.Context) error { return nil }

// ChangeBind 修改用户的绑定信息
func (self Ticket) ChangeBind(i echo.Context) error { return nil }

// CheckBind 显示用户的信息
func (self Ticket) CheckBind(i echo.Context) error { return nil }

// ViewTicket 查看用户的工单报修历史
func (self Ticket) ViewTicket(i echo.Context) error { return nil }

/* 管理API */

// GetTicketByUserBlock 返回一个片区所拥有的全部工单
func (self Ticket) GetTicketByUserBlock(i echo.Context) error { return nil }

// GetTicketDetail 返回工单的详情
func (self Ticket) GetTicketDetail(i echo.Context) error { return nil }
