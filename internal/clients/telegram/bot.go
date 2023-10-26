package telegram

import (
	"fmt"

	"github.com/go-bot/internal/handler"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot     *tgbotapi.BotAPI
	handler handler.Handler
}

func NewBot(tgString string, handler *handler.Handler) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(tgString)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize bot API: %w", err)
	}
	return &Bot{
		bot:     bot,
		handler: *handler,
	}, nil
}

func (b *Bot) Run() error {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	b.bot.Debug = true
	updates := b.bot.GetUpdatesChan(updateConfig)

	for update := range updates {

		if update.CallbackQuery != nil {
			go func(callback *tgbotapi.CallbackQuery) {

				responseMessage, err := b.handler.HandleCallback(callback)
				if err != nil {
					fmt.Println("Failed to handle callback:", err)
				}

				responseMessage.ChatID = callback.From.ID
				responseMessage.MessageID = callback.Message.MessageID
				responseMessage.ParseMode = "html"
				responseMessage.DisableWebPagePreview = true

				if responseMessage.Text != "" {
					_, err := b.bot.Send(responseMessage)
					if err != nil {
						fmt.Println("Failed to send callback:", err)
					}
					return
				}

			}(update.CallbackQuery)
		}

		if update.Message != nil {
			go func(message *tgbotapi.Message) {

				responseMessage, err := b.handler.HandleMessage(message)
				if err != nil {
					fmt.Println("Failed to handle message:", err)
				}

				responseMessage.ChatID = message.From.ID
				responseMessage.DisableWebPagePreview = true

				if responseMessage.Text != "" {
					_, err := b.bot.Send(responseMessage)
					if err != nil {
						fmt.Println("Failed to send message:", err)
					}
					return
				}

			}(update.Message)
		}
	}

	return nil
}
