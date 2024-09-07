package main

import (
	"fmt"
	"github.com/ZSCNetSupportDept/zscnetworksupport-api/config"
	"github.com/ZSCNetSupportDept/zscnetworksupport-api/use"
	"github.com/labstack/echo/v4"
	"os"

	"strconv"
)

var err error

// sysinit:this function loads the config file stored at project root,and connects to the database
func sysinit() {
	if len(os.Args) != 2 {
		panic("we accept one and only one argument as the config file path")
	}
	fmt.Println("Start and Load config file:")

	config.LoadConfig(os.Args[1])
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
