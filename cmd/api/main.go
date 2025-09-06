package main

import (
	"fmt"

	"github.com/jcbuitrago/anb-rising-stars/internal/platform/config"
	"github.com/jcbuitrago/anb-rising-stars/internal/platform/logger"
	sharederrors "github.com/jcbuitrago/anb-rising-stars/internal/shared/errors"
)

func main() {
	println("Initial API")

	cfg := config.Load()
	fmt.Printf("Env: %s, Port: %s\n", cfg.Env, cfg.Port)
	log := logger.New(cfg.LogLevel)
	log.Info("Starting server")
	log.Debug("This is a debug type message")

	err := sharederrors.NewValidationError("email", "invalid format")
	fmt.Println("Error: ", err)
	if sharederrors.IsValidationError(err) {
		fmt.Println("It is in facto a validation error")
	}
}
