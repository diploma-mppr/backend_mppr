package middleware

import "gitgub.com/diploma-mppr/backend_mppr/tools/authManager"

type CommonMiddleware struct {
	AuthManager authManager.AuthManager
}

func NewCommonMiddleware(authManager authManager.AuthManager) CommonMiddleware {
	return CommonMiddleware{
		AuthManager: authManager,
	}
}
