package app

import (
	"git-knowledge/db"
	"git-knowledge/logger"
	"git-knowledge/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

type BootStrap struct {
	engine *gin.Engine
	db     *db.Resource
	Dao    *Dao
	Api    *Api
}

func NewBootstrap() *BootStrap {
	b := BootStrap{}
	// 加载配置文件
	loadConfig()
	// 初始化日志
	logger.InitLogger(os.Getenv("LOG_LEVEL"), os.Getenv("LOG_DIR"))
	// 初始化数据库
	b.db = initDb()
	// 初始化web(gin)引擎
	b.initGinEngine()
	// 初始化Dao组件
	b.Dao = initDao(&b)
	// 初始化Api组件
	b.Api = initApi(&b)
	// 初始化gin router
	initRouter(b.engine.RouterGroup, b.Api)
	return &b
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

func (b *BootStrap) Start() {
	err := b.engine.Run(":8080")
	if err != nil {
		logger.Fatal("启动服务出现错误 %s", err)
	}
}

func (b *BootStrap) initGinEngine() {
	engine := gin.New()
	engine.Use(middlewares.GinLoggerMiddleware(logger.GetLogger()))
	engine.Use(middlewares.GinSessionMiddleware())
	b.engine = engine
}
