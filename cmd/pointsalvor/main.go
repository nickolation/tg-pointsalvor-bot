package main

import (
	"context"
	"log"

	botapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/nickolation/tg-pointsalvor-bot/auth"
	"github.com/nickolation/tg-pointsalvor-bot/engine"
	"github.com/nickolation/tg-pointsalvor-bot/engine/external/storage"
	"github.com/nickolation/tg-pointsalvor-bot/engine/handler"
	"github.com/nickolation/tg-pointsalvor-bot/ui"
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

	//open token 
	api, err := botapi.NewBotAPI("2020337404:AAHz7iU8yGaWwwPYhIjcu8HHY4QUCpnZKvo")
	if err != nil {
		log.Printf("error - [%s]", err.Error())
	}

	auth := auth.InitAuth(store.Client, api)

	ui := ui.NewUi(api)
	handler := handler.NewHandler(ui, nil, auth)

	if bot := engine.NewEngineBot(api, handler, nil); bot != nil {
		bot.StartEngine(ctx)
	}
}
