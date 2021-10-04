package controller

import (
	"git-knowledge/db"
	"github.com/gin-gonic/gin"
)

func ApplyLoginRouter(rg *gin.RouterGroup, resource *db.Resource) {
	rg.POST("registry", JSONHandler(registry))
}

func registry(ctx *gin.Context) (interface{}, error) {

	return nil, nil
}
