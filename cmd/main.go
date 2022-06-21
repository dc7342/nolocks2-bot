package main

import (
	"fmt"
	"nolocks-bot/internal/client"
	"nolocks-bot/internal/service"
	"nolocks-bot/internal/telegram"
	"nolocks-bot/pkg/config"
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
