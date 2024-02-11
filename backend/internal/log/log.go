package log

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func New(scope string) *log.Logger {
	return log.New(gin.DefaultWriter, fmt.Sprintf("[%s] ", scope), log.LstdFlags|log.Lmsgprefix)
}
