package util

import (
	"log"
	"time"
)

// logOperation logs an operation to stdout
func LogOperation(operation string, accountID int) {
	log.Printf("Time.now [%s] Operation [%s] on account [%d]\n", time.Now().Format(time.RFC3339), operation, accountID)
}
