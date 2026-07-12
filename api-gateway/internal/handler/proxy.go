package handler

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
)

func ReverseProxy(target, prefix string) echo.HandlerFunc {

	u, _ := url.Parse(target)

	proxy := httputil.NewSingleHostReverseProxy(u)

	originalDirector := proxy.Director

	proxy.Director = func(req *http.Request) {
		originalDirector(req)

		// Remove the gateway prefix
		req.URL.Path = strings.TrimPrefix(req.URL.Path, prefix)

		if req.URL.Path == "" {
			req.URL.Path = "/"
		}
	}

	return func(c echo.Context) error {
		proxy.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}
