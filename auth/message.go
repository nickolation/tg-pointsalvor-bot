package auth

type authMessageInterface interface {
	t() error
}

type authMessage struct {
}

func newAuthMessage() *authMessage {
	return &authMessage{}
}

func (msg *authMessage) t() error {
	return nil
}
