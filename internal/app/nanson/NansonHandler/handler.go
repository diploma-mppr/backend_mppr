package NansonHandler

import (
	"encoding/json"
	"fmt"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/models"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/nanson"
	"gitgub.com/diploma-mppr/backend_mppr/tools"
	"github.com/labstack/echo/v4"
	"net/http"
	"reflect"
	"strconv"
)

type HandlerNanson struct {
	UseCase nanson.UseCase
}

func NewHandlerNanson(useCase nanson.UseCase) *HandlerNanson {
	return &HandlerNanson{
		UseCase: useCase,
	}
}

func (h HandlerNanson) GetNanson(ctx echo.Context) error {
	id := ctx.QueryParam("id")
	fmt.Println("Param: ", id, " ", reflect.TypeOf(id))
	che, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("HandlerNanson GetNanson", err)
		return tools.CustomError(ctx, err, 0, "ParseInt")
	}

	data, err := h.UseCase.GetNanson(int(che))
	if err != nil {
		fmt.Println("HandlerNanson GetNanson", err)
		return tools.CustomError(ctx, err, 1, "UseCase")
	}

	result, err := json.Marshal(data)
	if err != nil {
		fmt.Println("HandlerNanson GetNanson", err)
		return tools.CustomError(ctx, err, 2, "Marshal")
	}

	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerNanson) SetNanson(ctx echo.Context) error {
	data := models.NansonJson{}
	if err := ctx.Bind(&data); err != nil {
		fmt.Println("HandlerNanson SetNanson", err)
		tools.CustomError(ctx, err, 0, "Bind")
	}

	fmt.Println(data)

	task, err := h.UseCase.SetNanson(&data)
	if err != nil {
		fmt.Println("HandlerNanson SetNanson", err)
		tools.CustomError(ctx, err, 1, "Usecase")
	}

	result, err := json.Marshal(task)
	if err != nil {
		fmt.Println("HandlerNanson SetNanson", err)
		tools.CustomError(ctx, err, 2, "Marshal")
	}

	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerNanson) UpdateNanson(ctx echo.Context) error {
	data := models.NansonJson{}
	if err := ctx.Bind(&data); err != nil {
		fmt.Println("HandlerNanson UpdateNanson", err)
		tools.CustomError(ctx, err, 0, "Bind")
	}

	fmt.Println(data)

	task, err := h.UseCase.UpdateNanson(&data)
	if err != nil {
		fmt.Println("HandlerNanson UpdateNanson", err)
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

func (h HandlerNanson) DeleteNanson(ctx echo.Context) error {
	data := models.NansonJson{}
	if err := ctx.Bind(&data); err != nil {
		fmt.Println("HandlerNanson DeleteNanson", err)
		tools.CustomError(ctx, err, 0, "Bind")
	}

	fmt.Println(data)

	err := h.UseCase.DeleteNanson(&data)
	if err != nil {
		fmt.Println("HandlerNanson DeleteNanson", err)
		tools.CustomError(ctx, err, 1, "DeleteNanson")
	}

	return ctx.JSONBlob(http.StatusOK, nil)
}
