package middleware

import (
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/models"
	"github.com/labstack/echo/v4"
)

const UserCtxKey = "user"
const TokenKeyCookie = "token"

type UserCtx struct {
	Id models.UserId
}

func (mw *CommonMiddleware) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		tokenCookie, err := ctx.Request().Cookie(TokenKeyCookie)

		if err != nil {
			return next(ctx)
		}

		payload, err := mw.AuthManager.ParseToken(tokenCookie.Value)
		if err != nil {
			return next(ctx)
		}

		//if _, err = govalidator.ValidateStruct(payload); err != nil {
		//	return next(ctx)
		//}

		ctx.Set(UserCtxKey, UserCtx{Id: models.UserId(payload.Id)})
		return next(ctx)
	}
}

func GetUserFromCtx(ctx echo.Context) *UserCtx {
	user, ok := ctx.Get(UserCtxKey).(UserCtx)
	if !ok {
		return nil
	}
	return &user
}
