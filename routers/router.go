package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/tiptok/examples_gincomm/controller"
	"github.com/tiptok/gocomm/pkg/mygin"
)

var ginEngin *gin.Engine

func InitRouter(addr string){
	//r := gin.Default() 使用default 默认使用两个中间件  log recovery

	ginEngin=gin.New()
	ginEngin.Use(gin.Logger())
	ginEngin.Use(gin.Recovery())

	//注册中间间
	base :=&mygin.BaseController{}
	ginEngin.Use(base.Prepare)

	LoadUserRouter()
	LoadSayRouter()
	ginEngin.Run(addr)//":8081"
}

func LoadSayRouter(){
	say:=ginEngin.Group("/say")
	{
		sayController := controller.SayController{}
		say.GET("/hello",sayController.Hello)
		say.GET("/genjwt",sayController.GenJwt)
	}
}

func LoadUserRouter(){
	v1:=ginEngin.Group("/v1")
	userGroup := v1.Group("user")
	{
		userController :=&controller.UserController{}
		//验证 jwt
		userGroup.Use(userController.JWTMiddleware())
		userGroup.POST("/select",userController.SelectUser)
	}
}
