package GRPCAuthHandler

import (
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/models"
	"gitgub.com/diploma-mppr/backend_mppr/internal/microservices/auth"
	"gitgub.com/diploma-mppr/backend_mppr/internal/microservices/auth/AuthPB"
	"golang.org/x/net/context"
)

//const (
//	tokenCookieKey = "token"
//)

type HandlerAuth struct {
	UseCase auth.UseCase
}

func NewHandlerAuth(usecase auth.UseCase) *HandlerAuth {
	return &HandlerAuth{
		UseCase: usecase,
	}
}

//
//func createTokenCookie(token string, domen string, exp time.Duration) *http.Cookie {
//	return &http.Cookie{
//		Name:     tokenCookieKey,
//		Value:    token,
//		HttpOnly: true,
//		Expires:  time.Now().Add(exp),
//		Domain:   domen,
//		Path:     "/",
//		//SameSite: http.SameSiteNoneMode,
//		Secure: true,
//	}
//}

type OK struct {
	Message string `json:"message"`
}

func (h HandlerAuth) Register(ctx context.Context, UserRequest *authpb.UserS) (*authpb.ResponseUser, error) {
	User, err := h.UseCase.Register(&models.UserJson{Username: UserRequest.GetUsername(), Password: UserRequest.GetPassword()})
	if err != nil {
		return nil, err
	}
	return &authpb.ResponseUser{Id: uint64(User.Id), Username: User.Username}, nil
}

func (h HandlerAuth) Login(ctx context.Context, UserRequest *authpb.UserS) (*authpb.ResponseUser, error) {
	User, err := h.UseCase.Login(&models.UserJson{Username: UserRequest.GetUsername(), Password: UserRequest.GetPassword()})
	if err != nil {
		return nil, err
	}
	return &authpb.ResponseUser{Id: uint64(User.Id), Username: User.Username}, nil
}

func (h HandlerAuth) GetUserById(ctx context.Context, UserId *authpb.UserId) (*authpb.ResponseUser, error) {
	User, err := h.UseCase.GetUserById(models.UserId(UserId.GetUserId()))
	if err != nil {
		return nil, err
	}
	return &authpb.ResponseUser{Id: uint64(User.Id), Username: User.Username}, nil
}
