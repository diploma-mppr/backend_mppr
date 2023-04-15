package task

import (
	"encoding/json"
	"fmt"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type HandlerTask struct {
	Ucase Ucase
}

func NewHandlerTask(ucase Ucase) *HandlerTask {
	return &HandlerTask{
		Ucase: ucase,
	}
}

//func (h HandlerTask) GetData(ctx echo.Context) error {
//	task, err := h.Ucase.GetData()
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//
//	result, _ := json.Marshal(task)
//	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
//	return ctx.JSONBlob(http.StatusOK, result)
//}

func (h HandlerTask) SetData(ctx echo.Context) error {
	data := models.DataJson{}
	if err := ctx.Bind(&data); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(data)

	task, err := h.Ucase.SetData(data)
	if err != nil {
		fmt.Println(err.Error())
	}

	result, _ := json.Marshal(task)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}
