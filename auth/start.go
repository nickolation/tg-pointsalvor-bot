package auth

import (
	"github.com/go-redis/redis/v8"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func InitAuth(db *redis.Client, api *tgbotapi.BotAPI) *AuthEngine {
	//init dependency // per-layers
	store := newAuthRepo(db)
	msg := newAuthMessage(api)
	service := newAuthService(store, msg)

	return newAuthEngine(service)
}
