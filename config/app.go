package config

type App struct {
	HTTPPort string `mapstructure:"HTTP_PORT" validate:"required"`
}
