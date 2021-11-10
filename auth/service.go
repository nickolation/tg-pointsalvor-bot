package auth

type authServiceInterface interface {
	//sign-up
	createTable() error

	//sign-in
	verifyAgent() error

	//foreign sign-in
	foreignVerificate() error
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

func (ar *authService) verifyAgent() error {
	return nil
}

func (ar *authService) foreignVerificate() error {
	return nil
}
