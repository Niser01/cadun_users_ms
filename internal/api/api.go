package api

import (
	"github.com/Niser01/CADUN_Users/tree/main/cadun_users_ms/internal/views"
	"github.com/labstack/echo/v4"
)

// It contains the view and the echo instance

type API struct {
	view views.View_interface
}

func New(view views.View_interface) *API {
	return &API{view: view}
}

func (a *API) Start(e *echo.Echo, address string) error {
	a.RegisterRoutes(e)
	return e.Start(address)
}
