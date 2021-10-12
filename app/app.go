package app

import (
	"git-knowledge/db"
	"git-knowledge/logger"
	"git-knowledge/middlewares"
	"git-knowledge/result"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"os"
)

// App 应用程序对象
type App struct {
	db           *db.Resource
	Dao          *Dao
	Api          *Api
	ut           *ut.UniversalTranslator
	validator    *validator.Validate
	errorHandler *result.ErrorHandler
	echo         *echo.Echo
}

func NewApp() *App {
	app := App{}
	// 加载配置文件
	loadConfig()
	// 初始化日志
	logger.InitLogger(os.Getenv("LOG_LEVEL"), os.Getenv("LOG_DIR"))
	// 初始化数据库
	app.db = initDb()
	// 初始化检验器
	app.initValidate()
	// 初始化翻译器
	app.initTranslator()
	// 初始化错误处理器
	app.initErrorHandler()
	// 初始化Dao组件
	app.Dao = initDao(&app)
	// 初始化Api组件
	app.Api = initApi(&app)
	// 初始化web引擎
	app.initEchoAndMiddleware()
	// 初始化路由
	app.initRouter()
	return &app
}

func loadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("加载配置文件.env出现错误")
	}
}

func initDb() *db.Resource {
	resource, err := db.NewResource(
		os.Getenv("MONGO_HOST"),
		os.Getenv("MONGO_PORT"),
		os.Getenv("MONGO_DATABASE"),
		os.Getenv("MONGO_USERNAME"),
		os.Getenv("MONGO_PASSWORD"),
	)
	if err != nil {
		logger.Fatal("连接mongodb出现错误", err)
	}
	return resource
}

func (a *App) Start() {
	err := a.echo.Start(":8080")
	if err != nil {
		logger.Fatal("启动服务出现错误 %s", err)
	}
}

type EchoValidator struct {
	validate *validator.Validate
}

func (e *EchoValidator) Validate(data interface{}) error {
	return e.validate.Struct(data)
}

func NewEchoValidator(validate *validator.Validate) *EchoValidator {
	return &EchoValidator{validate: validate}
}

func (a *App) initEchoAndMiddleware() {
	a.echo = echo.New()
	// 错误处理
	a.echo.HTTPErrorHandler = func(err error, context echo.Context) {
		trans, ok := a.ut.GetTranslator("zh")
		if !ok {
			trans, _ = a.ut.GetTranslator("zh")
		}
		resp := a.errorHandler.Handler(err, &trans)
		err = context.JSON(200, resp)
		if err != nil {
			panic(err)
		}
	}
	// 校验器
	a.echo.Validator = NewEchoValidator(a.validator)
	// 中间件
	a.echo.Use(middlewares.LoggerMiddleware(logger.GetLogger()))
	a.echo.Use(middlewares.SessionMiddleware())
}

func (a *App) initTranslator() {
	zhT := zh.New()
	enT := en.New()
	a.ut = ut.New(zhT, zhT, enT)
	translator, _ := a.ut.GetTranslator("zh")
	// 注册翻译器
	_ = zhTranslations.RegisterDefaultTranslations(a.validator, translator)
}

func (a *App) initErrorHandler() {
	a.errorHandler = result.NewErrorHandler()
}

func (a *App) initValidate() {
	a.validator = validator.New()
}
