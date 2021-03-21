package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Configurations struct {
	PORT         string	
	SQL_HOST     string
	SQL_PORT     string
	SQL_USERNAME string
	SQL_PASSWORD string
	SQL_DATABASE string
}

var configuration Configurations

func LoadConfiguration()  {
	viper.SetConfigName(os.Getenv("ENVNAME"))
	viper.AddConfigPath("../")
	viper.SetConfigType("env")
	
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		panic(err)
	}
}

func GetSqlConn() string {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		configuration.SQL_USERNAME,
		configuration.SQL_PASSWORD,
		configuration.SQL_HOST,
		configuration.SQL_PORT,
		configuration.SQL_DATABASE,
	)
	return connStr
}

func GetPort() string {
	return configuration.PORT
}
