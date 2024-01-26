package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	// DBHost         string `mapstructure:"POSTGRES_HOST"`
	// DBUserName     string `mapstructure:"POSTGRES_USER"`
	// DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	// DBName         string `mapstructure:"POSTGRES_DB"`
	// DBPort         string `mapstructure:"POSTGRES_PORT"`
	// ServerPort     string `mapstructure:"PORT"`

	// ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`

	AccessTokenSecret    string        `mapstructure:"ACCESS_TOKEN_SECRET"`
	AccessTokenExpiresIn time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRED_IN"`
	AccessTokenMaxAge    int           `mapstructure:"ACCESS_TOKEN_MAXAGE"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
