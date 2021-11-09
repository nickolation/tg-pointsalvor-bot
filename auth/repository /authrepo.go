package authrepository

type AuthRepoAdapter interface {
	CheckInRedis()
	OtherCheckInRedis()
	CreateNewAgentPair()
}

type AuthRepo struct {
}

func NewAuthRepo() *AuthRepo {
	return &AuthRepo{}
}

func (ar *AuthRepo) CheckInRedis() {}

func (ar *AuthRepo) OtherCheckInRedis() {}

func (ar *AuthRepo) CreateNewAgentPair() {}
