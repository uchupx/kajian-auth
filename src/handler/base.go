package handler

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/uchupx/kajian-api/pkg/logger"
)

type BaseHandler interface {
	InitRoutes(e *echo.Echo)
}

func HandleError(ctx echo.Context, err error) error {
	json_map := make(map[string]interface{})
	json.NewDecoder(ctx.Request().Body).Decode(&json_map)
	meta := map[string]interface{}{
		"body": json_map,
	}

	logger.Logger.WithFields(logrus.Fields{
		"method": ctx.Request().Method,
		"path":   *ctx.Request().URL,
		"meta":   meta,
	})

	return ctx.JSON(500, err)
}
