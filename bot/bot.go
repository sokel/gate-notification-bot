package bot

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"time"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot() (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &Bot{
		bot: bot,
	}, nil
}

func (b *Bot) Start() error {
	tk := time.NewTicker(reportPeriod)
	for {
		select {
		case <-tk.C:
			msg := tgbotapi.NewMessageToChannel(channelName, "Hi")
			_, err := b.bot.Send(msg)
			if err != nil {
				return fmt.Errorf("failed to send message: %v", err)
			}
		}
	}
}

