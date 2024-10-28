package main

import (
	"github.com/nitwhiz/svapi/internal/server"
	"log/slog"
	"os"
)

func main() {
	if err := server.Start(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
