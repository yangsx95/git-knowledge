package app

import (
	"git-knowledge/db"
	"git-knowledge/logger"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/mongo"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"log"
	"os"
	"time"
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

	resource, err := db.InitResource(os.Getenv("MONGO_HOST")+":"+os.Getenv("MONGO_PORT"), os.Getenv("MONGO_DATABASE"))
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
	// 日志中间件
	engine.Use(GinLoggerMiddleware(logger.GetLogger()))
	// session处理
	engine.Use(GinSessionMiddleware())

	return engine
}

func GinSessionMiddleware() gin.HandlerFunc {
	session, err := mgo.Dial(os.Getenv("MONGO_HOST") + "+" + os.Getenv("MONGO_PORT"))
	if err != nil {
		log.Fatalln("初始化Session出现异常", err)
	}
	c := session.DB("").C("sessions")
	store := mongo.NewStore(c, 3600, true, []byte("secret"))
	return sessions.Sessions("mongoSession", store)
}

func GinLoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		logger.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}
