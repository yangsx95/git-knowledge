package conf

import (
	"fmt"
	"testing"
)

func TestGetConfig(t *testing.T) {
	InitConfig("../git-knowledge.ini")
	fmt.Println(GetConfig().Log.Level)
	fmt.Println(GetConfig().Log.Dir)
	fmt.Println(GetConfig().Github.ClientId)
}
