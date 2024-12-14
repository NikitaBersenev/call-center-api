package main

import (
	"call-center/internal/app"
	"github.com/joho/godotenv"
	"log/slog"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading settings")
	}

	app.Run()
}
