package controller

import (
	"log"
	"main/common"
	"main/model"
	"main/response"
	"main/utility"
	"main/vo"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//IPostController i
type IPostController interface {
	RestController
	PageList(c *gin.Context)
	MyPosts(c *gin.Context)
}

//PostController p
type PostController struct {
	DB *gorm.DB
}

//NewPostController n
func NewPostController() IPostController {
	db := common.GetDB()
	db.AutoMigrate(model.Mypost{})
	return PostController{DB: db}
}

//Create c
func (p PostController) Create(c *gin.Context) {
	var requestPost vo.CreatePostRequest
	// 数据验证
	if err := c.Bind(&requestPost); err != nil {
		log.Println(err.Error())
		response.Fail(c, nil, "数据验证错误")
		return
	}

	// 获取登录用户 user
	user, _ := c.Get("user")

	// 创建post
	post := model.Mypost{
		UserID:      user.(model.User).ID,
		Name:        user.(model.User).Name,
		Description: requestPost.Description,
		Title:       requestPost.Title,
		Content:     utility.ReplaceO(requestPost.Content),
	}
	if err := p.DB.Create(&post).Error; err != nil {
		panic(err)
	}
	// storage.Storage(post.Name, post.Content)
	response.Success(c, nil, "创建成功")
}

//Update u
func (p PostController) Update(c *gin.Context) {
	var requestPost vo.CreatePostRequest
	// 数据验证
	if err := c.ShouldBind(&requestPost); err != nil {
		log.Println(err.Error())
		response.Fail(c, nil, "数据验证错误")
		return
	}

	// 获取path 中的id
	postID := c.Params.ByName("id")

	var post model.Mypost
	if p.DB.Where("id = ?", postID).First(&post).RecordNotFound() {
		response.Fail(c, nil, "文章不存在")
		return
	}

	// 判断当前用户是否为文章的作者
	// 获取登录用户 user
	user, _ := c.Get("user")
	userID := user.(model.User).ID
	if userID != post.UserID {
		response.Fail(c, nil, "文章不属于您，请勿非法操作")
		return
	}

	// 更新文章
	if err := p.DB.Model(&post).Update(requestPost).Error; err != nil {
		response.Fail(c, nil, "更新失败")
		return
	}
	// storage.Storage(post.Name, post.Content)
	response.Success(c, gin.H{"post": post}, "更新成功")
}

//Show s
func (p PostController) Show(c *gin.Context) {
	//获取path中的id
	postID := c.Params.ByName("id")

	var post model.Mypost
	if p.DB.Where("id=?", postID).First(&post).RecordNotFound() {
		response.Fail(c, nil, "文章不存在")
		return
	}
	response.Success(c, gin.H{"post": post}, "访问文章成功")
}

//Delete d
func (p PostController) Delete(c *gin.Context) {
	// 获取path 中的id
	postID := c.Params.ByName("id")

	var post model.Mypost
	if p.DB.Where("id = ?", postID).First(&post).RecordNotFound() {
		response.Fail(c, nil, "文章不存在")
		return
	}

	// 判断当前用户是否为文章的作者
	// 获取登录用户 user
	user, _ := c.Get("user")
	userID := user.(model.User).ID
	if userID != post.UserID {
		response.Fail(c, nil, "文章不属于您，请勿非法操作")
		return
	}

	p.DB.Delete(&post)

	response.Success(c, gin.H{"post": post}, "删除成功")
}

//PageList p
func (p PostController) PageList(c *gin.Context) {
	//获取分页参数
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize := 6

	//分页
	var posts []model.Mypost
	p.DB.Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&posts)

	//前端渲染需要知道总数
	var total int
	p.DB.Model(model.Mypost{}).Count(&total)
	response.Success(c, gin.H{"post": posts, "total_pages": total}, "成功")
}

//MyPosts m
func (p PostController) MyPosts(c *gin.Context) {
	//获取分页参数
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize := 6

	//分页
	var posts []model.Mypost
	user, _ := c.Get("user")
	p.DB.Order("created_at desc").Where("user_id=?", user.(model.User).ID).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&posts)

	var total int
	p.DB.Model(model.Mypost{}).Where("user_id=?", user.(model.User).ID).Count(&total)
	response.Success(c, gin.H{"post": posts, "total_pages": total}, "成功")
}
