package storage

import (
	"context"

	"github.com/go-redis/redis/v8"
)

//opts for init BotStorage
type StorageOptions struct {
	Port string
	Host string
	Ctx  context.Context
}

//implementation of repo-lay
type BotStorage struct {
	Client *redis.Client
}

func NewBotStorage(opt StorageOptions) (*BotStorage, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     opt.Host + opt.Port,
		Password: "",
		DB:       0,
	})

	if err := rdb.Ping(opt.Ctx).Err(); err != nil {
		return nil, err
	}

	return &BotStorage{
		Client: rdb,
	}, nil
}
