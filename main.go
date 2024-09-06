package main

import (
	"encoding/json"
	"fmt"

	"github.com/ZSCNetSupportDept/zscnetworksupport-api/use"
	"github.com/labstack/echo/v4"

	"os"
	"strconv"
)

var UseConfig *Config

// sysinit:this function loads the config file stored at project root,and connects to the database
func sysinit() *Config {
	fmt.Println("Start and Load config file:")
	useConfig, _ := LoadConfig("./config.json")
	fmt.Printf("use port %d\n", useConfig.Port)
	fmt.Printf("use database in : %s\n", useConfig.Database.Path)
	fmt.Printf("use database user : %s\n", useConfig.Database.User)
	return useConfig
}

func main() {
	UseConfig = sysinit()
	app := echo.New()

	use.ConfMWList(app)
	use.ConfRouterList(app)

	app.Logger.Fatal(app.Start(":" + strconv.Itoa(UseConfig.Port)))
}

func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening config file: %v", err)
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("error decoding config file: %v", err)
	}

	return &config, nil
}

type DatabaseConfig struct {
	Path     string `json:"path"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type Config struct {
	Port     int            `json:"port"`
	Database DatabaseConfig `json:"database"`
}
