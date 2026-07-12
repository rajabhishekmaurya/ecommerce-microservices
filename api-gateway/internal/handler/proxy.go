package handler

import (
	"net/http/httputil"
	"net/url"

	"github.com/labstack/echo/v4"
)

func ReverseProxy(target string) echo.HandlerFunc {

	u, _ := url.Parse(target)

	proxy := httputil.NewSingleHostReverseProxy(u)

	return func(c echo.Context) error {
		proxy.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}
