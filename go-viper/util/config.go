package util

import (
	"fmt"

	"github.com/spf13/viper"
)

type DbConfig struct {
	Address  string `mapstructure:"address"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	User     string `mapstructure:"user"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type Config struct {
	Db     DbConfig     `mapstructure:"db"`
	Server ServerConfig `mapstructure:"server"`
}

var vp *viper.Viper

func LoadConfig() (config *Config, err error) {
	vp = viper.New()

	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath("./util")
	vp.AddConfigPath(".")

	err = vp.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	err = vp.Unmarshal(&config)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return config, nil

}
