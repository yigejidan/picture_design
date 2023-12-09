package common

import (
	"time"
)

// ServerConfig 服务器配置
type ServerConfig struct {
	RunMode      string
	HttpPort     string
	ClientIp     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	UserName  string
	Password  string
	Host      string
	Port      string
	Database  string
	Charset   string
	ParseTime bool
}

// 定义全局配置
var (
	SvrConfig *ServerConfig
	DbConfig  *DatabaseConfig
)

// SetupSetting 读取配置到全局变量
func SetupSetting() error {
	s, err := NewSetting()
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &DbConfig)
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &SvrConfig)
	if err != nil {
		return err
	}
	return nil
}
