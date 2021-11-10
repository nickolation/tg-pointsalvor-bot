package handler

import (
	"github.com/nickolation/tg-pointsalvor-bot/auth"
	service "github.com/nickolation/tg-pointsalvor-bot/engine/service"
	"github.com/nickolation/tg-pointsalvor-bot/ui"
)

type Handler struct {
	svice *service.Service
	ui    *ui.Ui
	auth  *auth.Auth
}

func NewHandler(u *ui.Ui, s *service.Service, a *auth.Auth) *Handler {
	return &Handler{
		svice: s,
		ui:    u,
		auth:  a,
	}
}
