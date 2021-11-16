package result

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

type Serializable interface {
	Serialize(target interface{}) ([]byte, error)
	UnSerialize(data []byte, target interface{}) error
}

type JSONSerializer struct {
}

func (j *JSONSerializer) Serialize(target interface{}) ([]byte, error) {
	return json.Marshal(target)
}

func (j *JSONSerializer) UnSerialize(data []byte, target interface{}) error {
	return json.Unmarshal(data, target)
}

func GetSerializerByContentType(contentType string) Serializable {
	switch contentType {
	case "application/json":
		return new(JSONSerializer)
	case "text/plain":
		return new(TextSerializer)
	}
	return nil
}

type TextSerializer struct {
}

func (t *TextSerializer) Serialize(target interface{}) ([]byte, error) {
	return []byte(fmt.Sprint(target)), nil
}

func (t *TextSerializer) UnSerialize(data []byte, target interface{}) error {
	err := validateTargetInterface(target)
	if err != nil {
		return err
	}
	// 参数target只能是 []byte 或者 string
	switch target.(type) {
	case *[]byte:
		reflect.ValueOf(target).Elem().Set(reflect.ValueOf(data))
	case []byte:
		reflect.ValueOf(target).Set(reflect.ValueOf(data))
	case *string:
		reflect.ValueOf(target).Elem().Set(reflect.ValueOf(string(data)))
	case string:
		reflect.ValueOf(target).Set(reflect.ValueOf(data))
	default:
		return errors.New("只支持将data反序列化为类型 string 与 []byte")
	}
	return nil
}

func validateTargetInterface(target interface{}) error {
	rv := reflect.ValueOf(target)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New("参数是nil或者不是指针")
	}
	return nil
}
