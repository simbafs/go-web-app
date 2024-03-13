//go:build !dev

package assestserver

import (
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func New(assestFS fs.FS) gin.HandlerFunc {
	return gin.WrapH(http.FileServer(http.FS(assestFS)))
}
