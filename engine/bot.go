package engine

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	sdk "github.com/nickolation/pointsalvor"
	"github.com/nickolation/tg-pointsalvor-bot/engine/handler"
	"github.com/nickolation/tg-pointsalvor-bot/ui"
)

type EngineBot struct {
	bot     *tgbotapi.BotAPI
	agent   *sdk.Agent
	handler *handler.Handler	
}

func NewEngineBot(bot *tgbotapi.BotAPI, ag *sdk.Agent, hnd *handler.Handler, ui *ui.Ui) *EngineBot {
	return &EngineBot{
		bot:     bot,
		agent:   ag,
		handler: hnd,
	}
}
