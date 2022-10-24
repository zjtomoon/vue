package service

import (
	"myBlog/models"
	"myBlog/repo"
)

type GitUserService interface {
	SaveGitUserInfo(gitUser models.GitUser) (result models.Result)
}

type gitUserServices struct {}

func NewGitUserServices() GitUserService {
	return &gitUserServices{}
}

var gitUserRepo = repo.NewGitUserRepository()

func (u gitUserServices) SaveGitUserInfo(gitUser models.GitUser) (result models.Result) {
	gitUserRepo.SaveGitUserInfo(gitUser)
	return
}