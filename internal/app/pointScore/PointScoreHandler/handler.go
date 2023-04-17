package PointScoreHandler

import (
	"encoding/json"
	"fmt"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/models"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/pointScore"

	"gitgub.com/diploma-mppr/backend_mppr/tools"
	"github.com/labstack/echo/v4"
	"net/http"
	"reflect"
	"strconv"
)

type HandlerPointScore struct {
	UseCase pointScore.UseCase
}

func NewHandlerPointScore(useCase pointScore.UseCase) *HandlerPointScore {
	return &HandlerPointScore{
		UseCase: useCase,
	}
}

func (h HandlerPointScore) GetPointScore(ctx echo.Context) error {
	id := ctx.QueryParam("id")
	fmt.Println("Param: ", id, " ", reflect.TypeOf(id))
	che, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("HandlerPareto GetPointScore", err)
		return tools.CustomError(ctx, err, 0, "ParseInt")
	}

	data, err := h.UseCase.GetPointScore(int(che))
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

func (h HandlerPointScore) SetPointScore(ctx echo.Context) error {
	data := models.PointScoreJson{}
	if err := ctx.Bind(&data); err != nil {
		fmt.Println("HandlerPareto SetPointScore", err)
		tools.CustomError(ctx, err, 0, "Bind")
	}

	fmt.Println(data)

	task, err := h.UseCase.SetPointScore(&data)
	if err != nil {
		fmt.Println("HandlerPareto SetPointScore", err)
		tools.CustomError(ctx, err, 1, "Usecase")
	}

	result, err := json.Marshal(task)
	if err != nil {
		fmt.Println("HandlerPareto SetPointScore", err)
		tools.CustomError(ctx, err, 2, "Marshal")
	}

	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerPointScore) UpdatePointScore(ctx echo.Context) error {
	data := models.PointScoreJson{}
	if err := ctx.Bind(&data); err != nil {
		fmt.Println("HandlerPareto UpdatePointScore", err)
		tools.CustomError(ctx, err, 0, "Bind")
	}

	fmt.Println(data)

	task, err := h.UseCase.UpdatePointScore(&data)
	if err != nil {
		fmt.Println("HandlerPareto UpdatePointScore", err)
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

func (h HandlerPointScore) DeletePointScore(ctx echo.Context) error {
	data := models.PointScoreJson{}
	if err := ctx.Bind(&data); err != nil {
		fmt.Println("HandlerPareto DeletePointScore", err)
		tools.CustomError(ctx, err, 0, "Bind")
	}

	fmt.Println(data)

	err := h.UseCase.DeletePointScore(&data)
	if err != nil {
		fmt.Println("HandlerPareto DeletePointScore", err)
		tools.CustomError(ctx, err, 1, "DeletePointScore")
	}

	return ctx.JSONBlob(http.StatusOK, nil)
}
