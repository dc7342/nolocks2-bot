package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Telegram struct {
		Token   string
		Timeout int
	}

	Text struct {
		WelcomeMessage string
		UnknownMessage string
		AddMessage     string
		AddComment     string
		AddPhoto       string
		Done           string
		OnError        string
		Canceled       string
	}

	NoLocks struct {
		EndpointURL string
		User        string
		Pass        string
	}
}

func Read() (Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/nolocks/")
	viper.AddConfigPath("$HOME/.nolocks/")
	viper.AddConfigPath("/")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}
	var c Config

	err = viper.Unmarshal(&c)
	if err != nil {
		return Config{}, err
	}

	return c, nil
}
