package database

import (
	"github.com/ZSCNetSupportDept/zscnetworksupport-api/config"
	//"github.com/labstack/echo/v4"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectMySQL() (*gorm.DB, error) {
	i := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local\n", config.UseConfig.Database.User, config.UseConfig.Database.Password, config.UseConfig.Database.Path, config.UseConfig.Database.Port, config.UseConfig.Database.Name)
	db, err := gorm.Open(mysql.Open(i), &gorm.Config{})
	return db, err
}
func ConnectSQLite() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(config.UseConfig.Database.Path), &gorm.Config{})
	return db, err
}

func ConnectPGSQL() (*gorm.DB, error) {
	i := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", config.UseConfig.Database.Path, config.UseConfig.Database.User, config.UseConfig.Database.Password, config.UseConfig.Database.Name, config.UseConfig.Database.Port)
	db, err := gorm.Open(postgres.Open(i), &gorm.Config{})
	return db, err
}

// for those handlers that use a standalone database connection
func Connect() (db *gorm.DB, err error) {
	switch config.UseConfig.Database.Type {
	case "MySQL":
		db, err := ConnectMySQL()
		if err != nil {
			return Usingdb, nil
		}
		return db, err
	case "PostgreSQL":
		db, err := ConnectPGSQL()
		if err != nil {
			return Usingdb, nil
		}
		return db, err
	case "SQLite":
		db, err := ConnectSQLite()
		if err != nil {
			return Usingdb, nil
		}
		return db, err
	default:
		return Usingdb, nil
	}

}
