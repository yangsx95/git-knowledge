package dao

import (
	"context"
	"time"
)

func initContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 60*time.Second)
}
