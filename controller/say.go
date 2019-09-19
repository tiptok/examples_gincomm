package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tiptok/gocomm/common"
	"github.com/tiptok/gocomm/pkg/mygin"
)

type SayController struct {
	mygin.BaseController
}

func (this *SayController) GenJwt(c *gin.Context) {
	var msg *mygin.Message
	defer func() {
		this.Resp(c,msg)
	}()
	jwt,_ := common.GenerateToken("root","123456")
	msg = mygin.NewErrMessage(0,jwt)

	//fmt.Fprintf(out, "[GIN] %v |%s %3d %s| %13v | %15s |%s %-7s %s %s\n%s",
	//	end.Format("2006/01/02 - 15:04:05"),
	//	statusColor, statusCode, resetColor,
	//	latency,
	//	clientIP,
	//	methodColor, method, resetColor,
	//	path,
	//	comment,
	//}
}

func (this *SayController) Hello(c *gin.Context) {
	var msg *mygin.Message
	defer func() {
		this.Resp(c,msg)
	}()
	msg = mygin.NewErrMessage(0,"hello tiptok")
}
