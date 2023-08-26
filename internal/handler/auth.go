package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/uchupx/kajian-auth/internal/dto"
	"github.com/uchupx/kajian-auth/internal/service"
)

type AuthHandler struct {
	UserService *service.UserService
}

func (a *AuthHandler) InitRoutes(e *echo.Echo) {
	e.POST("/token", a.Auth)
}

func (a *AuthHandler) Auth(c echo.Context) error {
	var req dto.AuthRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	res, err := a.UserService.Login(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(200, res)
}
