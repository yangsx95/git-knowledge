package conf

import (
	"fmt"
	"testing"
)

func TestGetConfig(t *testing.T) {
	fmt.Println(GetConfig().Log.Level)
	fmt.Println(GetConfig().Log.Dir)
}
