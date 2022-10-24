package repo

import (
	"github.com/spf13/cast"
	"myBlog/dao"
	"myBlog/models"
)

type CommentRepository interface {
	SaveComment(m map[string]interface{}) (comment models.Comment)
	GetCommentList(m map[string]interface{}) (total int,comments []models.Comment)
	GetComment(commentID uint) (comment models.Comment)
}

type commentRepository struct {}

func NewCommentRepository() CommentRepository {
	return &commentRepository{}
}

var articleRepo = NewArticleRepository()

func (n commentRepository) SaveComment(m map[string]interface{}) (comment models.Comment) {
	db := dao.GetDB()
	article := articleRepo.GetArticle(cast.ToUint(m["ArticleID"]))
	comment.ArticleTitle = article.Title
	comment.Content = cast.ToString(m["Content"])
	comment.ArticleID = cast.ToUint(m["ArticleID"])
	comment.GitUserID = cast.ToUint(m["GitUserID"])
	comment.Username = cast.ToString(m["Username"])
	comment.AvatarUrl = cast.ToString(m["AvatarUrl"])
	comment.GithubUrl = cast.ToString(m["GithubUrl"])
	db.Save(&comment)
	return
}

func (n commentRepository) GetCommentList(m map[string]interface{}) (total int,comments []models.Comment) {
	db := dao.GetDB()
	var err error
	if m["ArticleID"] == nil {
		if m["Size"] == nil || m["Page"] == nil {
			err = db.Order("created_at desc").Find(&comments).Error
		}else {
			db.Model(&models.Comment{}).Count(&total)
			err = db.Limit(cast.ToInt(m["Size"])).Offset((cast.ToInt(m["Page"])-1) * cast.ToInt(m["Size"])).Order("created_at desc").Find(&comments).Error
		}
	} else {
		err = db.Where("article_id = ?",cast.ToUint(m["ArticleID"])).Order("created_at desc").Find(&comments).Error
	}
	if err != nil {
		panic("select Error")
	}
	return
}

func (n commentRepository) GetComment(commentID uint) (comment models.Comment) {
	db := dao.GetDB()
	db.First(&comment,commentID)
	return
}
