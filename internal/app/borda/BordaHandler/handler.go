package BordaHandler

import (
	"encoding/json"
	"fmt"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/borda"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/middleware"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/models"
	"gitgub.com/diploma-mppr/backend_mppr/tools"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
	"reflect"
	"strconv"
)

type HandlerBorda struct {
	UseCase borda.UseCase
}

func NewHandlerBorda(useCase borda.UseCase) *HandlerBorda {
	return &HandlerBorda{
		UseCase: useCase,
	}
}

func (h HandlerBorda) GetBorda(ctx echo.Context) error {
	user := middleware.GetUserFromCtx(ctx)
	if user == nil {
		return tools.CustomError(ctx, errors.Errorf("пользователь не в системе"), 0, "ошибка при запросе пользователя")
	}

	id := ctx.QueryParam("id")
	fmt.Println("Param: ", id, " ", reflect.TypeOf(id))
	che, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("HandlerBorda GetBorda", err)
		return tools.CustomError(ctx, err, 0, "ParseInt")
	}

	data, err := h.UseCase.GetBorda(int(che), int(user.Id))
	if err != nil {
		fmt.Println("HandlerBorda GetBorda", err)
		return tools.CustomError(ctx, err, 1, "UseCase")
	}

	result, err := json.Marshal(data)
	if err != nil {
		fmt.Println("HandlerBorda GetBorda", err)
		return tools.CustomError(ctx, err, 2, "Marshal")
	}

	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerBorda) SetBorda(ctx echo.Context) error {
	user := middleware.GetUserFromCtx(ctx)
	if user == nil {
		return tools.CustomError(ctx, errors.Errorf("пользователь не в системе"), 0, "ошибка при запросе пользователя")
	}

	data := models.BordaJson{}
	if err := ctx.Bind(&data); err != nil {
		fmt.Println("HandlerBorda SetBorda", err)
		tools.CustomError(ctx, err, 0, "Bind")
	}

	fmt.Println(data)

	task, err := h.UseCase.SetBorda(&data, int(user.Id))
	if err != nil {
		fmt.Println("HandlerBorda SetBorda", err)
		tools.CustomError(ctx, err, 1, "Usecase")
	}

	result, err := json.Marshal(task)
	if err != nil {
		fmt.Println("HandlerBorda SetBorda", err)
		tools.CustomError(ctx, err, 2, "Marshal")
	}

	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerBorda) UpdateBorda(ctx echo.Context) error {
	data := models.BordaJson{}
	if err := ctx.Bind(&data); err != nil {
		fmt.Println("HandlerBorda UpdateBorda", err)
		tools.CustomError(ctx, err, 0, "Bind")
	}

	fmt.Println(data)

	task, err := h.UseCase.UpdateBorda(&data)
	if err != nil {
		fmt.Println("HandlerBorda UpdateBorda", err)
		tools.CustomError(ctx, err, 1, "UseCase")
	}

	result, err := json.Marshal(task)
	if err != nil {
		fmt.Println(err.Error())
		tools.CustomError(ctx, err, 2, "Marshal")
	}
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerBorda) DeleteBorda(ctx echo.Context) error {
	data := models.BordaJson{}
	if err := ctx.Bind(&data); err != nil {
		fmt.Println("HandlerBorda DeleteBorda", err)
		tools.CustomError(ctx, err, 0, "Bind")
	}

	fmt.Println(data)

	err := h.UseCase.DeleteBorda(&data)
	if err != nil {
		fmt.Println("HandlerBorda DeleteBorda", err)
		tools.CustomError(ctx, err, 1, "DeleteBorda")
	}

	return ctx.JSONBlob(http.StatusOK, nil)
}
