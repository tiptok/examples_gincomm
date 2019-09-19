package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tiptok/gocomm/pkg/mygin"
	"github.com/tiptok/gocomm/pkg/log"

	"github.com/tiptok/examples_gincomm/services/user"
)

type UserController struct {
	mygin.BaseController
}

func(this *UserController)SelectUser(c *gin.Context){
	var msg *mygin.Message
	defer func(){
		this.Resp(c,msg)
	}()
	type Request struct {
		Mobile string  `json:"mobile"`
	}
	var req Request
	if err:=c.ShouldBindJSON(&req);err!=nil{
		msg = mygin.NewMessage(1)
		log.Error(err)
		return
	}
	if  len(req.Mobile)==0{
		msg = mygin.NewErrMessage(2,"mobile not empty.")
		return
	}
	msg = user.GetUserList(req.Mobile)
	return
}
