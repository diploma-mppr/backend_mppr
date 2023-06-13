package AuthHandler

import (
	"context"
	"encoding/json"
	"fmt"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/middleware"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/models"
	authpb "gitgub.com/diploma-mppr/backend_mppr/internal/microservices/auth/AuthPB"
	"gitgub.com/diploma-mppr/backend_mppr/tools"
	"gitgub.com/diploma-mppr/backend_mppr/tools/authManager"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"strconv"
	"time"
)

const (
	tokenCookieKey = "token"
)

type HandlerAuth struct {
	AuthManager authManager.AuthManager
}

func NewHandlerAuth(authManager authManager.AuthManager) *HandlerAuth {
	return &HandlerAuth{
		AuthManager: authManager,
	}
}

func createTokenCookie(token string, domen string, exp time.Duration) *http.Cookie {
	return &http.Cookie{
		Name:     tokenCookieKey,
		Value:    token,
		HttpOnly: true,
		Expires:  time.Now().Add(exp),
		Domain:   domen,
		Path:     "/",
		//SameSite: http.SameSiteNoneMode,
		Secure: true,
	}
}

type OK struct {
	Message string `json:"message"`
}

var (
	che authpb.AuthMicroserviceClient
)

func (h HandlerAuth) Register(ctx echo.Context) error {
	if middleware.GetUserFromCtx(ctx) != nil {
		return tools.CustomError(ctx, errors.Errorf("пользователь находится в системе"), 0, "пользователь уже зарегестрирован")
	}

	var UserRequest models.UserJson
	err := ctx.Bind(&UserRequest)
	if err != nil || len(UserRequest.Username) == 0 || len(UserRequest.Password) == 0 {
		return tools.CustomError(ctx, err, 1, "битый json на авторизацию")
	}

	fmt.Println(UserRequest)

	grcpConn, err := grpc.Dial(
		"127.0.0.1:8001",
		grpc.WithInsecure(),
	)
	if err != nil {
		fmt.Println("cant connect to grpc")
	}
	defer grcpConn.Close()
	che = authpb.NewAuthMicroserviceClient(grcpConn)
	User, err := che.Register(context.Background(), &authpb.UserS{Username: UserRequest.Username, Password: UserRequest.Password})

	token, err := h.AuthManager.CreateToken(authManager.NewTokenPayload(models.UserId(User.GetId()))) // подставить id пользователя полученного из usecase
	if err != nil {
		return tools.CustomError(ctx, err, 3, "отрыгнул jsw token или что то связанное с ним")
	}

	host, _, _ := net.SplitHostPort(ctx.Request().Host)
	tokenCookie := createTokenCookie(token, host, h.AuthManager.GetEpiryTime())

	ctx.SetCookie(tokenCookie)

	result, _ := json.Marshal(models.ResponseUserJson{Id: models.UserId(User.GetId()), Username: User.GetUsername()})
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerAuth) Login(ctx echo.Context) error {
	if middleware.GetUserFromCtx(ctx) != nil {
		return tools.CustomError(ctx, errors.Errorf("пользователь уже вошел в систему"), 0, "пользователь уже зарегестрирован")
	}

	var UserRequest models.UserJson
	err := ctx.Bind(&UserRequest)
	if err != nil || len(UserRequest.Username) == 0 || len(UserRequest.Password) == 0 {
		return tools.CustomError(ctx, err, 1, "битый json на логин")
	}

	grcpConn, err := grpc.Dial(
		"127.0.0.1:8001",
		grpc.WithInsecure(),
	)
	if err != nil {
		fmt.Println("cant connect to grpc")
	}
	defer grcpConn.Close()
	che = authpb.NewAuthMicroserviceClient(grcpConn)
	User, err := che.Login(context.Background(), &authpb.UserS{Username: UserRequest.Username, Password: UserRequest.Password})

	token, err := h.AuthManager.CreateToken(authManager.NewTokenPayload(models.UserId(User.GetId()))) // подставить id пользователя полученного из usecase
	if err != nil {
		return tools.CustomError(ctx, err, 2, "отрыгнул jsw token или что то связанное с ним")
	}

	host, _, _ := net.SplitHostPort(ctx.Request().Host)
	tokenCookie := createTokenCookie(token, host, h.AuthManager.GetEpiryTime())

	ctx.SetCookie(tokenCookie)

	result, _ := json.Marshal(models.ResponseUserJson{Id: models.UserId(User.GetId()), Username: User.GetUsername()})
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerAuth) Logout(ctx echo.Context) error {
	if middleware.GetUserFromCtx(ctx) == nil {
		return tools.CustomError(ctx, errors.Errorf("пользователь не в системе"), 2, "ошибка при логине")
	}
	host, _, _ := net.SplitHostPort(ctx.Request().Host)
	resetTokenCookie := createTokenCookie("", host, -time.Hour)

	ctx.SetCookie(resetTokenCookie)

	var che = OK{
		Message: "успешно вышли из системы",
	}

	result, _ := json.Marshal(che)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerAuth) GetUserById(ctx echo.Context) error {
	user := middleware.GetUserFromCtx(ctx)
	if user == nil {
		return tools.CustomError(ctx, errors.Errorf("пользователь не в системе"), 0, "ошибка при запросе пользователя")
	}

	grcpConn, err := grpc.Dial(
		"127.0.0.1:8001",
		grpc.WithInsecure(),
	)
	if err != nil {
		fmt.Println("cant connect to grpc")
	}
	defer grcpConn.Close()
	che = authpb.NewAuthMicroserviceClient(grcpConn)
	User, err := che.GetUserById(context.Background(), &authpb.UserId{UserId: uint64(user.Id)})

	result, _ := json.Marshal(models.ResponseUserJson{Id: models.UserId(User.GetId()), Username: User.GetUsername()})
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}
