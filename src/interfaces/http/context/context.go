package context

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"shoeshop-backend/src/infrastructure/errors"
	"shoeshop-backend/src/infrastructure/logger"
)

type (
	ApplicationContext struct {
		echo.Context
	}

	Success struct {
		Code    string      `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	Failed struct {
		Code    string      `json:"code"`
		Message string      `json:"message"`
		Error   string      `json:"error"`
		Data    interface{} `json:"data"`
	}
)

func (sc *ApplicationContext) Success(data interface{}) error {
	hc := http.StatusOK
	if data == nil {
		data = struct{}{}
	}

	res := Success{
		Code:    "00",
		Message: "success",
		Data:    data,
	}

	reqTime := sc.Context.Get("RequestTime").(time.Time)
	sc.Context.Logger().Info("Outgoing",
		logger.ToField("rt", fmt.Sprint(time.Since(reqTime).Milliseconds(), " ms")),
		logger.ToField("response", res),
		logger.ToField("http_code", hc))

	return sc.JSON(hc, res)
}

func (sc *ApplicationContext) SuccessWithMeta(data, meta interface{}) error {
	hc := http.StatusOK
	res := Success{
		Code:    "00",
		Message: "success",
		Data:    data,
	}

	reqTime := sc.Context.Get("RequestTime").(time.Time)
	sc.Context.Logger().Info("Outgoing",
		logger.ToField("rt", fmt.Sprint(time.Since(reqTime).Milliseconds(), " ms")),
		logger.ToField("response", res),
		logger.ToField("http_code", hc))

	return sc.JSON(hc, res)
}

func (sc *ApplicationContext) Fail(err error) error {
	return sc.FailWithData(err, nil)
}

func (sc *ApplicationContext) FailWithData(err error, data interface{}) error {
	var (
		ed = errors.ExtractError(err)
	)

	if data == nil {
		data = struct{}{}
	}

	res := Failed{
		Code:    ed.Code,
		Message: ed.Message,
		Error:   ed.FullMessage,
		Data:    data,
	}

	reqTime := sc.Context.Get("RequestTime").(time.Time)
	sc.Context.Logger().Info("Outgoing",
		logger.ToField("rt", fmt.Sprint(time.Since(reqTime).Milliseconds(), " ms")),
		logger.ToField("response", res),
		logger.ToField("http_code", ed.HttpCode))

	return sc.JSON(ed.HttpCode, res)
}

func (sc *ApplicationContext) Raw(hc int, data interface{}) error {
	if data == nil {
		data = struct{}{}
	}

	reqTime := sc.Context.Get("RequestTime").(time.Time)
	sc.Context.Logger().Info("Outgoing",
		logger.ToField("rt", fmt.Sprint(time.Since(reqTime).Milliseconds(), " ms")),
		logger.ToField("response", data),
		logger.ToField("http_code", hc))

	return sc.JSON(hc, data)
}

func NewEmptyApplicationContext(parent echo.Context) *ApplicationContext {
	return &ApplicationContext{parent}
}

func NewApplicationContext(parent echo.Context) (*ApplicationContext, error) {
	pctx, ok := parent.(*ApplicationContext)
	if !ok {
		return nil, errors.ErrSession
	}

	return pctx, nil
}
