package auth

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

type authRepoInterface interface {
	//sign-in
	searchAgent(ctx context.Context, db *redis.Client, chatId string) error

	//foreign sign-in
	foreignAuth()

	//sign-up
	casheAgent()
}

type authRepo struct {
	db *redis.Client
}

func newAuthRepo(db *redis.Client) *authRepo {
	return &authRepo{
		db: db,
	}
}

func (ar *authRepo) searchAgent(ctx context.Context, db *redis.Client, chatId string) error {
	if chatId == "" {
		return errNilChatId
	}

	iter := db.Scan(ctx, 0, "auth*", 0).Iterator()
	if err := iter.Err(); err != nil {
		return err
	}

	for iter.Next(ctx) {
		key := iter.Val()
		status := IsAuth(key, chatId)

		if status.Status {
			log.Printf("auth is succes: chat_id - [%s]", chatId)
			return nil
		}
	}

	return errAuth
}

func (ar *authRepo) foreignAuth() {}

func (ar *authRepo) casheAgent() {}
