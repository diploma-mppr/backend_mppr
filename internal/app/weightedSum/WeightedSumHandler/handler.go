package WeightedSumHandler

import (
	"encoding/json"
	"fmt"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/models"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/weightedSum"
	"gitgub.com/diploma-mppr/backend_mppr/tools"
	"github.com/labstack/echo/v4"
	"net/http"
	"reflect"
	"strconv"
)

type HandlerWeightedSum struct {
	UseCase weightedSum.UseCase
}

func NewHandlerWeightedSum(useCase weightedSum.UseCase) *HandlerWeightedSum {
	return &HandlerWeightedSum{
		UseCase: useCase,
	}
}

func (h HandlerWeightedSum) GetWeightedSum(ctx echo.Context) error {
	id := ctx.QueryParam("id")
	fmt.Println("Param: ", id, " ", reflect.TypeOf(id))
	che, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("HandlerWeightedSum GetWeightedSum", err)
		return tools.CustomError(ctx, err, 0, "ParseInt")
	}

	data, err := h.UseCase.GetWeightedSum(int(che))
	if err != nil {
		fmt.Println("HandlerWeightedSum GetWeightedSum", err)
		return tools.CustomError(ctx, err, 1, "UseCase")
	}

	result, err := json.Marshal(data)
	if err != nil {
		fmt.Println("HandlerWeightedSum GetWeightedSum", err)
		return tools.CustomError(ctx, err, 2, "Marshal")
	}

	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerWeightedSum) SetWeightedSum(ctx echo.Context) error {
	data := models.WeightedSumJson{}
	if err := ctx.Bind(&data); err != nil {
		fmt.Println("HandlerWeightedSum SetWeightedSum", err)
		tools.CustomError(ctx, err, 0, "Bind")
	}

	fmt.Println(data)

	task, err := h.UseCase.SetWeightedSum(&data)
	if err != nil {
		fmt.Println("HandlerWeightedSum SetWeightedSum", err)
		tools.CustomError(ctx, err, 1, "Usecase")
	}

	result, err := json.Marshal(task)
	if err != nil {
		fmt.Println("HandlerWeightedSum SetWeightedSum", err)
		tools.CustomError(ctx, err, 2, "Marshal")
	}

	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerWeightedSum) UpdateWeightedSum(ctx echo.Context) error {
	data := models.WeightedSumJson{}
	if err := ctx.Bind(&data); err != nil {
		fmt.Println("HandlerWeightedSum UpdateWeightedSum", err)
		tools.CustomError(ctx, err, 0, "Bind")
	}

	fmt.Println(data)

	task, err := h.UseCase.UpdateWeightedSum(&data)
	if err != nil {
		fmt.Println("HandlerWeightedSum UpdateWeightedSum", err)
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

func (h HandlerWeightedSum) DeleteWeightedSum(ctx echo.Context) error {
	data := models.WeightedSumJson{}
	if err := ctx.Bind(&data); err != nil {
		fmt.Println("HandlerWeightedSum DeleteWeightedSum", err)
		tools.CustomError(ctx, err, 0, "Bind")
	}

	fmt.Println(data)

	err := h.UseCase.DeleteWeightedSum(&data)
	if err != nil {
		fmt.Println("HandlerWeightedSum DeleteWeightedSum", err)
		tools.CustomError(ctx, err, 1, "DeleteWeightedSum")
	}

	return ctx.JSONBlob(http.StatusOK, nil)
}
