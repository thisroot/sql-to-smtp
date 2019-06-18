package config

import (
	"github.com/go-sql-driver/mysql"
	"net"
	"strconv"
)

type MYSQLConfig struct {
	Scheme   string `json:"scheme"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int64 `json:"port"`
	Database string `json:"db"`
}

func (c *MYSQLConfig) GetMysqlConnectionConfig() *mysql.Config {
	config := &mysql.Config{
		User: c.User,
		Passwd: c.Password,
		DBName: c.Database,
		Net: "tcp",
		Addr: net.JoinHostPort(c.Host, strconv.FormatInt(c.Port, 10) ),
		AllowNativePasswords: true,
	}
	return config
}
