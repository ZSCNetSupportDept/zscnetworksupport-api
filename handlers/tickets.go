package handler

import (
	"gorm.io/gorm"
)

type ticket struct {
	usingdb *gorm.DB
}
