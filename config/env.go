package config

import (
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

func LoadConfiguration() {
	if len(os.Getenv("ENVNAME")) == 0 {
		panic("Please specify .env file name")
	}
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

func GetPort() string {
	return configuration.PORT
}
