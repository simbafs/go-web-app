//go:build dev

package assestserver

import (
	"io/fs"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func New(assestFS fs.FS) gin.HandlerFunc {
	remote, err := url.Parse("http://localhost:3001")
	if err != nil {
		panic(err)
	}
	return func(c *gin.Context) {
		proxy := httputil.NewSingleHostReverseProxy(remote)

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
