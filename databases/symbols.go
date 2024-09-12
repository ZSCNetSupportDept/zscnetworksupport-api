package database

import (
	"gorm.io/gorm"
	"time"
)

// 网维成员的记录
type Member struct {
	ID     int `gorm:"PrimaryKey"` //工号
	Name   string
	Access int //详见文档
	Phone  int
}

// 普通用户的记录
type User struct {
	ID    int `gorm:"PrimaryKey"` //学号
	Name  string
	ISP   int //详见文档
	Block int //详见文档
	Room  int
	Phone int
}

// 工单的记录
type Ticket struct {
	gorm.Model        //定义了编号，创建，更新和删除的时间
	Userid     int    //报修人的学号
	Status     int    //详见文档
	notes      string //描述
}

// 新生报名的记录
type Volunteer struct {
	ID      int    `gorm:"PrimaryKey" json:"ID"` //学号
	Name    string `json:"Name"`
	Phone   int    `json:"Phone"`
	Major   int    `json:"Major"`   //详见文档
	College int    `json:"College"` //看文档
}

// 工单修改的记录
type TicketTrace struct {
	Issue        *Ticket   //该记录所指向的工单
	Operator     *Member   //修改操作人
	StatusChange int       //变更为
	CreatedAt    time.Time //创建的时间
	Notes        string    //修改备注
}

var Usingdb *gorm.DB //handler所使用的主数据库连接
