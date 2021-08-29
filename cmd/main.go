package main

import (
	"context"
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	level, err := log.ParseLevel(os.Getenv("LOGLVL"))
	if err != nil {
		level = log.DebugLevel
	}
	log.SetLevel(level)
}

func main() {
	ctx := context.Background()

	app := newApp(ctx, nil)

	if err := app.Start(ctx); err != nil {
		log.Errorf("bad start app: %s", err)
	}
}
