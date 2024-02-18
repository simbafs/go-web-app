package fileserver

// This package implements a route that directs incoming traffic to different destinations based on whether it's in development or release mode.
// In development mode, it uses httputil.ReverseProxy. In release mode, it utilizes an embedded file server.

import (
	"io/fs"
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

func FileServer(assestFS fs.FS, mode string) gin.HandlerFunc {
	// https://stackoverflow.com/questions/36357791/
	if mode == gin.DebugMode {
		return proxy
	} else {
		return gin.WrapH(http.FileServer(http.FS(assestFS)))
	}
}

// https://github.com/golang/go/issues/43431#issuecomment-752662261

func CD(embedFS fs.FS, root string) fs.FS {
	newFS, err := fs.Sub(embedFS, root)
	if err != nil {
		panic(err)
	}
	return newFS
}
