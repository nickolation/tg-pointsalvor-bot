package auth

func InitAuth() *AuthEngine {
	//init dependency // per-layers
	return newAuthEngine(newAuthService(newAuthRepo(), newAuthMessage()))
}
