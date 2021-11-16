package app

import (
	"errors"
	"git-knowledge/api/v1/vo"
	"git-knowledge/logger"
	"git-knowledge/result"
	"git-knowledge/util"
	ut "github.com/go-playground/universal-translator"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
	"reflect"
)

func (a *App) Handler(apiMethod interface{}) func(context echo.Context) error {
	validateHandlerMethod(apiMethod)

	mT := reflect.TypeOf(apiMethod)
	mV := reflect.ValueOf(apiMethod)

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
				// 填充登录信息
				if context.Get("_userid") != "" {
					// FieldByName 如果没有找到对应的字段的Value将会返回Zero零值
					field := structV.Elem().FieldByName("LoginInfo")
					// 判断字段Value是否是零值，也就是判断字段是否存在
					if field.IsValid() {
						// 设置字段内容
						field.FieldByName("Userid").Set(reflect.ValueOf(context.Get("_userid")))
					}
				}
				pV.Elem().Set(structV)
			}
			reqVal := pV.Elem()
			// 将请求信息绑定到参数对象中
			err := context.Bind(reqVal.Interface())
			if err != nil {
				return err
			}
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

		var err error
		// 根据Accept头决定要返回的数据类型
		accept := context.Request().Header.Get("Accept")
		switch accept {
		case "application/json":
			err = context.JSON(200, response)
		case "application/xml":
			err = context.XML(200, response)
		default:
			err = context.JSON(200, response)
		}
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

	// 返回值类别只能是(data, error)或者 (data) 或者  (error)
	// 其中 data 可以是 struct 或者 slice 结构
	for i := 0; i < rLen; i++ {
		rT := mT.Out(i)
		if rT.Kind() == reflect.Ptr {
			rT = rT.Elem()
		}
		if rT.Kind() != reflect.Struct && rT.Kind() != reflect.Slice && !util.IsErrType(rT) {
			panic("handler处理器方法的返回值必须为struct、slice或者error类型" + mT.Name())
		}

		if rLen == 2 &&
			((i == 0 && (rT.Kind() != reflect.Struct) && (rT.Kind() != reflect.Slice)) ||
				(i == 1 && !util.IsErrType(rT))) {
			panic("handler处理器方法的多返回值格式必须为 (struct,error) 或者 (slice, error)" + mT.Name())
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
		rv1 := rts[1].Interface()
		if rv1 != nil {
			return handler.Handler(rv1.(error), translator)
		}
		return result.Build(result.CodeOk).WithData(rv0)
	}
}

func WebsocketHandler(handlers map[string]interface{}) func(c echo.Context) error {
	validateWebsocketHandler(handlers)

	return func(c echo.Context) error {
		websocket.Handler(func(ws *websocket.Conn) {
			defer ws.Close()
			// 不断接收客户端消息，当接收到消息，创建协程处理消息
			for {
				receiveMsg := result.NewEmptyMessage()
				err := websocket.JSON.Receive(ws, &receiveMsg)
				if err != nil {
					logger.Error("websocket消息解析出错, %s", err)
					_ = websocket.JSON.Send(ws, result.NewErrorMessage("", err))
					continue
				}

				// 获取处理器
				handler, ok := handlers[receiveMsg.Func]
				if !ok { // 未知操作
					_ = websocket.JSON.Send(ws, result.NewErrorMessage("", errors.New("未知的消息操作")))
					continue
				}

				// 获取消息序列化器
				serializer := result.GetSerializerByContentType(receiveMsg.ContentType)
				if serializer == nil {
					_ = websocket.JSON.Send(ws, result.NewErrorMessage("", errors.New("不支持的消息类型")))
					continue
				}

				// 调用处理器完成消息处理
				mT := reflect.TypeOf(handler)
				mV := reflect.ValueOf(handler)
				// 准备调用处理器方法所需的参数
				invokeParams := make([]reflect.Value, 0)
				for i := 0; i < mT.NumIn(); i++ {
					inT := mT.In(i) // 参数类型
					if inT.String() == "vo.WebsocketSender" {
						// 构造websocketSender
						wss := vo.WebsocketSender{Conn: ws, Func: receiveMsg.Func, ContentType: receiveMsg.ContentType}
						invokeParams = append(invokeParams, reflect.ValueOf(wss))
					} else if inT.Kind() == reflect.String {
						inV := reflect.New(inT)
						inV.Set(reflect.ValueOf(receiveMsg.Content))
						invokeParams = append(invokeParams, inV)
					} else if inT.Kind() == reflect.Struct || inT.Kind() == reflect.Slice {
						//err := serializer.UnSerialize([]byte(receiveMsg.Content), inV.Interface())
						//if err != nil {
						//	_ = websocket.JSON.Send(ws, result.NewErrorMessage("", err))
						//}
					} else if inT.Kind() == reflect.Ptr { // 指针类型
						inV := reflect.New(inT)
						invokeParams = append(invokeParams, inV.Elem())
						elemT := inT.Elem()
						elemV := reflect.New(elemT)
						if elemT.Kind() == reflect.String {
							elemV.Elem().Set(reflect.ValueOf(receiveMsg.Content))
						} else if elemT.Kind() == reflect.Struct || elemT.Kind() == reflect.Slice {
							err := serializer.UnSerialize([]byte(receiveMsg.Content), elemV.Interface())
							if err != nil {
								_ = websocket.JSON.Send(ws, result.NewErrorMessage("", err))
							}
						}

						inV.Elem().Set(elemV)
					}
				}
				// 调用处理器
				rVs := mV.Call(invokeParams)
				for _, rV := range rVs {
					if !rV.IsNil() {
						_ = websocket.JSON.Send(ws, result.NewErrorMessage("", rV.Interface().(error)))
						break
					}
				}
			}
		}).ServeHTTP(c.Response(), c.Request())
		return nil
	}
}

func validateWebsocketHandler(handlers map[string]interface{}) {
	// 校验handlers
	for _, apiMethod := range handlers {
		if apiMethod == nil {
			panic("handler处理器不可以为nil")
		}

		mT := reflect.TypeOf(apiMethod)

		// handler必须是一个函数或者方法
		if mT.Kind() != reflect.Func {
			panic("handler处理器只能是方法" + mT.Name())
		}

		// 返回值只许是error类型
		if !(mT.NumOut() == 0 || (mT.NumOut() == 1 && util.IsErrType(mT.Out(0)))) {
			panic("handler处理器方法的返回值只能是错误类型" + mT.Name())
		}

		// 请求参数数量至多为2
		pLen := mT.NumIn()
		if pLen > 2 {
			panic("handler处理器方法至多只能有两个请求参数" + mT.Name())
		}

		// 请求参数类型只能是 vo.WebsocketSender（也是一个struct）, string/*string、slice/*slice、struct/*struct 等类型
		for i := 0; i < pLen; i++ {
			pT := mT.In(i)
			if pT.Kind() == reflect.Ptr {
				pT = pT.Elem()
			}
			if pT.Kind() != reflect.Struct &&
				pT.Kind() != reflect.Slice &&
				pT.Kind() != reflect.String {
				panic("handler处理器方法的参数类型只能为为 *websocket.Conn、string、slice、struct类型 " + mT.Name())
			}
		}
	}
}
