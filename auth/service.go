package auth

type authServiceInterface interface {
	createTable() error
	checkAgent() error
	otherHandChecking() error
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

func (ar *authService) createTable() error {
	return nil
}

func (ar *authService) checkAgent() error {
	return nil
}

func (ar *authService) otherHandChecking() error {
	return nil
}
