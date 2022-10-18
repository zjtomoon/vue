package service

import "myBlog/models"

type ArticleService interface {
	GetArticleList(m map[string]interface{}) (result models.Result)
	SaveArtice(m map[string]interface{}) (result models.Result)
	GetArticle(m map[string]interface{}) (result models.Result)
	GetArticleCount() (result models.Result)
}

type articleService struct {}

func NewArticleService() ArticleService {
	return &articleService{}
}


