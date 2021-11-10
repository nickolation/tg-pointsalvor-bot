package storage

import (
	"context"

	"github.com/go-redis/redis/v8"
)

//opts for init BotStorage
type StorageOptions struct {
	port string
	host string
	ctx  context.Context
}

//implementation of repo-lay
type BotStorage struct {
	client *redis.Client
}

func NewBotStorage(opt StorageOptions) (*BotStorage, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     opt.host + opt.port,
		Password: "",
		DB:       0,
	})

	if err := rdb.Ping(opt.ctx).Err(); err != nil {
		return nil, err
	}

	return &BotStorage{
		client: rdb,
	}, nil
}
