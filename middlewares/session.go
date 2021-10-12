package middlewares

import (
	"github.com/globalsign/mgo"
	"github.com/kidstuff/mongostore"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"log"
	"os"
)

func SessionMiddleware() echo.MiddlewareFunc {
	s, err := mgo.Dial(os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT"))
	if err != nil {
		log.Fatalln("初始化Session中间件出现异常", err)
	}
	c := s.DB("").C("sessions")
	store := mongostore.NewMongoStore(c, 3600, true, []byte("secret"))
	return session.Middleware(store)
}
