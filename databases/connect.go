package database

import (
	"github.com/ZSCNetSupportDept/zscnetworksupport-api/config"
	//"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectMySQL() {

}
func ConnectSQLite() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(config.UseConfig.Database.Path), &gorm.Config{})
	return db, err
}
