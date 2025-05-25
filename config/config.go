package config

import (
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

const (
	AppDebug = "debug"
	AppTest  = "test"
	AppProd  = "prod"
)

type Config struct {
	AppMode  string   `mapstructure:"APP_MODE" validate:"required"`
	App      App      `mapstructure:"APP" validate:"required"`
	Database Database `mapstructure:"DATABASE" validate:"required"`
}

var C Config

func Init(path string) {
	readInConfig(path)
	setConfig()
}

func readInConfig(path string) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}

func setConfig() {
	err := viper.Unmarshal(&C)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	v := validator.New()
	if err := v.Struct(&C); err != nil {
		log.Fatalf("config validation error, %v", err)
	}
}
