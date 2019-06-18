package config

import "github.com/go-sql-driver/mysql"

type Configuration struct {
	DB *mysql.Config
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

	Config.DB = mySqlRawConfig.GetMysqlConnectionConfig()
}