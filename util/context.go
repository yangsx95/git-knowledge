package util

import (
	"context"
	"time"
)

func GetContextWithTimeout60Second() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 60*time.Second)
}
