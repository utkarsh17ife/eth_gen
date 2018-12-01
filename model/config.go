package model

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AbiPath string `json:"abiPath"`
}

func LoanConfig() *Config {

	config := &Config{}
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicf("Fatal error config file: %s \n", err)
	}
	viper.Unmarshal(config)
	return config

}
