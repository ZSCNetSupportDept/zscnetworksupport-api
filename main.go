package main

import (
	"fmt"

	"github.com/ZSCNetSupportDept/zscnetworksupport-api/config"
	"github.com/ZSCNetSupportDept/zscnetworksupport-api/use"
	"github.com/labstack/echo/v4"

	"strconv"
)

var err error

// sysinit:this function loads the config file stored at project root,and connects to the database
func sysinit() {
	fmt.Println("Start and Load config file:")

	config.LoadConfig("./config.json")
	config.CheckConfig()

	fmt.Printf("use port %d\n", config.UseConfig.Port)
	fmt.Printf("use database in : %s\n", config.UseConfig.Database.Path)
	fmt.Printf("use database user : %s\n", config.UseConfig.Database.User)

}

func main() {
	sysinit()
	app := echo.New()

	use.ConfMWList(app)
	use.ConfRouterList(app)

	app.Logger.Fatal(app.Start(":" + strconv.Itoa(config.UseConfig.Port)))
}
