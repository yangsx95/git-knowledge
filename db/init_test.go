package db

import (
	"fmt"
	"testing"
)

func TestInitResource(t *testing.T) {
	resource, err := InitResource("127.0.0.1", "27017", "test", "root", "root123")
	if err != nil {
		panic(err)
	}
	cols := resource.DB.Collection("test")
	fmt.Println(cols)
}
