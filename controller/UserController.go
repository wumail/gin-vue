package controller

import (
	"log"
	"main/common"
	"main/dto"
	"main/model"
	"main/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//Register 注册
func Register(c *gin.Context) {
	db := common.GetDB()
	//1 使用map 获取请求参数
	// var requestMap = make(map[string]string)
	// json.NewDecoder(c.Request.Body).Decode(&requestMap)
	//
	//2 使用结构体获取请求参数
	// var requestUser = model.User{}
	// json.NewDecoder(c.Request.Body).Decode(&requestUser)
	//
	//3 gin框架提供的Bind函数绑定请求参数
	var requestUser = model.User{}
	c.Bind(&requestUser)
	//获取参数
	name := requestUser.Name
	// telephone := requestUser.Telephone
	password := requestUser.Password
	group := "user"
	//校验参数
	//如果没有用户名则随机一个用户名
	if len(name) < 8 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户名必须大于8位")
		// log.Println(name, telephone)
		return
	}

	// //验证手机号位数
	// if len(telephone) != 11 {
	// 	response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
	// 	// log.Println(name, telephone)
	// 	return
	// }

	//验证密码位数
	if len(password) < 8 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码必须大于8位")
		log.Println(name, password)
		return
	}

	// log.Println(name, telephone, password)
	// //验证手机号是否存在
	// if isTelephoneExist(db, telephone) {
	// 	c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已存在"})
	// 	return
	// }
	// 验证用户名是否存在
	if isNameExist(db, name) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已存在"})
		return
	}

	//创建用户
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}
	newUser := model.User{
		Name: name,
		// Telephone: telephone,
		Password: string(hashedpassword),
		Group:    group,
	}
	db.Create(&newUser)

	//发放token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error:%v", err)
		return
	}

	//返回结果
	// c.JSON(http.StatusOK, gin.H{
	// 	"code": 200,
	// 	"data": gin.H{"token": token},
	// 	"msg":  "登录成功",
	// })
	response.Success(c, gin.H{"token": token}, "注册成功")

}

//Login 登录
func Login(c *gin.Context) {
	db := common.GetDB()
	//获取数据
	var requestUser = model.User{}
	c.Bind(&requestUser)
	//获取参数
	// telephone := requestUser.Telephone
	name := requestUser.Name
	password := requestUser.Password
	//校验参数
	//验证手机号位数
	// if len(telephone) != 11 {
	// 	response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
	// 	// log.Println(name, telephone)
	// 	return
	// }
	// 验证用户名位数
	if len(password) < 8 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码必须大于8位")
		log.Println(name, password)
		return
	}

	//验证密码位数
	if len(password) < 8 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码必须大于8位")
		// log.Println(name, password)
		return
	}

	//用户是否存在
	var user model.User
	db.Where("name=?", name).First(&user)
	if user.ID == 0 {
		response.Response(c, http.StatusInternalServerError, 500, nil, "用户不存在")
		return
	}
	//验证密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(c, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}

	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error:%v", err)
		return
	}

	//返回结果
	// c.JSON(http.StatusOK, gin.H{
	// 	"code": 200,
	// 	"data": gin.H{"token": token},
	// 	"msg":  "登录成功",
	// })
	response.Success(c, gin.H{"token": token}, "登录成功")

}

//Info 查看用户信息
func Info(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{"code": 200, "user": dto.ToUserDto(user.(model.User))})
}

// //验证手机号是否存在
// func isTelephoneExist(db *gorm.DB, telephone string) bool {
// 	var user model.User
// 	db.Where("telephone=?", telephone).First(&user)
// 	if user.ID != 0 {
// 		return true
// 	}
// 	return false
// }
// 验证用户名是否存在
func isNameExist(db *gorm.DB, name string) bool {
	var user model.User
	db.Where("name=?", name).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
