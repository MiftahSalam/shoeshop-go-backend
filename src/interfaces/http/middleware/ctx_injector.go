package middleware

import (
	"bytes"
	"io/ioutil"
	"time"

	"github.com/labstack/echo/v4"

	"shoeshop-backend/src/infrastructure/logger"
)

const (
	ExternalID = "X-EXTERNAL-ID"
	JourneyID  = "X-JOURNEY-ID"
	ChainID    = "X-CHAIN-ID"
)

type (
	ContextInjectorMiddleware interface {
		Injector(next echo.HandlerFunc) echo.HandlerFunc
	}

	contextInjectorMiddleware struct {
		logger logger.Logger
	}
)

func (i *contextInjectorMiddleware) Injector(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			rid = c.Request().Header.Get(ExternalID)
		)
		if len(rid) == 0 {
			rid, _ = GenerateUUID()
		}

		// print request time
		var bodyBytes []byte
		if c.Request().Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
			// Restore the io.ReadCloser to its original state
			c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		c.Set("RequestTime", time.Now())

		if !i.skipper(c) {
			i.logger.Info("Incoming",
				logger.ToField("url", c.Request().URL.String()),
				logger.ToField("header", c.Request().Header),
				logger.ToField("request", string(bodyBytes)),
				logger.ToField("rid", rid))
		}

		return h(c)
	}
}

func (i *contextInjectorMiddleware) skipper(c echo.Context) (skip bool) {
	url := c.Request().URL.String()
	if url == "/" {
		skip = false
		return
	}

	return
}

func NewContextInjectorMiddleware(logger logger.Logger) (ContextInjectorMiddleware, error) {
	return &contextInjectorMiddleware{logger: logger}, nil
}
