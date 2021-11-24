package auth

import (
	"context"
	"log"
)

type Auth interface {
	SignIn(ctx context.Context, chatId int64) (*KeyTokenOpt, error)
	SignUp(ctx context.Context, opt *KeyTokenOpt) error
	MakeTemp(ctx context.Context, chatId int64, temp string) error
	SearchTemp(ctx context.Context, chatId int64) (string, error)
	DeleteTemp(ctx context.Context, chatId int64) error
	AlreadyAuth(chatId int64) error
}

type AuthEngine struct {
	Engine AuthServiceInterface
}

func newAuthEngine(engine AuthServiceInterface) *AuthEngine {
	return &AuthEngine{
		Engine: engine,
	}
}

func (eng *AuthEngine) SignIn(ctx context.Context, chatId int64) (*KeyTokenOpt, error) {
	opt, err := eng.Engine.VerifyAgent(ctx, chatId)
	if err != nil {
		//		test-log
		log.Printf("error with verify - [%s]", err.Error())
		return nil, err
	}

	return opt, nil
}

func (eng *AuthEngine) SignUp(ctx context.Context, opt *KeyTokenOpt) error {
	if err := eng.Engine.CreateTable(ctx, opt); err != nil {
		//		test-log
		log.Printf("error with signUp - [%s]", err.Error())
		return err
	}

	return nil
}


func (eng *AuthEngine) MakeTemp(ctx context.Context, chatId int64, temp string) error {
	if err := eng.Engine.MakeTemp(ctx, chatId, temp); err != nil {
		//		test-log
		log.Printf("error with signUp - [%s]", err.Error())
		return err
	}

	return nil
}


func (eng *AuthEngine) SearchTemp(ctx context.Context, chatId int64) (string, error) {
	val, err := eng.Engine.SearchTemp(ctx, chatId)
	if err != nil {
		//		test-log
		log.Printf("error with message - [%s]", err.Error())

		return "", err
	}

	return val, nil
}

func (eng *AuthEngine) AlreadyAuth(chatId int64) error {
	if err := eng.Engine.AlreadyAuthorized(chatId); err != nil {
		//		test-log
		log.Printf("error with message - [%s]", err.Error())

		return err
	}

	log.Printf("already auth - [%d]", chatId)
	return nil
}


func (eng *AuthEngine) DeleteTemp(ctx context.Context, chatId int64) error {
	if err := eng.Engine.DeleteTemp(ctx, chatId); err != nil {
		return err 
	}
	
	return nil
}