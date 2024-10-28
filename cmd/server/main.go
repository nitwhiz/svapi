package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nitwhiz/svapi/internal/server"
	"log/slog"
	"os"
)

var isRelease = false

func main() {
	if isRelease {
		gin.SetMode(gin.ReleaseMode)
	}

	if err := server.Start(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
