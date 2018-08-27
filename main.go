package main

import (
	"context"
	"net/http"

	"github.com/chapsuk/mserv"
	"github.com/im-kulikov/helium"
	"github.com/im-kulikov/helium/grace"
	"github.com/im-kulikov/helium/logger"
	"github.com/im-kulikov/helium/module"
	"github.com/im-kulikov/helium/settings"
	"github.com/im-kulikov/helium/web"
	"github.com/labstack/echo"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

type (
	app struct {
		dig.In

		Logger  *zap.SugaredLogger
		Servers mserv.Server
	}
)

func (a app) Run(ctx context.Context) error {
	a.Logger.Infow("run servers")
	a.Servers.Start()

	<-ctx.Done()

	a.Logger.Infow("stop servers")
	a.Servers.Stop()

	return nil
}

var mod = module.Module{
	{Constructor: newApp},
	{Constructor: newHandler},
	{Constructor: web.NewBinder},
	{Constructor: web.NewValidator},
	{Constructor: web.NewEngine},
}.
	Append(grace.Module).
	Append(logger.Module).
	Append(settings.Module).
	Append(web.ServersModule)

func newApp(instance app) helium.App {
	return instance
}

func newHandler(e *echo.Echo) http.Handler {
	return e
}

func main() {
	h, err := helium.New(&settings.App{
		File:         "",
		Type:         "",
		Name:         "",
		BuildTime:    "",
		BuildVersion: "",
	}, mod)

	if err != nil {
		panic(err)
	}

	if err = h.Run(); err != nil {
		panic(err)
	}
}
