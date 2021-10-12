package app

import (
	"fmt"
	"testing"
)

type TApi struct {
}

type TReq struct {
	name string
	age  int
}

func (t *TApi) tMethod() {
	fmt.Println("hello")
}

func (t *TApi) tMethodWithParam(req *TReq) {
	fmt.Println(req.name)
}

func TestHandler(t *testing.T) {
}
