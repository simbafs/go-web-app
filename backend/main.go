package main

import (
	"backend/api"
	"backend/internal/fileserver"
	"backend/internal/staticfs"
	"backend/pkg/websocket"
	"embed"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	flag "github.com/spf13/pflag"
)

// go embed ignore files begin with '_' or '.', 'all:' tells go embed to embed all files

//go:embed all:static/*
var rawStatic embed.FS

var static = staticfs.NewStatic(rawStatic, "static")

var (
	Mode       = "debug"
	Version    = "dev"
	CommitHash = "n/a"
	BuildTime  = "n/a"
)

var logger = log.New(gin.DefaultWriter, "[main] ", log.LstdFlags|log.Lmsgprefix)

func run(addr string) error {
	gin.SetMode(Mode)
	r := gin.Default()

	io := websocket.Route(r)
	api.Route(r, io)
	fileserver.Route(r, static, Mode)

	logger.Printf("Server is running at %s\n", addr)
	return r.Run(addr)
}

func main() {
	addr := flag.StringP("addr", "a", ":3000", "server address")
	version := flag.BoolP("version", "v", false, "show version")
	flag.StringVarP(&Mode, "mode", "m", Mode, "server mode")
	flag.Parse()

	if *version {
		fmt.Printf("Version: %s\nCommitHash: %s\nBuildTime: %s\n", Version, CommitHash, BuildTime)
		return
	}

	if err := run(*addr); err != nil {
		logger.Printf("Oops, there's an error: %v\n", err)
	}
}
