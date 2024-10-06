package handler

import (
	"github.com/ZSCNetSupportDept/zscnetworksupport-api/databases"
	// "github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

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
