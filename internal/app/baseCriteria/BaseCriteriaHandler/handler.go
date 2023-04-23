package BaseCriteriaHandler

import (
	"encoding/json"
	"fmt"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/baseCriteria"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/middleware"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/models"
	"gitgub.com/diploma-mppr/backend_mppr/tools"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
	"reflect"
	"strconv"
)

type HandlerBaseCriteria struct {
	UseCase baseCriteria.UseCase
}

func NewHandlerBaseCriteria(useCase baseCriteria.UseCase) *HandlerBaseCriteria {
	return &HandlerBaseCriteria{
		UseCase: useCase,
	}
}

func (h HandlerBaseCriteria) GetBaseCriteria(ctx echo.Context) error {
	user := middleware.GetUserFromCtx(ctx)
	if user == nil {
		return tools.CustomError(ctx, errors.Errorf("пользователь не в системе"), 0, "ошибка при запросе пользователя")
	}

	id := ctx.QueryParam("id")
	fmt.Println("Param: ", id, " ", reflect.TypeOf(id))
	che, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("HandlerPareto GetPareto", err)
		return tools.CustomError(ctx, err, 0, "ParseInt")
	}

	data, err := h.UseCase.GetBaseCriteria(int(che), int(user.Id))
	if err != nil {
		fmt.Println("HandlerPareto GetPareto", err)
		return tools.CustomError(ctx, err, 1, "UseCase")
	}

	result, err := json.Marshal(data)
	if err != nil {
		fmt.Println("HandlerPareto GetPareto", err)
		return tools.CustomError(ctx, err, 2, "Marshal")
	}

	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerBaseCriteria) SetBaseCriteria(ctx echo.Context) error {
	user := middleware.GetUserFromCtx(ctx)
	if user == nil {
		return tools.CustomError(ctx, errors.Errorf("пользователь не в системе"), 0, "ошибка при запросе пользователя")
	}

	data := models.BaseCriteriaJson{}
	if err := ctx.Bind(&data); err != nil {
		fmt.Println("HandlerPareto SetPareto", err)
		tools.CustomError(ctx, err, 0, "Bind")
	}

	fmt.Println(data)

	task, err := h.UseCase.SetBaseCriteria(&data, int(user.Id))
	if err != nil {
		fmt.Println("HandlerPareto SetPareto", err)
		tools.CustomError(ctx, err, 1, "Usecase")
	}

	result, err := json.Marshal(task)
	if err != nil {
		fmt.Println("HandlerPareto SetPareto", err)
		tools.CustomError(ctx, err, 2, "Marshal")
	}

	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerBaseCriteria) UpdateBaseCriteria(ctx echo.Context) error {
	data := models.BaseCriteriaJson{}
	if err := ctx.Bind(&data); err != nil {
		fmt.Println("HandlerPareto UpdatePareto", err)
		tools.CustomError(ctx, err, 0, "Bind")
	}

	fmt.Println(data)

	task, err := h.UseCase.UpdateBaseCriteria(&data)
	if err != nil {
		fmt.Println("HandlerPareto UpdatePareto", err)
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

func (h HandlerBaseCriteria) DeleteBaseCriteria(ctx echo.Context) error {
	data := models.BaseCriteriaJson{}
	if err := ctx.Bind(&data); err != nil {
		fmt.Println("HandlerPareto DeletePareto", err)
		tools.CustomError(ctx, err, 0, "Bind")
	}

	fmt.Println(data)

	err := h.UseCase.DeleteBaseCriteria(&data)
	if err != nil {
		fmt.Println("HandlerPareto DeletePareto", err)
		tools.CustomError(ctx, err, 1, "DeletePareto")
	}

	return ctx.JSONBlob(http.StatusOK, nil)
}
