package service

import (
	"ms-go-blog/dao"
	"ms-go-blog/models"
)

func SearchPost(condition string) []models.SearchResp  {
	posts,_ := dao.GetPostSearch(condition)
	var searchResps []models.SearchResp
	for _,post := range posts{
		searchResps = append(searchResps,models.SearchResp{
			post.Pid,
			post.Title,
		})
	}
	return searchResps
}