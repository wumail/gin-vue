package controller

import (
	"main/model"
	"main/repository"
	"main/response"
	"main/vo"
	"strconv"

	"github.com/gin-gonic/gin"
)

//ICategoryController i
type ICategoryController interface {
	RestController
}

//CategoryController c
type CategoryController struct {
	Repository repository.CategoryRepository
}

//NewCategoryController n
func NewCategoryController() ICategoryController {
	repository := repository.NewCategoryRepository()
	repository.DB.AutoMigrate(model.Category{})

	return CategoryController{Repository: repository}
}

//Create c
func (cg CategoryController) Create(c *gin.Context) {
	var requestCategory vo.CreateCategoryRequest

	if err := c.ShouldBind(&requestCategory); err != nil {
		response.Fail(c, nil, "数据验证错误，分类名必填")
		return
	}
	category, err := cg.Repository.Create(requestCategory.Name)
	if err != nil {
		response.Fail(c, nil, "创建失败")
		panic(err)
	}
	response.Success(c, gin.H{"category": category}, "创建成功")
}

//Update u
func (cg CategoryController) Update(c *gin.Context) {
	//绑定body中的参数
	var requestCategory vo.CreateCategoryRequest

	if err := c.ShouldBind(&requestCategory); err != nil {
		response.Fail(c, nil, "数据验证错误，分类名必填")
		return
	}
	//获取path中的参数
	categoryID, _ := strconv.Atoi(c.Params.ByName("id"))

	updateCategory, err := cg.Repository.SelectByID(categoryID)
	if err != nil {
		response.Fail(c, nil, "分类不存在")
		return
	}
	//更新分类
	//map
	//struct
	//name value
	category, err := cg.Repository.Update(*updateCategory, requestCategory.Name)
	if err != nil {
		response.Fail(c, nil, "修改失败")
		return
	}

	response.Success(c, gin.H{"category": category}, "修改成功")
}

//Show s
func (cg CategoryController) Show(c *gin.Context) {
	//获取path中的参数
	categoryID, _ := strconv.Atoi(c.Params.ByName("id"))

	category, err := cg.Repository.SelectByID(categoryID)
	if err != nil {
		response.Fail(c, nil, "分类不存在")
		return
	}
	response.Success(c, gin.H{"category": category}, "")
}

//Delete d
func (cg CategoryController) Delete(c *gin.Context) {
	//获取path中的参数
	categoryID, _ := strconv.Atoi(c.Params.ByName("id"))

	if err := cg.Repository.DeleteByID(categoryID); err != nil {
		response.Fail(c, nil, "删除失败，请重试")
		return
	}
	response.Success(c, nil, "删除成功")
}
