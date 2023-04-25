package config

import (
	"fmt"
	"github.com/pelletier/go-toml/v2"
	"os"
)

type DatabaseConnections struct {
	DatabaseConnections []DatabaseConfig `toml:"database_connections"`
}

type DatabaseConfig struct {
	BusinessId string `toml:"businessId"`
	IP         string `toml:"ip"`
	Port       string `toml:"port"`
	Username   string `toml:"username"`
	Password   string `toml:"password"`
	DBName     string `toml:"dbname"`
	TableName  string `toml:"table_name"`
}

func LoadConfig(configPath string) (*DatabaseConnections, error) {
	// 打开配置文件
	file, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(file))
	// 解析配置文件
	cfg := DatabaseConnections{}
	err = toml.Unmarshal(file, &cfg)
	if err != nil {
		return nil, err
	}

	fmt.Println("cfg.ImageRoute", cfg)

	return &cfg, nil
}
