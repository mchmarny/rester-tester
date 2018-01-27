package util

import (
	"log"
	"os"

	"github.com/google/uuid"
)

var (
	logger = log.New(os.Stdout, "[util] ", log.Lshortfile|log.Ldate|log.Ltime)
)

// GetUUIDv4 returns UUID v4
func GetUUIDv4() string {
	id, err := uuid.NewRandom()
	if err != nil {
		logger.Fatalf("Error while getting id: %v\n", err)
	}
	return id.String()
}
