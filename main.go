package main

import (
	"context"
	"fmt"

	"github.com/Niser01/CADUN_Users/tree/main/cadun_users_ms/DB"
	"github.com/Niser01/CADUN_Users/tree/main/cadun_users_ms/internal/api"
	"github.com/Niser01/CADUN_Users/tree/main/cadun_users_ms/internal/views"
	"github.com/Niser01/CADUN_Users/tree/main/cadun_users_ms/settings"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

// main is the entry point of the application

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			DB.New,
			views.New,
			api.New,
			echo.New,
		),
		fx.Invoke(setLifeCycle),
	)

	app.Run()
}

// setLifeCycle is used to set the lifecycle of the application and start the server on a goroutine so that it doesn't block the main thread of execution

func setLifeCycle(lc fx.Lifecycle, a *api.API, s *settings.Settings, e *echo.Echo) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			address := fmt.Sprintf(":%s", s.Port)
			go a.Start(e, address)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
