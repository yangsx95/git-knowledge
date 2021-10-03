package app

import (
	"git-knowledge/controller"
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
}

func NewBootstrap() *BootStrap {
	err := godotenv.Load(".env")
	if err != nil {
		panic("加载配置文件.env出现错误")
	}

	logger.InitLogger(os.Getenv("LOG_LEVEL"), os.Getenv("LOG_DIR"))

	resource, err := db.InitResource(
		os.Getenv("MONGO_HOST"),
		os.Getenv("MONGO_PORT"),
		os.Getenv("MONGO_DATABASE"),
		os.Getenv("MONGO_USERNAME"),
		os.Getenv("MONGO_PASSWORD"),
	)
	if err != nil {
		logger.Fatal("连接mongodb出现错误", err)
	}

	engine := InitGinEngine()

	b := BootStrap{
		engine: engine,
		db:     resource,
	}
	return &b
}

func (b *BootStrap) Start() {
	err := b.engine.Run(":8080")
	if err != nil {
		logger.Fatal("启动服务出现错误 %s", err)
	}
}

func InitGinEngine() *gin.Engine {
	engine := gin.New()
	engine.Use(middlewares.GinLoggerMiddleware(logger.GetLogger()))
	engine.Use(middlewares.GinSessionMiddleware())

	controller.ApplyLoginRouter(&engine.RouterGroup, nil)

	return engine
}
