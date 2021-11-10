package auth

import "github.com/go-redis/redis/v8"

type authRepoInterface interface {
	//sign-in
	searchAgent()

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

func (ar *authRepo) searchAgent() {}

func (ar *authRepo) foreignAuth() {}

func (ar *authRepo) casheAgent() {}
