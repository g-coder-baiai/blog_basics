package service

import (
	"errors"
	"log"
	"ms-go-blog/dao"
	"ms-go-blog/models"
	"ms-go-blog/utils"
)

func Login(userName,passwd string)(*models.LoginRes,error){

	passwd = utils.Md5Crypt(passwd,"baiai")

	user := dao.GetUser(userName,passwd)
	if user == nil{
		return nil,errors.New("账号密码不正确")
	}

	uid := user.Uid

	// 生成token
	token,err:=utils.Award(&uid)
	if err!=nil{
		log.Println(err)
		return nil,errors.New("token生成错误")
	}

	var userInfo models.UserInfo
	userInfo.Uid=user.Uid
	userInfo.UserName=user.UserName
	userInfo.Avatar=user.Avatar

	var lr =&models.LoginRes{
		token,
		userInfo,
	}




	return lr,nil
}