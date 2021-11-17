package auth

import (
	"context"
	"log"
)

type AuthServiceInterface interface {
	//sign-up
	CreateTable(ctx context.Context, opt *KeyTokenOpt) error

	//sign-in
	VerifyAgent(ctx context.Context, chatId int64) (string, error)
}

type authService struct {
	authRepoInterface
	authMessageInterface
}

func newAuthService(store authRepoInterface, msg authMessageInterface) *authService {
	return &authService{
		authRepoInterface:    store,
		authMessageInterface: msg,
	}
}

func (ar *authService) CreateTable(ctx context.Context, opt *KeyTokenOpt) error {
	if err := ar.authRepoInterface.casheAgent(ctx, opt); err != nil {
		return err
	}
	
	if err := ar.authMessageInterface.SuccesAuthorized(opt.ChatId); err != nil {
		//		test-log
		log.Printf("error with message - [%s]", err.Error())
		return err 

	}
	

	return nil
}

func (ar *authService) VerifyAgent(ctx context.Context, chatId int64) (string, error) {
	token, err := ar.authRepoInterface.searchAgent(ctx, chatId)
	if err != nil {
		return "", err
	}

	return token, nil
}
