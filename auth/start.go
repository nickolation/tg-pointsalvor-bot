package auth

import (
	"github.com/go-redis/redis/v8"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func InitAuth(db *redis.Client, api *tgbotapi.BotAPI) *AuthEngine {
	//init dependency // per-layers
	return newAuthEngine(newAuthService(newAuthRepo(db), newAuthMessage(api)))
}
