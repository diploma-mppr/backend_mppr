package PairComparisonCriteriaHandler

import (
	"encoding/json"
	"fmt"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/middleware"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/models"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/pairComparisonCriteria"
	"gitgub.com/diploma-mppr/backend_mppr/tools"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
	"reflect"
	"strconv"
)

type HandlerPairComparisonCriteria struct {
	UseCase pairComparisonCriteria.UseCase
}

func NewHandlerPairComparisonCriteria(useCase pairComparisonCriteria.UseCase) *HandlerPairComparisonCriteria {
	return &HandlerPairComparisonCriteria{
		UseCase: useCase,
	}
}

func (h HandlerPairComparisonCriteria) GetPairComparisonCriteria(ctx echo.Context) error {
	user := middleware.GetUserFromCtx(ctx)
	if user == nil {
		return tools.CustomError(ctx, errors.Errorf("пользователь не в системе"), 0, "ошибка при запросе пользователя")
	}

	id := ctx.QueryParam("id")
	fmt.Println("Param: ", id, " ", reflect.TypeOf(id))
	che, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("HandlerPairComparisonCriteria GetPairComparisonCriteria", err)
		return tools.CustomError(ctx, err, 0, "ParseInt")
	}

	data, err := h.UseCase.GetPairComparisonCriteria(int(che), int(user.Id))
	if err != nil {
		fmt.Println("HandlerPairComparisonCriteria GetPairComparisonCriteria", err)
		return tools.CustomError(ctx, err, 1, "UseCase")
	}

	result, err := json.Marshal(data)
	if err != nil {
		fmt.Println("HandlerPairComparisonCriteria GetPairComparisonCriteria", err)
		return tools.CustomError(ctx, err, 2, "Marshal")
	}

	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerPairComparisonCriteria) SetPairComparisonCriteria(ctx echo.Context) error {
	user := middleware.GetUserFromCtx(ctx)
	if user == nil {
		return tools.CustomError(ctx, errors.Errorf("пользователь не в системе"), 0, "ошибка при запросе пользователя")
	}

	data := models.PairComparisonCriteriaJson{}
	if err := ctx.Bind(&data); err != nil {
		fmt.Println("HandlerPairComparisonCriteria SetPairComparisonCriteria", err)
		tools.CustomError(ctx, err, 0, "Bind")
	}

	fmt.Println(data)

	task, err := h.UseCase.SetPairComparisonCriteria(&data, int(user.Id))
	if err != nil {
		fmt.Println("HandlerPairComparisonCriteria SetPairComparisonCriteria", err)
		tools.CustomError(ctx, err, 1, "Usecase")
	}

	result, err := json.Marshal(task)
	if err != nil {
		fmt.Println("HandlerPairComparisonCriteria SetPairComparisonCriteria", err)
		tools.CustomError(ctx, err, 2, "Marshal")
	}

	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerPairComparisonCriteria) UpdatePairComparisonCriteria(ctx echo.Context) error {
	data := models.PairComparisonCriteriaJson{}
	if err := ctx.Bind(&data); err != nil {
		fmt.Println("HandlerPairComparisonCriteria UpdatePairComparisonCriteria", err)
		tools.CustomError(ctx, err, 0, "Bind")
	}

	fmt.Println(data)

	task, err := h.UseCase.UpdatePairComparisonCriteria(&data)
	if err != nil {
		fmt.Println("HandlerPairComparisonCriteria UpdatePairComparisonCriteria", err)
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

func (h HandlerPairComparisonCriteria) DeletePairComparisonCriteria(ctx echo.Context) error {
	data := models.PairComparisonCriteriaJson{}
	if err := ctx.Bind(&data); err != nil {
		fmt.Println("HandlerPairComparisonCriteria DeletePairComparisonCriteria", err)
		tools.CustomError(ctx, err, 0, "Bind")
	}

	fmt.Println(data)

	err := h.UseCase.DeletePairComparisonCriteria(&data)
	if err != nil {
		fmt.Println("HandlerPairComparisonCriteria DeletePairComparisonCriteria", err)
		tools.CustomError(ctx, err, 1, "DeletePairComparisonCriteria")
	}

	return ctx.JSONBlob(http.StatusOK, nil)
}
