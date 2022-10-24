package repo

import (
	"myBlog/dao"
	"myBlog/models"
)

type GitUserRepository interface {
	SaveGitUserInfo(gitUser models.GitUser) ()
}

type gitUserRepository struct {}

func NewGitUserRepository() GitUserRepository {
	return &gitUserRepository{}
}

func (n gitUserRepository) SaveGitUserInfo(gitUser models.GitUser) () {
	db := dao.GetDB()
	db.Save(&gitUser)
}
