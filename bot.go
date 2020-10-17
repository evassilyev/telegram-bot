package tgbot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	token   string
	chat    int64
	bot     *tgbotapi.BotAPI
	Updates tgbotapi.UpdatesChannel
}

func NewBot(token string, chat int64) (*Bot, error) {

	b, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.GetUpdatesChan(u)

	return &Bot{
		token:   token,
		chat:    chat,
		bot:     b,
		Updates: updates,
	}, err
}
func (b *Bot) SendMessage(message string) error {
	msg := tgbotapi.NewMessage(b.chat, message)
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) ReplyToMessage(msg *tgbotapi.Message, text string) error {
	rm := tgbotapi.NewMessage(msg.Chat.ID, text)
	rm.ReplyToMessageID = msg.MessageID
	_, err := b.bot.Send(rm)
	return err
}
