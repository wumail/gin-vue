package repository

import (
	"main/common"
	"main/model"

	"github.com/jinzhu/gorm"
)

//CategoryRepository c
type CategoryRepository struct {
	DB *gorm.DB
}

//NewCategoryRepository n
func NewCategoryRepository() CategoryRepository {
	return CategoryRepository{DB: common.GetDB()}
}

//Create c
func (cr CategoryRepository) Create(name string) (*model.Category, error) {
	category := model.Category{
		Name: name,
	}
	if err := cr.DB.Create(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

//Update u
func (cr CategoryRepository) Update(category model.Category, name string) (*model.Category, error) {
	if err := cr.DB.Model(&category).Update("name", name).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

//SelectByID s
func (cr CategoryRepository) SelectByID(id int) (*model.Category, error) {
	var category model.Category
	if err := cr.DB.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

//DeleteByID d
func (cr CategoryRepository) DeleteByID(id int) error {
	if err := cr.DB.Delete(model.Category{}, id).Error; err != nil {
		return err
	}
	return nil
}
