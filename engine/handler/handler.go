package handler

import (
	"context"

	"github.com/nickolation/tg-pointsalvor-bot/auth"
	"github.com/nickolation/tg-pointsalvor-bot/engine/service"
	"github.com/nickolation/tg-pointsalvor-bot/ui"
)

type HandlerAdapter interface {
	HandleCallback()
	HandleMessage(ctx context.Context, chatId int64, data string) error
	HandleStart(ctx context.Context, chatId int64) error
	HandleForeignCommand(chatId int64) error
}

type Handler struct {
	svice *service.Service
	ui    ui.UiAdapter
	auth  auth.Auth
}

func NewHandler(u ui.UiAdapter, s *service.Service, a auth.Auth) *Handler {
	return &Handler{
		svice: s,
		ui:    u,
		auth:  a,
	}
}
