package app

import (
	"git-knowledge/result"
	"git-knowledge/util"
	ut "github.com/go-playground/universal-translator"
	"github.com/labstack/echo/v4"
	"reflect"
)

func (a *App) Handler(apiMethod interface{}) func(context echo.Context) error {
	validateHandlerMethod(apiMethod)

	mT := reflect.TypeOf(apiMethod)
	mV := reflect.ValueOf(apiMethod)

	// 构造gin路由处理函数
	return func(context echo.Context) error {
		translator, ok := a.ut.GetTranslator("zh")
		if !ok {
			translator, _ = a.ut.GetTranslator("zh")
		}

		// 准备参数对象列表
		pVs := make([]reflect.Value, 0)
		for i := 0; i < mT.NumIn(); i++ {
			// 当前参数的类型和构造的值
			pT := mT.In(i)
			pV := reflect.New(pT)

			// 如果是指针，则需要构造struct，并指向该指针
			if pT.Kind() == reflect.Ptr {
				structV := reflect.New(pT.Elem())
				pV.Elem().Set(structV)
			}
			// 将请求信息绑定到参数对象中
			err := context.Bind(pV.Interface())
			if err != nil {
				return err
			}
			reqVal := pV.Elem()
			// 校验结构体
			err = context.Validate(reqVal.Interface())
			if err != nil {
				return err
			}
			pVs = append(pVs, reqVal)
		}

		// 调用函数/方法
		rts := mV.Call(pVs)

		// 根据函数返回值生成结果，并返回响应体
		response := generateResult(a.errorHandler, &translator, rts)
		err := context.JSON(200, response)
		return err
	}
}

// validateHandlerMethod 校验传入方法是否是一个合法的handler方法
func validateHandlerMethod(apiMethod interface{}) {
	if apiMethod == nil {
		panic("handler处理器不可以为nil")
	}

	mT := reflect.TypeOf(apiMethod)

	// handler必须是一个函数或者方法
	if mT.Kind() != reflect.Func {
		panic("handler处理器只能是方法" + mT.Name())
	}

	// 参数大于1个
	pLen := mT.NumIn()
	if pLen > 1 {
		panic("handler处理器方法至多只能有一个参数" + mT.Name())
	}

	// 参数类别必须为结构体
	for i := 0; i < pLen; i++ {
		pT := mT.In(i)
		var tT reflect.Type
		if pT.Kind() == reflect.Ptr {
			tT = pT.Elem()
		}
		if tT.Kind() != reflect.Struct {
			panic("handler处理器方法的参数必须为struct类型" + mT.Name())
		}
	}

	// 返回值数量至多为2
	rLen := mT.NumOut()
	if rLen > 2 {
		panic("handler处理器方法至多只能有两个返回值" + mT.Name())
	}

	// 返回值类别只能是struct或者error类型，并且数量要 <=2，切返回值类别不能相同
	for i := 0; i < rLen; i++ {
		rT := mT.Out(i)
		if rT.Kind() == reflect.Ptr {
			rT = rT.Elem()
		}
		if rT.Kind() != reflect.Struct && !util.IsErrType(rT) {
			panic("handler处理器方法的返回值必须为struct或者error类型" + mT.Name())
		}

		if rLen == 2 &&
			(i == 0 && rT.Kind() != reflect.Struct) &&
			(i == 1 && util.IsErrType(rT)) {
			panic("handler处理器方法的多返回值格式必须为 (struct,error)" + mT.Name())
		}
	}

}

func generateResult(handler *result.ErrorHandler, translator *ut.Translator, rts []reflect.Value) *result.Response {
	if len(rts) == 0 { // 无返回值
		return result.Build(result.CodeOk)
	}

	if len(rts) == 1 { // 单返回值
		rv := rts[0].Interface()
		// error
		if util.IsErrType(rts[0].Type()) && rv != nil {
			return handler.Handler(rv.(error), translator)
		}
		// success
		return result.Build(result.CodeOk).WithData(rv)
	} else { // 双返回值
		rv0 := rts[0].Interface()
		rv1 := rts[1].Interface().(error)
		if rv1 != nil {
			return handler.Handler(rv1.(error), translator)
		}
		return result.Build(result.CodeOk).WithData(rv0)
	}
}
