package tools

import (
	"encoding/json"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/models"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

func CustomError(ctx echo.Context, err error, number int, comment string) error {
	if err == nil {
		err = errors.Errorf("")
	}
	che := models.CustomError{
		Number:  number,
		Comment: comment,
		Error:   err.Error(),
	}
	result, _ := json.Marshal(che)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusInternalServerError, result)
}
