package repo

import (
	"myBlog/dao"
	"myBlog/models"

	"github.com/spf13/cast"
)

type CategoryRepository interface {
	SaveCategory(m map[string]interface{}) (category models.Category)
	GetCategoryList(m map[string]interface{})(categories []models.Category)
	GetCategory(CategoryID uint)(category models.Category)
}

type categoryRepository struct{}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{}
}

func (n categoryRepository) SaveCategory(m map[string]interface{}) (category models.Category) {
	category.Name = cast.ToString(m["Name"])
	db := dao.GetDB()
	db.Save(&category)
	return
}

func (n categoryRepository) GetCategory(categoryID uint) (category models.Category) {
	db := dao.GetDB()
	db.First(&category,categoryID)
	return
}

func (n categoryRepository) GetCategoryList(m map[string]interface{}) (categories []models.Category) {
	db := dao.GetDB()
	err := db.Find(&categories).Error
	if err != nil {
		panic("select Error")
	}
	return
}