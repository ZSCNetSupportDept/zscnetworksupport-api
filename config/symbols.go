package config

// our program will use this variable as the config
var UseConfig *Config

type DatabaseConfig struct {
	Type     string `json:"type"`
	Port     int    `json:"port"`
	Path     string `json:"path"`
	User     string `json:"user"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Config struct {
	Port     int            `json:"port"`
	Database DatabaseConfig `json:"database"`
}
