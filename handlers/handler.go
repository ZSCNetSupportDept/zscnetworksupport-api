package handler

import (
	"github.com/ZSCNetSupportDept/zscnetworksupport-api/databases"
	// "github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitHandlers(usingdb *gorm.DB) (*Recruit, *Ticket, *Test) {
	// initialize the handler symbol
	recruitment, err := NewRecruitHandler(database.Usingdb)
	if err != nil {
		panic("error in initializing recruitment handler symbol")
	}
	tickets, err := NewTicketHandler(database.Usingdb)
	if err != nil {
		panic("error in initializing tickets handler symbol")
	}
	test, err := NewTestHandler(database.Usingdb)
	if err != nil {
		panic("error in initializing test handler symbol")
	}
	return recruitment, tickets, test
}

// NewRecruitHandler新建一个招新handler的实例
func NewRecruitHandler(usingdb *gorm.DB) (handler *Recruit, err error) {
	zx := new(Recruit)
	zx.usingdb = database.Usingdb
	return zx, nil
}

// NewTicketHandler新建一个工单handler的实例
func NewTicketHandler(usingdb *gorm.DB) (handler *Ticket, err error) {
	gd := new(Ticket)
	gd.usingdb = database.Usingdb
	return gd, nil

}

func NewTestHandler(usingdb *gorm.DB) (handler *Test, err error) {
	cs := new(Test)
	cs.Usingdb = database.Usingdb
	return cs, nil
}
