package app

import (
	"git-knowledge/result"
	"git-knowledge/util"
	"github.com/gin-gonic/gin"
	"reflect"
)

func Handler(apiMethod interface{}) func(context *gin.Context) {
	validateHandlerMethod(apiMethod)

	mT := reflect.TypeOf(apiMethod)
	mV := reflect.ValueOf(apiMethod)

	// 构造gin路由处理函数
	return func(context *gin.Context) {
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
			err := context.ShouldBind(pV.Interface())
			if err != nil {
				context.JSON(200, result.Build(result.CodeReqParamErr).WithDetail(err.Error()))
				return
			}
			pVs = append(pVs, pV.Elem())
		}

		// 调用函数/方法
		rts := mV.Call(pVs)

		// 根据函数返回值生成结果
		response := generateResult(rts)

		// 设置http响应体
		context.JSON(200, response)
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

func generateResult(rts []reflect.Value) *result.Response {
	if len(rts) == 0 { // 无返回值
		return result.Build(result.CodeOk)
	}

	if len(rts) == 1 { // 单返回值
		rv := rts[0].Interface()
		// error
		if util.IsErrType(rts[0].Type()) && rv != nil {
			return errToResp(rv.(error))
		}
		// success
		return result.Build(result.CodeOk).WithData(rv)
	} else { // 双返回值
		rv0 := rts[0].Interface()
		rv1 := rts[1].Interface().(error)
		if rv1 != nil {
			return errToResp(rv1)
		}
		return result.Build(result.CodeOk).WithData(rv0)
	}
}

func errToResp(err error) *result.Response {
	if err == nil {
		panic("err不能为nil")
	}
	var response *result.Response
	switch err.(type) {
	case result.ServiceError: // 如果是服务错误，使用服务码构建返回对象
		e := err.(result.ServiceError)
		response = result.Build(e.Code).WithDetail(e.Detail)
	default: // 如果是未知异常，则抛出系统内部错误
		response = result.Build(result.CodeInnerError).WithDetail(err.Error())
	}
	return response
}
