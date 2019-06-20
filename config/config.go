package config

import (
	"github.com/go-sql-driver/mysql"
	"gopkg.in/gomail.v2"
)

type Configuration struct {
	DB *mysql.Config
	SMTP *gomail.Dialer
}

var Config = &Configuration{}

func init()  {
	var mySqlRawConfig = MYSQLConfig {
		"mysql",
		"devst",
		"c9VRqVdUJ1PLEsLB",
		"212.32.239.1",
		3306,
		"devst",
	}

	var smtConfig = SMTPConfig{
		Host:"5.79.119.5",
		Port: 25,
	}

	Config.DB = mySqlRawConfig.GetMysqlConnectionConfig()
	Config.SMTP = smtConfig.GetSMTPConnectionConfig()
}