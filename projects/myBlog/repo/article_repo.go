package repo

import (
	"github.com/spf13/cast"
	"myBlog/dao"
	"myBlog/models"
)

type ArticleRepository interface {
	GetArticleList(m map[string]interface{}) (total int,articles []models.Article)
	GetArticle(articleID uint) (article models.Article)
	SaveArticle(m map[string]interface{}) (article models.Article,err error)
	SaveArticleCategories(articleID uint,categoryStr string) (article models.Article)
	GetArticleCount() (map[string]int)
}

func NewArticleRepository() ArticleRepository {
	return &articleRepository{}
}

type articleRepository struct {}

func (n articleRepository) GetArticleCount() (map[string]int) {
	lifeCount := 0
	skillCount := 0
	temp := make(map[string]int)
	dao.GetDB().Model(&models.Article{}).Where("personal = ?",1).Count(&lifeCount)
	dao.GetDB().Model(&models.Article{}).Where("personal = ?",0).Count(&lifeCount)
	temp["Life"] = lifeCount
	temp["Skill"] = skillCount
	return temp
}

func (n articleRepository) GetArticle(articleID uint) (article models.Article) {
	db := dao.GetDB()
	db.First(&article,articleID)
	return

}

func (n articleRepository) GetArticleList(m map[string]interface{}) (total int,articles []models.Article) {
	db := dao.GetDB()
	var err error
	if m["Personal"] == nil {
		dao.GetDB().Model(&models.Article{}).Count(&total)
		err = db.Limit(cast.ToInt(m["Size"])).Offset((cast.ToInt(m["Page"]) - 1) * cast.ToInt(m["Size"])).Order("created_at desc").Find(&articles).Error
	} else {
		dao.GetDB().Model(&models.Article{}).Where("personal = ?",cast.ToInt(m["Personal"])).Count(&total)
		err = db.Limit(cast.ToInt(m["Size"])).Offset((cast.ToInt(m["Page"]) - 1) * cast.ToInt(m["Size"])).Where("personal = ?",cast.ToInt(m["Personal"])).Order("created_at desc").Find(&articles).Error
	}
	if err != nil {
		panic("select Error")
	}
	return
}

func (n articleRepository) SaveArticle(m map[string]interface{}) (article models.Article,err error) {
	content := cast.ToString(m["Content"])
	article.Content = content
	article.Title = cast.ToString(m["Title"])
	article.Personal = cast.ToUint(m["Personal"])
	article.Tags = cast.ToString(m["Tags"])
	db := dao.GetDB()
	err = db.Save(&article).Error
	return
}


func (n articleRepository) SaveArticleCategories(articleID uint,categoryStr string) (article models.Article) {
	article.ID = articleID
	db := dao.GetDB()
	db.Model(&article).Update("categories",categoryStr)
	return 
}

