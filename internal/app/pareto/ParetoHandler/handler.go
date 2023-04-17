package ParetoHandler

import (
	"encoding/json"
	"fmt"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/models"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/pareto"
	"gitgub.com/diploma-mppr/backend_mppr/tools"

	"github.com/labstack/echo/v4"
	"net/http"
	"reflect"
	"strconv"
)

type HandlerPareto struct {
	UseCase pareto.UseCase
}

func NewHandlerPareto(useCase pareto.UseCase) *HandlerPareto {
	return &HandlerPareto{
		UseCase: useCase,
	}
}

func (h HandlerPareto) GetPareto(ctx echo.Context) error {
	id := ctx.QueryParam("id")
	fmt.Println("Param: ", id, " ", reflect.TypeOf(id))
	che, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("HandlerPareto GetPareto", err)
		return tools.CustomError(ctx, err, 0, "ParseInt")
	}

	data, err := h.UseCase.GetPareto(int(che))
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

func (h HandlerPareto) SetPareto(ctx echo.Context) error {
	data := models.ParetoJson{}
	if err := ctx.Bind(&data); err != nil {
		fmt.Println("HandlerPareto SetPareto", err)
		tools.CustomError(ctx, err, 0, "Bind")
	}

	fmt.Println(data)

	task, err := h.UseCase.SetPareto(&data)
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

func (h HandlerPareto) UpdatePareto(ctx echo.Context) error {
	data := models.ParetoJson{}
	if err := ctx.Bind(&data); err != nil {
		fmt.Println("HandlerPareto UpdatePareto", err)
		tools.CustomError(ctx, err, 0, "Bind")
	}

	fmt.Println(data)

	task, err := h.UseCase.UpdatePareto(&data)
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

func (h HandlerPareto) DeletePareto(ctx echo.Context) error {
	data := models.ParetoJson{}
	if err := ctx.Bind(&data); err != nil {
		fmt.Println("HandlerPareto DeletePareto", err)
		tools.CustomError(ctx, err, 0, "Bind")
	}

	fmt.Println(data)

	err := h.UseCase.DeletePareto(&data)
	if err != nil {
		fmt.Println("HandlerPareto DeletePareto", err)
		tools.CustomError(ctx, err, 1, "DeletePareto")
	}

	return ctx.JSONBlob(http.StatusOK, nil)
}
