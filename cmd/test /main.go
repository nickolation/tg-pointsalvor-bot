package main

import (
	"context"
	"log"

	botapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/nickolation/tg-pointsalvor-bot/auth"
	"github.com/nickolation/tg-pointsalvor-bot/engine"
	"github.com/nickolation/tg-pointsalvor-bot/engine/external/storage"
	"github.com/nickolation/tg-pointsalvor-bot/engine/handler"
)

func main() {
	ctx := context.Background()

	store, err := storage.NewBotStorage(storage.StorageOptions{
		Port: "7000",
		Host: "localhost:",
		Ctx:  ctx,
	})
	if err != nil {
		log.Printf("error - [%s]", err.Error())
	}

	//		token is dirty 
	api, err := botapi.NewBotAPI("<token>")
	if err != nil {
		log.Printf("error - [%s]", err.Error())
	}

	auth := auth.InitAuth(store.Client, api)

	handler := handler.NewHandler(nil, nil, auth)

	if bot := engine.NewEngineBot(api, handler, nil); bot != nil {
		bot.StartEngine(ctx)
	}
}








