package controller

import (
	"git-knowledge/dao"
	"git-knowledge/db"
	"github.com/gin-gonic/gin"
)

func ApplyLoginRouter(rg *gin.RouterGroup, resource *db.Resource) {
	userDao := dao.GetUserDaoInstance(resource)
	rg.POST("login", login(userDao))
}

func login(userDao dao.UserDao) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

	}
}
