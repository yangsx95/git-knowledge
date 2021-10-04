package db

import (
	"fmt"
	"testing"
)

func InitResourceForTest() (*Resource, error) {
	return NewResource("127.0.0.1", "27017", "test", "root", "root123")
}

func TestInitResource(t *testing.T) {
	resource, err := InitResourceForTest()
	if err != nil {
		panic(err)
	}
	cols := resource.DB.Collection("test")
	fmt.Println(cols)
}
