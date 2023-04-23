package DataHandler

import (
	"encoding/json"
	"fmt"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/data"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/middleware"
	"gitgub.com/diploma-mppr/backend_mppr/tools"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

type HandlerData struct {
	UseCase data.UseCase
}

func NewHandlerData(useCase data.UseCase) *HandlerData {
	return &HandlerData{
		UseCase: useCase,
	}
}

func (h HandlerData) GetAll(ctx echo.Context) error {
	user := middleware.GetUserFromCtx(ctx)
	if user == nil {
		return tools.CustomError(ctx, errors.Errorf("пользователь не в системе"), 0, "ошибка при запросе пользователя")
	}

	data, err := h.UseCase.GetAll(int(user.Id))
	if err != nil {
		fmt.Println("HandlerData GetAll", err)
		return tools.CustomError(ctx, err, 0, "UseCase")
	}

	result, err := json.Marshal(data)
	if err != nil {
		return tools.CustomError(ctx, err, 1, "Marshal")
	}

	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}
