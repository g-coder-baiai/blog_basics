package views

import (
	"errors"
	"log"
	"ms-go-blog/common"
	"ms-go-blog/service"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Index(w http.ResponseWriter,r *http.Request)  {
	index := common.Template.Index
	//页面上涉及到的所有的数据，必须有定义
	//数据库查询
	if err:=r.ParseForm();err!=nil{
		log.Println("表单获取失败")
		index.WriteError(w,errors.New("系统错误"))
	}

	pageStr:=r.Form.Get("page")
	pageSize := 10
	page:=1
	if pageStr!=""{
		page,_=strconv.Atoi(pageStr)
	}
	path := r.URL.Path
	slug := strings.TrimPrefix(path,"/")
	hr,err := service.GetAllIndexInfo(slug,page,pageSize)
	if err != nil {
		log.Println("Index获取数据出错：",err)
		index.WriteError(w,errors.New("系统错误，请联系管理员!!"))
	}
	index.WriteData(w,hr)
}




