package auth

import (
	"context"
	"log"
)

type AuthServiceInterface interface {
	//sign-up
	CreateTable(ctx context.Context, opt *KeyTokenOpt) error

	//sign-in
	VerifyAgent(ctx context.Context, chatId int64) (*KeyTokenOpt, error)

	//don't cashe agent and enjoy the life 
	//send message about already auth
	AlreadyAuthorized(chatId int64) error

	//make temp row in redis
	MakeTemp(ctx context.Context, chatId int64, temp string) error

	//search temp value
	SearchTemp(ctx context.Context, chatId int64) (string, error)

	//delete temp value in service-level
	DeleteTemp(ctx context.Context, chatId int64) error
}

type authService struct {
	authRepoAdapter
	authMessageAdapter
}

func newAuthService(store authRepoAdapter, msg authMessageAdapter) *authService {
	return &authService{
		authRepoAdapter:    store,
		authMessageAdapter: msg,
	}
}

func (as *authService) CreateTable(ctx context.Context, opt *KeyTokenOpt) error {
	if err := as.authRepoAdapter.casheAgent(ctx, opt); err != nil {
		return err
	}

	if err := as.authMessageAdapter.SuccesAuthorized(opt.ChatId); err != nil {
		//		test-log	
		log.Printf("error with message - [%s]", err.Error())
		return err

	}

	return nil
}

func (as *authService) VerifyAgent(ctx context.Context, chatId int64) (*KeyTokenOpt, error) {
	opt, err := as.authRepoAdapter.searchAgent(ctx, chatId)
	if err != nil {
		return nil, err
	}

	return opt, nil
}

func (as *authService) MakeTemp(ctx context.Context, chatId int64, temp string) error {
	if err := as.authRepoAdapter.casheTempValue(ctx, chatId, temp); err != nil {
		//		test-log
		log.Printf("error with message - [%s]", err.Error())

		return err
	}

	return nil
}

func (as *authService) SearchTemp(ctx context.Context, chatId int64) (string, error) {
	val, err := as.authRepoAdapter.searchTempValue(ctx, chatId)
	if err != nil {
		//		test-log
		log.Printf("error with message - [%s]", err.Error())

		return "", err
	}

	return val, nil
}


func (as *authService) AlreadyAuthorized(chatId int64) error {
	if err := as.authMessageAdapter.AlreadyAuthorized(chatId); err != nil {
		//		test-log
		log.Printf("error with message - [%s]", err.Error())
		return err
	}

	return nil
}

func (as *authService) DeleteTemp(ctx context.Context, chatId int64) error {
	if err := as.authRepoAdapter.deleteTempValue(ctx, chatId); err != nil {
		return err
	}
	
	return nil
}