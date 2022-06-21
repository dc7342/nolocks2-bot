package telegram

import (
	"fmt"
	tm "github.com/and3rson/telemux/v2"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"nolocks-bot/internal/entity"
)

const (
	cmdTextStart  = "start"
	cmdTextAdd    = "add"
	cmdTextSkip   = "skip"
	cmdTextCancel = "cancel"

	dataStateLocation = "location"
	dataStateComment  = "comment"
)

// HandlerFunction for /start command
func (t *Telegram) startCmd(u *tm.Update) {
	answer := tgbotapi.NewMessage(u.Message.Chat.ID, t.conf.Text.WelcomeMessage)
	_, err := t.bot.Send(answer)
	if err != nil {
		panic(err)
	}
}

func (t *Telegram) addCmd(u *tm.Update) {
	// Next handler awaiting location.
	u.PersistenceContext.SetState(stateLocation)
	msg := tgbotapi.NewMessage(u.Message.Chat.ID, t.conf.Text.AddMessage)
	t.bot.Send(msg)
}

func (t *Telegram) newLocation(u *tm.Update) {
	data := u.PersistenceContext.GetData()
	// Remember location from the message.
	loc := entity.NewLocation(u.Message.Location.Latitude, u.Message.Location.Latitude)
	data[dataStateLocation] = loc
	u.PersistenceContext.SetData(data)
	// Next handler awaiting comment.
	u.PersistenceContext.SetState(stateComment)

	msg := tgbotapi.NewMessage(u.Message.Chat.ID, t.conf.Text.AddComment)
	t.bot.Send(msg)

}

func (t *Telegram) newComment(u *tm.Update) {
	data := u.PersistenceContext.GetData()
	// Get location from previous step.
	loc := data[dataStateLocation].(*entity.Location)
	// Remember comment from the message.
	loc.Comment = u.Message.Text
	data[dataStateLocation] = loc
	// Remember changes.
	u.PersistenceContext.SetData(data)
	// Next handler awaiting photo or skip.
	u.PersistenceContext.SetState(statePhoto)

	msg := tgbotapi.NewMessage(u.Message.Chat.ID, t.conf.Text.AddPhoto)
	t.bot.Send(msg)

}
func (t *Telegram) newPhoto(u *tm.Update) {
	// Get location and comment from the previous steps.
	data := u.PersistenceContext.GetData()
	loc := data[dataStateLocation].(*entity.Location)

	// Clean up context.
	u.PersistenceContext.SetState(stateDefault)
	u.PersistenceContext.ClearData()

	// Get photo from the message.
	photoID := u.Message.Photo[0].FileID
	photoFile, _ := u.Bot.GetFile(tgbotapi.FileConfig{FileID: photoID})
	// Set URL of the Photo from telegram servers.
	loc.ImageUrl = photoFile.Link(t.conf.Telegram.Token)

	// Send to our API.
	if err := t.services.Add(loc); err != nil {
		logError(err)
	}

	msg := tgbotapi.NewMessage(u.Message.Chat.ID, t.conf.Text.Done)
	t.bot.Send(msg)
}

func (t *Telegram) skip(u *tm.Update) {
	// Get location and comment from the previous steps.
	data := u.PersistenceContext.GetData()
	loc := data[dataStateLocation].(*entity.Location)

	// Clean up context.
	u.PersistenceContext.SetState(stateDefault)
	u.PersistenceContext.ClearData()

	// Send to our API.
	if err := t.services.Add(loc); err != nil {
		logError(err)
	}

	msg := tgbotapi.NewMessage(u.Message.Chat.ID, t.conf.Text.Done)
	t.bot.Send(msg)
}

func (t *Telegram) cancel(u *tm.Update) {
	u.PersistenceContext.ClearData()
	u.PersistenceContext.SetState(stateDefault)
	t.bot.Send(tgbotapi.NewMessage(u.Message.Chat.ID, t.conf.Text.Canceled))
}

func (t *Telegram) onError(u *tm.Update, err error, stackTrace string) {
	chat := u.EffectiveChat()
	if chat != nil {
		t.bot.Send(tgbotapi.NewMessage(
			chat.ID,
			fmt.Sprintf("%s: %s", t.conf.Text.OnError, err),
		))
		log.Printf("Warning! An error occurred: %s", err)
	}
}
