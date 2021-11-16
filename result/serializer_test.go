package result

import (
	"fmt"
	"testing"
)

func TestTextSerializer_Serialize(t *testing.T) {
	serializer := new(TextSerializer)
	text, err := serializer.Serialize("abc")
	panicError(err)
	fmt.Println(string(text))

	type A struct {
		A string
	}
	text, err = serializer.Serialize(A{A: "aaa"})
	panicError(err)
	fmt.Println(string(text))
}

func TestTextSerializer_UnSerialize(t *testing.T) {
	serializer := new(TextSerializer)
	str := ""
	err := serializer.UnSerialize([]byte("abc"), &str)
	panicError(err)
	fmt.Println(str)
}

func panicError(err error) {
	if err != nil {
		panic(err)
	}
}
