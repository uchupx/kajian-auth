package handler

import "github.com/labstack/echo/v4"

type BaseHandler interface {
	InitRoutes(e *echo.Echo)
}
