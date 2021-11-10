package auth

type Auth interface {
	SingIn()
	ForeignSignIn()
	SignUp()
}

type AuthEngine struct {
	engine *authServiceInterface
}

func newAuthEngine(engine authServiceInterface) *AuthEngine {
	return &AuthEngine{
		engine: &engine,
	}
}

func (eng *AuthEngine) SignIn() {}

func (eng *AuthEngine) ForeignSignIn() {}

func (eng *AuthEngine) SignUp() {}
