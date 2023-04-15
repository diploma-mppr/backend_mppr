package task

import (
	"encoding/json"
	"fmt"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/models"
	"gitgub.com/diploma-mppr/backend_mppr/tools"

	"github.com/labstack/echo/v4"
	"net/http"
	"reflect"
	"strconv"
)

type HandlerTask struct {
	UseCase UseCase
}

func NewHandlerTask(ucase UseCase) *HandlerTask {
	return &HandlerTask{
		UseCase: ucase,
	}
}

func (h HandlerTask) GetData(ctx echo.Context) error {
	id := ctx.QueryParam("id")
	fmt.Println("Param: ", id, " ", reflect.TypeOf(id))
	che, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return tools.CustomError(ctx, err, 0, "ParseInt")
	}

	task, err := h.UseCase.GetData(int(che))
	if err != nil {
		fmt.Println(err.Error())
	}

	result, _ := json.Marshal(task)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerTask) SetData(ctx echo.Context) error {
	data := models.DataJson{}
	if err := ctx.Bind(&data); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(data)

	task, err := h.UseCase.SetData(&data)
	if err != nil {
		fmt.Println(err.Error())
	}

	result, _ := json.Marshal(task)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}
