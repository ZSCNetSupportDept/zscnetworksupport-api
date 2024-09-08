package database

import (
	"gorm.io/gorm"
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
	Id    int `gorm:"PrimaryKey"` //学号
	Name  string
	ISP   int //详见文档
	Block int //详见文档
	Room  int
	Phone int
}

// 工单的记录
type Ticket struct {
	gorm.Model     //定义了编号，创建，更新和删除的时间
	Userid     int //报修人的学号
	Status     int //详见文档
}

// 新生报名的记录
type Volunteer struct {
	ID      int `gorm:"PrimaryKey"` //学号
	Name    string
	Phone   int
	Major   int //详见文档
	College int //看文档
}
