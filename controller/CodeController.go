package controller

import (
	"log"
	"main/common"
	"main/model"
	"main/response"
	"main/storage"
	"main/utility"
	"main/vo"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//CodeSaveController c
type CodeSaveController interface {
	RestController
	CodeList(c *gin.Context)
}

//CodeController p
type CodeController struct {
	DB *gorm.DB
}

//NewCodeController n
func NewCodeController() CodeSaveController {
	db := common.GetDB()
	db.AutoMigrate(model.CodeSave{})
	return CodeController{DB: db}
}

//Create c
func (csv CodeController) Create(c *gin.Context) {
	var requestCode vo.CreateCodeRequest
	// 数据验证
	if err := c.Bind(&requestCode); err != nil {
		log.Println(err.Error())
		response.Fail(c, nil, "数据验证错误")
		return
	}

	// 获取登录用户 user
	user, _ := c.Get("user")
	// 创建post
	code := model.CodeSave{
		UserID: user.(model.User).ID,
		Name:   user.(model.User).Name,
		Code:   utility.ReplaceO(requestCode.Code),
		Title:  requestCode.Title,
	}
	if err := csv.DB.Create(&code).Error; err != nil {
		panic(err)
	}
	// storage.StorageCode(code.Name, code.Title, code.Code)
	storage.StorageCode(code.Code)

	response.Success(c, nil, "创建成功")
}

//Update u
func (csv CodeController) Update(c *gin.Context) {
	var requestCode vo.CreateCodeRequest
	// 数据验证
	if err := c.ShouldBind(&requestCode); err != nil {
		log.Println(err.Error())
		response.Fail(c, nil, "数据验证错误")
		return
	}

	// 获取path 中的id
	codeID := c.Params.ByName("id")

	var code model.CodeSave
	if csv.DB.Where("id = ?", codeID).First(&code).RecordNotFound() {
		response.Fail(c, nil, "代码不存在")
		return
	}

	// 判断当前用户是否为代码的作者
	// 获取登录用户 user
	user, _ := c.Get("user")
	userID := user.(model.User).ID
	if userID != code.UserID {
		response.Fail(c, nil, "代码不属于您，请勿非法操作")
		return
	}

	// 更新文章
	if err := csv.DB.Model(&code).Update(requestCode).Error; err != nil {
		response.Fail(c, nil, "更新失败")
		return
	}
	// storage.StorageCode(code.Name, code.Title, code.Code)
	storage.StorageCode(code.Code)

	response.Success(c, gin.H{"code": code}, "更新成功")
}

//Show s
func (csv CodeController) Show(c *gin.Context) {
	//获取path中的id
	// categoryID := c.Params.ByName("cid")
	codeID := c.Params.ByName("id")

	var code model.CodeSave
	if csv.DB.Where("id=?", codeID).First(&code).RecordNotFound() {
		response.Fail(c, nil, "代码不存在")
		return
	}
	response.Success(c, gin.H{"code": code}, "访问文章成功")
}

//Delete d
func (csv CodeController) Delete(c *gin.Context) {
	// 获取path 中的id
	codeID := c.Params.ByName("id")

	var code model.CodeSave
	if csv.DB.Where("id = ?", codeID).First(&code).RecordNotFound() {
		response.Fail(c, nil, "代码不存在")
		return
	}

	// 判断当前用户是否为文章的作者
	// 获取登录用户 user
	user, _ := c.Get("user")
	userID := user.(model.User).ID
	if userID != code.UserID {
		response.Fail(c, nil, "代码不属于您，请勿非法操作")
		return
	}

	csv.DB.Delete(&code)
	response.Success(c, gin.H{"code": code}, "删除成功")
}

//CodeList c
func (csv CodeController) CodeList(c *gin.Context) {
	user, _ := c.Get("user")
	//获取path中的id
	var codes []model.CodeSave
	csv.DB.Where("user_id=?", user.(model.User).ID).Find(&codes)

	var total int
	csv.DB.Model(model.CodeSave{}).Count(&total)
	response.Success(c, gin.H{"code": codes, "total_pages": total}, "成功")
}

//IfCodeCreated i
func (csv CodeController) IfCodeCreated(c *gin.Context) {
	user, _ := c.Get("user")

	var codes []model.CodeSave
	csv.DB.Where("user_id=?", user.(model.User).ID).Find(&codes)

	var total int
	csv.DB.Model(model.CodeSave{}).Count(&total)
	response.Success(c, gin.H{"code": codes, "total_pages": total}, "成功")
}
