package bootstrap

import (
	"git-knowledge/conf"
	"git-knowledge/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type BootStrap struct {
	engine *gin.Engine
}

func NewBootstrap() *BootStrap {
	conf.InitConfig("./git-knowledge.ini")
	logger.InitLog()
	engine := InitGinEngine()

	b := BootStrap{
		engine: engine,
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
	engine.Use(GinLogger(logger.GetLogger()))
	return engine
}

func GinLogger(logger *zap.Logger) gin.HandlerFunc {
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
