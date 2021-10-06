package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/mongo"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"log"
	"os"
)

func GinSessionMiddleware() gin.HandlerFunc {
	session, err := mgo.Dial(os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT"))
	if err != nil {
		log.Fatalln("初始化Session出现异常", err)
	}
	c := session.DB("").C("sessions")
	store := mongo.NewStore(c, 3600, true, []byte("secret"))
	return sessions.Sessions("mongoSession", store)
}
