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
		//test-log
		log.Printf("error with getting updates - [%s]", err.Error())

		return err
	}

	//iterating update and handling the signals from the chat
	for update := range updates {
		msg := update.Message

		if msg == nil {
			continue
		}

		//data of message
		chatId := msg.Chat.ID
		data := msg.Text

		if cmd := msg.Command(); cmd != "" {
			if cmd == "start" {
				if err := bot.handler.HandleStart(ctx, chatId); err != nil {
					//		test-log
					log.Printf("error with start message - [%s]", err.Error())
				}
				continue
			}

			//		test-logs
			log.Printf("what is the command? - [%s]", cmd)
			if err := bot.handler.HandleForeignCommand(chatId); err != nil {
				//		...
				log.Printf("error with handling foreign command - [%s]", err)
			}

			continue
		}

		if err := bot.handler.HandleMessage(ctx, chatId, data); err != nil {
			//		test-log
			log.Printf("error with handling - [%s]", err.Error())
		}
	}

	return nil
}
