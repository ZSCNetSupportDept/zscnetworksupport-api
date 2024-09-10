package database

import (
	"fmt"
	"github.com/ZSCNetSupportDept/zscnetworksupport-api/config"
)

// try to establish a main connection that handlers will use
func init() {
	switch config.UseConfig.Database.Type {
	case "MySQL":
		db, err := ConnectMySQL()
		if err != nil {
			fmt.Println(fmt.Errorf("error in initializing database config: %v", err))
			panic("panic at init()")
		}
		Usingdb = db
	case "PostgreSQL":
		db, err := ConnectPGSQL()
		if err != nil {
			fmt.Println(fmt.Errorf("error in initializing database config: %v", err))
			panic("panic at init()")
		}
		Usingdb = db
	case "SQLite":
		db, err := ConnectSQLite()
		if err != nil {
			fmt.Println(fmt.Errorf("error in initializing database config: %v", err))
			panic("panic at init()")
		}
		Usingdb = db
	default:
		panic("Please enter a valid database Type field")
	}

}
