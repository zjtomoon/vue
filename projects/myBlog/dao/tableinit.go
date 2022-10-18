package dao

import "myBlog/models"

func Createtable()  {
	GetDB().AutoMigrate(
		&models.User{},
		&models.Article{},
		&models.Category{},
		&models.ArtcileCategory{},
		&models.GitUser{},
		&models.Comment{},
		)
}
