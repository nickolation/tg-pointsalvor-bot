package engine

import (
	"context"
	"log"

	botapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (bot *EngineBot) StartEngine(ctx context.Context) error {
	//new update config and cfg setting
	cfg := botapi.NewUpdate(0)
	cfg.Timeout = 60

	//get update in channel with own goroutine
	//block operation
	updates, err := bot.bot.GetUpdatesChan(cfg)
	if err != nil {
		return err
	}

	//iteratung uodate and handling signals from chat
	for update := range updates {
		msg := update.Message

		if msg == nil {
			continue
		}

		if msg.IsCommand() {
			bot.handler.HandleStart(ctx, msg.Chat.ID)
			continue
		}

		if err := bot.handler.HandleMessage(ctx, msg.Chat.ID, msg.Text); err != nil {
			//		test-log
			log.Printf("error with handling - [%s]", err.Error())
			return err
		}
	}

	return nil
}
