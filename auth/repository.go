package auth

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	sdk "github.com/nickolation/pointsalvor"
)

type authRepoInterface interface {
	//sign-in
	searchAgent(ctx context.Context, db *redis.Client, chatId string) (string, error)

	//foreign sign-in
	foreignAuth()

	//sign-up
	casheAgent(ctx context.Context, db *redis.Client, chatId string, ag *sdk.Agent) error
}

type authRepo struct {
	db *redis.Client

	//pointsalvor agent for toodist-requests
	ag *sdk.Agent
}

func newAuthRepo(db *redis.Client) *authRepo {
	return &authRepo{
		db: db,
	}
}

//search agent in database by key
//cheking the auth status of user and selecting the todoist-token
func (ar *authRepo) searchAgent(ctx context.Context, db *redis.Client, chatId string) (string, error) {
	if chatId == "" {
		return "", errNilChatId
	}

	iter := db.Scan(ctx, 0, "auth*", 0).Iterator()
	if err := iter.Err(); err != nil {
		return "", err
	}

	for iter.Next(ctx) {
		key := iter.Val()
		status := isAuth(key, chatId).status

		if status {
			//auth succes log
			log.Printf("auth is succes: chat_id - [%s]", chatId)

			token, err := db.Get(ctx, key).Result()
			if err != nil {
				return "", err
			}

			return token, nil
		}
	}

	return "", errAuth
}

func (ar *authRepo) foreignAuth() {}

//cashe agent in database according to the auth key-value scheme
//validation opt inner data on nillable
func (ar *authRepo) casheAgent(ctx context.Context, db *redis.Client, opt *KeyTokenOpt) error {
	var (
		token  = opt.token
		chatId = opt.chatId
	)

	if chatId == "" || token == "" {
		return errNilOpt
	}

	//link agent with same token
	ag, err := linkAgent(token)
	if err != nil {
		return err
	}

	//make table
	var tableName = callTable(chatId)

	proj, err := ag.AddProject(ctx, tableName)
	if err != nil {
		return err
	}

	//make key for auth: key- value
	key, err := makeKey(chatId, proj.Id)
	if err != nil {
		return nil
	}
	//log-key
	log.Printf("key - [%s]", key)

	//make pair in redis
	if err = db.Set(ctx, key, token, 0).Err(); err != nil {
		return err
	}

	//test-logic
	val, err := db.Get(ctx, key).Result()
	if err != nil {
		//log key-value
		log.Printf("token value for key [%s] - [%s]", key, val)
		return err
	}
	//same log
	log.Printf("token value for key [%s] - [%s]", key, val)
	//test-logic

	return nil
}
