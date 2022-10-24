package service

import (
	"github.com/spf13/cast"
	"myBlog/models"
	"myBlog/repo"
	"strings"
)

type ArticleService interface {
	GetArticleList(m map[string]interface{}) (result models.Result)
	SaveArticle(m map[string]interface{}) (result models.Result)
	GetArticle(m map[string]interface{}) (result models.Result)
	GetArticleCount() (result models.Result)
}

type articleService struct {}

func NewArticleService() ArticleService {
	return &articleService{}
}

var articleRepo = repo.NewArticleRepository()
var articleCategoryRepo = repo.NewArticleCategoryRepository()

func (u articleService) GetArticleList(m map[string]interface{}) (result models.Result) {
	result.Code = 0
	total,articles := articleRepo.GetArticleList(m)
	maps := make(map[string]interface{},2)
	maps["Total"] = total
	maps["List"] = articles
	result.Data = maps
	return
}

func (u articleService) SaveArticle(m map[string]interface{}) (result models.Result) {
	article,err := articleRepo.SaveArticle(m)
	if err != nil {
		result.Code = -1
		result.Msg = cast.ToString(err)
		return
	}
	articleID := article.ID
	ids := strings.Split(cast.ToString(m["CategoryIDs"]),",")
	for _,i := range ids {
		category := categoryRepo.GetCategory(cast.ToUint(i))
		articleCategoryRepo.SaveArticleCategory(articleID,cast.ToUint(i),category.Name)
	}
	categories := articleCategoryRepo.GetArticleCategoryList(articleID)
	categoriesArr := []string{}
	for _,item := range categories{
		categoriesArr = append(categoriesArr,item.CategoryName)
	}
	article2Save := articleRepo.SaveArticleCategories(articleID,strings.Join(categoriesArr,","))
	article.Categories = article2Save.Categories
	maps := make(map[string]interface{},1)
	maps["article"] = article
	result.Code = 0
	result.Data = maps
	result.Msg = "保存成功"
	return
}

func (u articleService) GetArticle(m map[string]interface{}) (result models.Result) {
	result.Code = 0
	article := articleRepo.GetArticle(cast.ToUint(m["ID"]))
	result.Data = article
	return
}


func (u articleService) GetArticleCount() (result models.Result) {
	result.Code = 0
	count := articleRepo.GetArticleCount()
	result.Data = count
	return
}
