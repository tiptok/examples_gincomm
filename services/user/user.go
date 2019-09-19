package user

import (
	"github.com/tiptok/examples_gincomm/models"
	"github.com/tiptok/gocomm/pkg/log"
	"github.com/tiptok/gocomm/pkg/mygin"
)

func GetUserList(mobile string)*mygin.Message{
	 u,err:=models.GetUserByMobile(mobile)
	 log.Debug("GetUserList mobile:",mobile)
	 if err!=nil{
	 	log.Error(err)
	 	log.Error(mobile)
	 	return mygin.NewErrMessage(1)
	 }
	 msg :=mygin.NewMessage(0)
	 msg.Data = u
	 return msg
}
