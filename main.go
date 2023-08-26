package main

import (
	"github.com/labstack/echo/v4"
	"github.com/uchupx/kajian-auth/config"
	"github.com/uchupx/kajian-auth/internal"
)

func main() {
	e := echo.New()
	e.Debug = true
	i := &internal.Internal{}
	conf := config.GetConfig()

	i.InitRoutes(conf, e)
	if err := e.Start(":8081"); err != nil {
		e.Logger.Fatal(err.Error())
	}
}
