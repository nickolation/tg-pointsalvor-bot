package engine

import (
	botapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (bot *EngineBot) StartEngine() error {
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
		if update.Message == nil {
			continue
		}
	}

	return nil
}
