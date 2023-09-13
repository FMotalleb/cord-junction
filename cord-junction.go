package main

import (
	"fmt"
	"net/http/httputil"
	"net/url"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	routerGroup := e.Group("/")
	routerGroup.Use(func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			req := context.Request()
			res := context.Response().Writer
			url, _ := url.Parse(fmt.Sprintf("%s://%s%s", "https", req.Host, req.RequestURI))

			println(url.Scheme)
			url.Scheme = "https"
			proxy := httputil.NewSingleHostReverseProxy(url)
			//trim reverseProxyRoutePrefix
			path := req.URL.Path
			req.URL.Path = path

			// ServeHttp is non blocking and uses a go routine under the hood
			proxy.ServeHTTP(res, req)
			return nil
		}
	})
	e.Start(":3390")
}
