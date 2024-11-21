package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/uchupx/kajian-auth/src/middlware"
)

type UserHandler struct {
}

func (u UserHandler) InitRoutes(e *echo.Echo, m *middlware.Middleware) {
	// e.POST("/user", u.AddUser)
	e.GET("/user", u.GetUser)
	e.GET("/user/me", u.GetMe, m.Authorization)
	e.PUT("/user", u.UpdateUser)
	e.DELETE("/user", u.DeleteUser)
}

func (u UserHandler) GetUser(c echo.Context) error {
	return c.JSON(200, "get user")
}

func (u UserHandler) GetMe(c echo.Context) error {
	return c.JSON(200, "get me")
}

func (u UserHandler) UpdateUser(c echo.Context) error {
	return c.JSON(200, "update user")
}

func (u UserHandler) DeleteUser(c echo.Context) error {
	return c.JSON(200, "delete user")
}
