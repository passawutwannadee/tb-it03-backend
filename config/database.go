package config

type Database struct {
	Username string `mapstructure:"USERNAME" validate:"required"`
	Password string `mapstructure:"PASSWORD" validate:"required"`
	Name     string `mapstructure:"NAME" validate:"required"`
	Host     string `mapstructure:"HOST" validate:"required"`
	Port     int    `mapstructure:"PORT" validate:"required"`
}
