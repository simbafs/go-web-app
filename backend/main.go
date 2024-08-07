package main

import (
	"backend/api"
	"backend/internal/assestserver"
	"backend/internal/log"
	"backend/internal/tree"
	"embed"
	"fmt"

	"github.com/gin-gonic/gin"
	flag "github.com/spf13/pflag"
)

// go embed ignore files begin with '_' or '.'. Adding 'all:' in comment tells go embed to include all files

//go:embed all:static/*
var assestFS embed.FS

var (
	Mode       = "debug"
	Version    = "dev"
	CommitHash = "n/a"
	BuildTime  = "n/a"
)

var logger = log.New("main")

func run(addr string) error {
	gin.SetMode(Mode)
	r := gin.Default()

	api.Route(r)
	r.Use(assestserver.New(assestserver.CD(assestFS, "static")))

	logger.Printf("Server is running at %s\n", addr)
	return r.Run(addr)
}

func main() {
	addr := flag.StringP("addr", "a", ":3000", "server address")
	version := flag.BoolP("version", "v", false, "show version")
	flag.StringVarP(&Mode, "mode", "m", Mode, "server mode")
	list := flag.BoolP("list", "l", false, "list all files in static folder")
	flag.Parse()

	if *version {
		fmt.Printf("Version: %s\nCommitHash: %s\nBuildTime: %s\n", Version, CommitHash, BuildTime)
		return
	}

	if *list {
		dirs, err := tree.Tree(assestFS)
		if err != nil {
			logger.Printf("Oops, there's an error: %v\n", err)
			return
		}
		fmt.Println(dirs)
		return
	}

	if err := run(*addr); err != nil {
		logger.Printf("Oops, there's an error: %v\n", err)
	}
}
