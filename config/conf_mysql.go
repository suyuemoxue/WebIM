package config

import (
	"fmt"
)

type MySQL struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DB       string `yaml:"db"`
	User     string `yaml:"users"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_level"` //日志等级
}

func (ms *MySQL) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", ms.User, ms.Password, ms.Host, ms.Port, ms.DB)
}
