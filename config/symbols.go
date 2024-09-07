package config

// our program will use this variable as the config
var UseConfig *Config

type DatabaseConfig struct {
	Path     string `json:"path"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type Config struct {
	Port     int            `json:"port"`
	Database DatabaseConfig `json:"database"`
}
