package config

import (
	"github.com/BurntSushi/toml"
	"strconv"
)

type Config struct {
	Application AppConfig `toml:"application"`
	Database    DBConfig  `toml:"database"`
}

type AppConfig struct {
	SecretFile string `toml:"secret_file"`
}

type DBConfig struct {
	SQLType string `toml:"sql_type"`
	Host    string `toml:"host"`
	Port    int    `toml:"port"`
	Schema  string `toml:"schema"`
	User    string `toml:"user"`
	Pass    string `toml:"pass"`
}

func (config DBConfig) GetConnInfo() string {
	var connstr string
	connstr += "user=" + config.User
	connstr += " "
	connstr += "password=" + config.Pass
	connstr += " "
	connstr += "host=" + config.Host
	connstr += " "
	connstr += "port=" + strconv.Itoa(config.Port)
	connstr += " "
	connstr += "dbname=" + config.Schema
	connstr += " "
	connstr += "sslmode=disable"
	return connstr
}

func LoadFile(path string) (*Config, error) {
	var config Config
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		return nil, err
	}
	return &config, err
}
