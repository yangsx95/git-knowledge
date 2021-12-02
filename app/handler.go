package app

import (
	"errors"
	"git-knowledge/api/v1/vo"
	"git-knowledge/logger"
	"git-knowledge/result"
	"git-knowledge/util"
	"git-knowledge/ws"
	ut "github.com/go-playground/universal-translator"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
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

func (a *App) WebsocketHandler(handlers map[string]interface{}) func(c echo.Context) error {
	validateWebsocketHandler(handlers)

	return func(c echo.Context) (err error) {
		upgrader := websocket.Upgrader{}
		var (
			wsConn     *websocket.Conn
			connection *ws.Connection
		)
		if wsConn, err = upgrader.Upgrade(c.Response(), c.Request(), nil); err != nil {
			return err
		}

		// 创建Connection对象
		connection = ws.InitConnection(wsConn)
		defer connection.Close()

		// 用户上线
		a.wsConnMap.Online("", connection)

		// 循环读取并处理消息
		for {
		ReceiveNew:
			receiveMsg := result.NewEmptyMessage()
			if err = connection.ReadJSON(receiveMsg); err != nil {
				logger.Error("websocket读取消息出错, %s", err)
				if err = connection.SendJSON(result.NewErrorMessage("", errors.New("消息格式有误"))); err != nil {
					return
				}
				continue
			}

			var (
				handler    interface{}
				ok         bool
				serializer result.Serializable
			)

			// 获取处理器
			if handler, ok = handlers[receiveMsg.Func]; !ok {
				// 这里如果发送消息出错，就代表连接有问题，中断循环，且与客户端连接中断
				if err = connection.SendJSON(result.NewErrorMessage("", errors.New("未知的消息操作"))); err != nil {
					return
				}
				continue
			}

			// 获取消息序列化器
			if serializer = result.GetSerializerByContentType(receiveMsg.ContentType); serializer == nil {
				if err = connection.SendJSON(result.NewErrorMessage("", errors.New("无法解析的消息类型"))); err != nil {
					return
				}
				continue
			}

			// 调用处理器完成消息处理
			mT := reflect.TypeOf(handler)
			mV := reflect.ValueOf(handler)
			invokeParams := make([]reflect.Value, 0)

			var v reflect.Value
			for i := 0; i < mT.NumIn(); i++ {
				if v, err = process(mT.In(i), receiveMsg, connection, serializer); err != nil {
					if err = connection.SendJSON(result.NewErrorMessage("", errors.New("消息内容解析失败，请确保消息合法"))); err != nil {
						return
					}
					goto ReceiveNew
				}
				invokeParams = append(invokeParams, v)
			}
			// 调用处理器
			rVs := mV.Call(invokeParams)
			for _, rV := range rVs {
				if !rV.IsNil() {
					break
				}
			}
		}
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
func process(t reflect.Type, receiveMsg *result.Message, connection *ws.Connection, serializer result.Serializable) (v reflect.Value, err error) {
	v = reflect.New(t) // new方法返回的是指针
	switch t.Kind() {
	case reflect.String:
		v.Elem().Set(reflect.ValueOf(receiveMsg.Content))
		v = v.Elem()
	case reflect.Struct, reflect.Slice:
		switch t.String() {
		case "vo.WebsocketSender":
			v.Elem().Set(reflect.ValueOf(vo.WebsocketSender{Conn: connection, Func: receiveMsg.Func, ContentType: receiveMsg.ContentType}))
			v = v.Elem()
		default:
			if err = serializer.UnSerialize([]byte(receiveMsg.Content), v.Interface()); err != nil {
				return
			}
			v = v.Elem()
		}
	case reflect.Ptr:
		v = reflect.New(t.Elem())
		var vv reflect.Value
		if vv, err = process(t.Elem(), receiveMsg, connection, serializer); err != nil {
			return
		}
		v.Elem().Set(vv)
	default:
	}
	return
}
