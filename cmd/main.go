package main

import (
	"fmt"
	"github.com/dc7342/nolocks2/internal/client"
	"github.com/dc7342/nolocks2/internal/service"
	"github.com/dc7342/nolocks2/internal/telegram"
	"github.com/dc7342/nolocks2/pkg/config"
)

func main() {
	c, err := config.Read()
	if err != nil {
		panic(fmt.Errorf("fatal error during reading config file: %w", err))
	}

	httpClient := client.NewClient(c.NoLocks)
	if err != nil {
		panic(err)
	}

	serv := service.NewService(httpClient)
	bot := telegram.NewTelegram(serv, c)
	if err := bot.Start(); err != nil {
		panic(err)
	}
}
