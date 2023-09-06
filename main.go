package main

import (
	"fmt"
	"log"
	"net"

	"github.com/labstack/echo/v4"
	"github.com/uchupx/kajian-auth/config"
	"github.com/uchupx/kajian-auth/internal"
	"google.golang.org/grpc"
)

func main() {
	e := echo.New()
	e.Debug = true
	conf := config.GetConfig()
	i := &internal.Internal{}

	// initial logger
	i.InitLogger(conf)

	go runGRPCServer(conf, i)
	runAPIServer(conf, e, i)
}

func runAPIServer(conf *config.Config, e *echo.Echo, i *internal.Internal) {
	i.InitRoutes(conf, e)
	if err := e.Start(":" + conf.App.Port); err != nil {
		e.Logger.Fatal(err.Error())
	}
}

func runGRPCServer(conf *config.Config, i *internal.Internal) {
	listener, err := net.Listen("tcp", ":"+conf.App.GRPCPort)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	i.InitRoutesGRPC(conf, s)
	fmt.Println("Server is running at port 8081")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
