package auth

type authRepoInterface interface {
	checkInRedis()
	otherCheckInRedis()
	createNewAgentPair()
}

type authRepo struct {
}

func newAuthRepo() *authRepo {
	return &authRepo{}
}

func (ar *authRepo) checkInRedis() {}

func (ar *authRepo) otherCheckInRedis() {}

func (ar *authRepo) createNewAgentPair() {}
