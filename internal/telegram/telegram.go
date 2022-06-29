package telegram

import (
	tm "github.com/and3rson/telemux/v2"
	"github.com/dc7342/nolocks2/internal/service"
	"github.com/dc7342/nolocks2/pkg/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Telegram struct {
	bot      *tgbotapi.BotAPI
	services *service.Service
	conf     config.Config

	cmds *tm.Mux
}

func NewTelegram(serv *service.Service, conf config.Config) *Telegram {
	return &Telegram{
		services: serv,
		conf:     conf,
	}
}

func (t *Telegram) Start() error {
	bot, err := tgbotapi.NewBotAPI(t.conf.Telegram.Token)
	if err != nil {
		return err
	}
	t.bot = bot

	u := tgbotapi.NewUpdate(0)
	u.Timeout = t.conf.Telegram.Timeout
	updates := t.bot.GetUpdatesChan(u)

	t.initMux()

	for update := range updates {
		t.cmds.Dispatch(bot, update)
	}

	return nil
}
