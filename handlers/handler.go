package handler

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Test struct {
	Usingdb *gorm.DB
}

type Recruit struct {
	Usingdb *gorm.DB
}

type Ticket struct {
	Usingdb *gorm.DB
}

type Handler interface {
	IsHandler()
}

func (self Test) IsHandler() bool {
	return true
}
