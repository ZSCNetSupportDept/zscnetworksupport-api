package model

import (
	"gorm.io/gorm"
	"time"
)

// Ticket是一个工单的对象模型，其中：
// gorm.Model是gorm框架内置的一个字段，它定义了工单的编号，创建时间，更新时间，删除时间（工单删除时不会真的删除，而是会标记上），
// Userid是报修人的学号，根据这个学号将报修人的信息与工单信息组合起来，
// Status是工单的状态，0是刚报修，1是已推迟，2是已上报，3是已解决，9是已取消
// notes是工单情况的描述，用户填写的
type Ticket struct {
	gorm.Model
	Userid int
	Status int
	Notes  string
}

// TicketTrace是一个工单修改的对象模型，工单状态的更改都对应了一个对象修改，其中：
// Issue是该记录对应工单的编号
// Operator是修改操作人的工号
// StatusChange是该工单变更后状态的代码
// CreatedAt是该修改创建的时间
// Notes是该工单情况的备注描述
type TicketTrace struct {
	Issue        uint
	Operator     int
	StatusChange int
	CreatedAt    time.Time
	Notes        string
}

// Create 将创建一个工单
func (self *Ticket) Create(userid int, notes string) error

// GetByID通过用户学号来查询所有报修
func (self *Ticket) GetByID(userid int) error

// GetByBlock通过楼号来查询所有报修
func (self *Ticket) GetByBlock(block int) error

// Trace添加一条处理记录
func (self *Ticket) Trace(issue uint, operator int, statusChange int, notes string) error

// Cancel取消报修
func (self *Ticket) Cancel(id uint) error
