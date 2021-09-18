package bootstrap

import "github.com/gin-gonic/gin"

type BootStrap struct {
	engine *gin.Engine
}

func New() *BootStrap {
	engine := gin.New()
	engine.Use()

	b := BootStrap{
		engine: engine,
	}
	return &b
}

func initLogger() {

}
