package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Bot struct {
	Bot *tgbotapi.BotAPI
}

func New(token string) *Bot {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil
	}

	return &Bot{Bot: bot}
}
