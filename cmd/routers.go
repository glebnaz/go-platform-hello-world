package main

import (
	phttp "github.com/glebnaz/go-platform/http"
	"github.com/labstack/echo"
	"net/http"
)

var routers = []phttp.Router{
	{
		Path: "/hello",
		Handler: func(c echo.Context) error {
			return c.String(200, "hello")
		},
		Method: http.MethodGet,
	},
}
