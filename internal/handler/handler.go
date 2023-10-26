package handler

import (
	"fmt"
	"strconv"

	"github.com/go-bot/internal/service"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) HandleMessage(message *tgbotapi.Message) (*tgbotapi.MessageConfig, error) {
	switch message.Text {
	case "/start":
		h.services.UserService.AddUser(message.From.ID)
		msg := &tgbotapi.MessageConfig{Text: ""}
		msg.ReplyMarkup = h.services.GetCategoriesInlineKeyboard()
		fmt.Print(message)
		return msg, nil
	case "/menu":
		msg := &tgbotapi.MessageConfig{Text: ""}
		msg.ReplyMarkup = h.services.GetCategoriesInlineKeyboard()
		return msg, nil
	default:
		msg := &tgbotapi.MessageConfig{Text: "Извините, я вас не понимаю"}
		return msg, nil
	}
}

func (h *Handler) HandleCallback(callback *tgbotapi.CallbackQuery) (*tgbotapi.EditMessageTextConfig, error) {

	if callback.Data == "-1" {

		msg := &tgbotapi.EditMessageTextConfig{Text: ""}
		keyboardCategories := h.services.GetCategoriesInlineKeyboard()
		msg.ReplyMarkup = &keyboardCategories
		return msg, nil
	}

	if _, err := strconv.Atoi(callback.Data); err == nil {
		msg := &tgbotapi.EditMessageTextConfig{Text: ""}
		keyboard := h.services.LinkService.GetResourceNamesInlineKeyboard(callback.Data)
		msg.ReplyMarkup = &keyboard
		return msg, nil
	}

	var messageTemplate string = ""

	data := h.services.LinkService.GetLinks(callback.Data)

	for i := 0; i < len(data); i++ {
		messageTemplate += data[i]
	}

	msg := &tgbotapi.EditMessageTextConfig{Text: messageTemplate}

	keyboardBack := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Назад", "-1"),
		),
	)

	msg.ReplyMarkup = &keyboardBack

	return msg, nil

}
