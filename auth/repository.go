package auth

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

type authRepoAdapter interface {
	//sign-in
	searchAgent(ctx context.Context, chatId int64) (string, error)

	//sign-up
	casheAgent(ctx context.Context, opt *KeyTokenOpt) error

	//check temp row in redis for validation work with model or token
	searchTempValue(ctx context.Context, chatId int64) (string, error)

	//cashe temp row in redis for given chatId and value of temp
	casheTempValue(ctx context.Context, chatId int64, temp string) error
}

type authRepo struct {

	//object to connection to the redis-base
	Db *redis.Client
}

func newAuthRepo(db *redis.Client) *authRepo {
	return &authRepo{
		Db: db,
	}
}

//search agent in database by key
//cheking the auth status of user and selecting the todoist-token
func (ar *authRepo) searchAgent(ctx context.Context, chatId int64) (string, error) {
	if chatId == 0 {
		return "", errNilChatId
	}

	iter := ar.Db.Scan(ctx, 0, "auth*", 0).Iterator()
	if err := iter.Err(); err != nil {
		return "", err
	}

	for iter.Next(ctx) {
		key := iter.Val()
		status := isAuth(key, chatId).status

		if status {
			//auth succes log
			log.Printf("auth is succes: chat_id - [%d]", chatId)

			token, err := ar.Db.Get(ctx, key).Result()
			if err != nil {
				return "", err
			}

			return token, nil
		}
	}

	return "", errAuth
}


//cashe agent in database according to the auth key-value scheme
//validation opt inner data on nillable
func (ar *authRepo) casheAgent(ctx context.Context, opt *KeyTokenOpt) error {
	var (
		token  = opt.Token
		chatId = opt.ChatId
	)

	if chatId == 0 || token == "" {
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
	key, err := makeAuthKey(chatId, int64(proj.Id))
	if err != nil {
		return nil
	}

	//log-key
	log.Printf("key - [%s]", key)

	//make pair in redis
	if err = ar.Db.Set(ctx, key, token, 0).Err(); err != nil {
		return err
	}

	//test-logic
	val, err := ar.Db.Get(ctx, key).Result()
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



//determine temp value for given chatId	
func (ar *authRepo) searchTempValue(ctx context.Context, chatId int64) (string, error) {
	key, err := makeTempKey(chatId)
	//		test-log
	log.Printf("temp key is [%s]", key)
	if err != nil {
		return "", err
	}


	val, err := ar.Db.Get(ctx, key).Result()
	//		test-log
	log.Printf("val temp is [%s]", val)
	if err != nil {
		return "", err
	}

	return val, nil
}


func (ar *authRepo) casheTempValue(ctx context.Context, chatId int64, temp string) error {
	if temp == "" {
		return errNilTemp
	}

	key, err := makeTempKey(chatId)
	//		test-log
	log.Printf("temp key is [%s]", key)
	if err != nil {
		return err
	}

	val, err := ar.Db.Set(ctx, key, temp, 0).Result()
	//		test-log
	log.Printf("val temp creation is [%s]", val)
	if err != nil {
		return err
	}

	return nil
}





