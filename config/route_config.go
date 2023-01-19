package config

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	RouteConfig interface {
		AddRoute(e *echo.Echo)
	}
	routeConfig struct {
		Path string
		Json []byte
	}
)

func NewRouteConfig(p string, j []byte) RouteConfig {
	return routeConfig{p, j}
}

func (r routeConfig) AddRoute(e *echo.Echo) {
	var res interface{}
	err := json.Unmarshal(r.Json, &res)
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.GET(r.Path, func(c echo.Context) error {
		return c.JSON(http.StatusOK, res)
	})
}
