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

//IQeustController i
type IQeustController interface {
	RestController
	QuestList(c *gin.Context)
	// MyQuests(c *gin.Context)
}

//QuestController p
type QuestController struct {
	DB *gorm.DB
}

//NewQuestController n
func NewQuestController() IQeustController {
	db := common.GetDB()
	db.AutoMigrate(model.Question{})
	return QuestController{DB: db}
}

//Create c
func (q QuestController) Create(c *gin.Context) {
	var requestQuest vo.CreateQuestRequest
	// 数据验证
	if err := c.Bind(&requestQuest); err != nil {
		log.Println(err.Error())
		response.Fail(c, nil, "数据验证错误")
		return
	}

	// 获取登录用户 user
	user, _ := c.Get("user")

	// 创建post
	quest := model.Question{
		UserID:      user.(model.User).ID,
		Name:        user.(model.User).Name,
		Description: requestQuest.Description,
		Title:       requestQuest.Title,
		Content:     utility.ReplaceO(requestQuest.Content),
	}
	if err := q.DB.Create(&quest).Error; err != nil {
		panic(err)
	}
	response.Success(c, nil, "创建成功")
}

//Update u
func (q QuestController) Update(c *gin.Context) {
	var requestQuest vo.CreatePostRequest
	// 数据验证
	if err := c.ShouldBind(&requestQuest); err != nil {
		log.Println(err.Error())
		response.Fail(c, nil, "数据验证错误")
		return
	}

	// 获取path 中的id
	questID := c.Params.ByName("id")

	var quest model.Question
	if q.DB.Where("id = ?", questID).First(&quest).RecordNotFound() {
		response.Fail(c, nil, "文章不存在")
		return
	}

	// 判断当前用户是否为文章的作者
	// 获取登录用户 user
	user, _ := c.Get("user")
	userID := user.(model.User).ID
	if userID != quest.UserID {
		response.Fail(c, nil, "文章不属于您，请勿非法操作")
		return
	}

	// 更新文章
	if err := q.DB.Model(&quest).Update(requestQuest).Error; err != nil {
		response.Fail(c, nil, "更新失败")
		return
	}

	response.Success(c, gin.H{"quest": quest}, "更新成功")
}

//Show s
func (q QuestController) Show(c *gin.Context) {
	//获取path中的id
	questID := c.Params.ByName("id")

	var quest model.Question
	if q.DB.Where("id=?", questID).First(&quest).RecordNotFound() {
		response.Fail(c, nil, "文章不存在")
		return
	}
	response.Success(c, gin.H{"quest": quest}, "访问文章成功")
}

//Delete d
func (q QuestController) Delete(c *gin.Context) {
	// 获取path 中的id
	questID := c.Params.ByName("id")

	var quest model.Question
	if q.DB.Where("id = ?", questID).First(&quest).RecordNotFound() {
		response.Fail(c, nil, "文章不存在")
		return
	}

	// 判断当前用户是否为文章的作者
	// 获取登录用户 user
	user, _ := c.Get("user")
	userID := user.(model.User).ID
	if userID != quest.UserID {
		response.Fail(c, nil, "文章不属于您，请勿非法操作")
		return
	}

	q.DB.Delete(&quest)

	response.Success(c, gin.H{"quest": quest}, "删除成功")
}

//QuestList p
func (q QuestController) QuestList(c *gin.Context) {
	//获取分页参数
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize := 6

	//分页
	var quests []model.Question
	q.DB.Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&quests)

	//前端渲染需要知道总数
	var total int
	q.DB.Model(model.Question{}).Count(&total)
	response.Success(c, gin.H{"quests": quests, "total_pages": total}, "成功")
}

// //MyQuests m
// func (q QuestController) MyQuests(c *gin.Context) {
// 	//获取分页参数
// 	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
// 	pageSize := 6

// 	//分页
// 	var quests []model.Question
// 	user, _ := c.Get("user")
// 	q.DB.Order("created_at desc").Where("user_id=?", user.(model.User).ID).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&quests)

// 	var total int
// 	q.DB.Model(model.Question{}).Where("user_id=?", user.(model.User).ID).Count(&total)
// 	response.Success(c, gin.H{"quests": quests, "total_pages": total}, "成功")
// }
