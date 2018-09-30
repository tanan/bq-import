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
	Host   string
	Port   int
	Schema string
	User   string
	Pass   string
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
