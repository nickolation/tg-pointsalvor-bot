package auth

import (
	"context"
	"log"
)

type Auth interface {
	SingIn(ctx context.Context, chatId int64) (string, error)
	SignUp(ctx context.Context, opt *KeyTokenOpt) error
}

type AuthEngine struct {
	Engine AuthServiceInterface
}

func newAuthEngine(engine AuthServiceInterface) *AuthEngine {
	return &AuthEngine{
		Engine: engine,
	}
}

func (eng *AuthEngine) SignIn(ctx context.Context, chatId int64) (string, error) {
	token, err := eng.Engine.VerifyAgent(ctx, chatId)
	if err != nil {
		//		test-log
		log.Printf("error with verify - [%s]", err.Error())
		return "", err
	}

	return token, nil
}

func (eng *AuthEngine) SignUp(ctx context.Context, opt *KeyTokenOpt) error {

	if err := eng.Engine.CreateTable(ctx, opt); err != nil {
		//		test-log
		log.Printf("error with signUp - [%s]", err.Error())
		return err
	}

	return nil
}
