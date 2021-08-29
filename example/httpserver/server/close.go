package server

import (
	"github.com/whereabouts/sdk/logger"
)

func Close() {
	// close db
	logger.Println("doing something on shutdown...")
}
