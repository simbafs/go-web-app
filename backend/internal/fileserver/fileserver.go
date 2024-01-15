package fileserver
// This package implements a route that directs incoming traffic to different destinations based on whether it's in development or release mode.
// In development mode, it uses httputil.ReverseProxy. In release mode, it utilizes an embedded file server.

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func proxy(c *gin.Context) {
	remote, err := url.Parse("http://localhost:3001")
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)

	proxy.ServeHTTP(c.Writer, c.Request)
}

func Route(r *gin.Engine, static http.FileSystem, mode string) {
	// https://stackoverflow.com/questions/36357791/
	if mode == gin.DebugMode {
		r.Use(proxy)
	} else {
		r.NoRoute(gin.WrapH(http.FileServer(static)))
	}
}
