package util

import (
	"fmt"
	"time"
)


// logOperation logs an operation to stdout
func LogOperation(operation string, accountID int) {
    fmt.Printf("[%s] Operation %s on account %d\n", time.Now().Format(time.RFC3339), operation, accountID)
}