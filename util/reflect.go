package util

import "reflect"

// IsErrType 判断类型是否是 error
func IsErrType(t reflect.Type) bool {
	errorInterface := reflect.TypeOf((*error)(nil)).Elem()
	return t.Implements(errorInterface)
}

// MethodCall 函数调用
func MethodCall() {

}
