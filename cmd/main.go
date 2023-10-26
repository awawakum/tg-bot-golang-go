package main

import (
	"log"

	"github.com/go-bot/internal/clients/telegram"
	"github.com/go-bot/internal/config"
	"github.com/go-bot/internal/database"
	"github.com/go-bot/internal/handler"
	"github.com/go-bot/internal/service"
	"github.com/sirupsen/logrus"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		logrus.Fatal("False init config: ", err)
	}

	db, err := database.ConnectToDB(&cfg.DBConfig)
	if err != nil {
		logrus.Fatal("False db connection: ", err)
	}

	database := database.NewDataBase(db)
	services := service.NewService(database, cfg.ApiUrl)
	handlers := handler.NewHandler(services)

	bot, err := telegram.NewBot(cfg.TelegramToken, handlers)
	if err != nil {
		log.Fatalf("Failed to initialize Telegram bot: %v", err)
	}

	if err := bot.Run(); err != nil {
		log.Fatalf("Failed to run the bot: %v", err)
	}
}
