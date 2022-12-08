package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config interface {
	Save() error
	Load() error
}

var configModels []Config

func Register(config Config) {
	configModels = append(configModels, config)
}

func Initialize() error {
	if err := initViper(); err != nil {
		log.Println("Could not init Viper!")
		return err
	}
	return nil
}

func initViper() error {
	viper.SetConfigName("config/config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Println("Could not accesee config file ", err)
		return err
	}
	err := readConfig()
	if err != nil {
		log.Println("Could not read config file, ", err)
		return err
	}
	return nil
}

func readConfig() error {
	for _, model := range configModels {
		err := model.Load()
		if err != nil {
			return err
		}
	}
	return nil
}
