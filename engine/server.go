package sengine

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	sdk "github.com/nickolation/pointsalvor"
	hnd "github.com/nickolation/tg-pointsalvor-bot/internal/API/handler"
	store "github.com/nickolation/tg-pointsalvor-bot/internal/API/repository/storage"
	svice "github.com/nickolation/tg-pointsalvor-bot/internal/API/service"
)

type BotEngine struct {
	bot   *tgbotapi.BotAPI
	agent *sdk.Agent
	token string

	handler *hnd.Handler
	service *svice.Service
	storage *store.Storage
}

func NewBotEngine(b *tgbotapi.BotAPI, a *sdk.Agent, t string,
	h *hnd.Handler, s *svice.Service, st *store.Storage) *BotEngine {
	return &BotEngine{
		bot:     b,
		agent:   a,
		token:   t,
		handler: h,
		service: s,
		storage: st,
	}
}
