package repo

import (
	"myBlog/dao"
	"myBlog/models"
)

type ArticleCategoryReposity interface {
	SaveArticleCategory(articleID uint,categoryID uint,categoryName string) (articleCategory models.ArtcileCategory)
	GetArticleCategoryList(articleID uint) (articleCategories []models.ArtcileCategory)
}

func NewArticleCategoryRepository() ArticleCategoryReposity {
	return &articleCategoryRepository{}
}

type articleCategoryRepository struct {}


func (n articleCategoryRepository) SaveArticleCategory(articleID uint,categoryID uint,categoryName string) (articleCategory models.ArtcileCategory) {
	articleCategory.ArticleID = articleID
	articleCategory.CategoryID = categoryID
	articleCategory.CategoryName = categoryName
	db := dao.GetDB()
	db.Save(&articleCategory)
	return
}

func (n articleCategoryRepository) GetArticleCategoryList(articleID uint) (articleCategories []models.ArtcileCategory) {
	db := dao.GetDB()
	db.Where("article_id = ?",articleID).Find(&articleCategories)
	return
}
