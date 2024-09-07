package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// LoadConfig:this function will load a config file ,decode the file and store the config in the config.UseConfig
func LoadConfig(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(fmt.Errorf("error opening config file: %v", err))
		panic("panic at LoadConfig")
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		fmt.Println(fmt.Errorf("error opening config file: %v", err))
		panic("panic at Loadconfig")
	}
	*UseConfig = config
}
func CheckConfig() {
	//here to check config validity
	if UseConfig.Port < 1024 && UseConfig.Port > 65535 {
		panic("at CheckConfig:please enter a valid port to serve(1024<port<65535)")
	}
}
