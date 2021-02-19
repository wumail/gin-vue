package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// {
// 	code:20001,
// 	data:xxx,
// 	msg:xx
// }

//Response 统一请求返回
func Response(c *gin.Context, httpstatus int, code int, data gin.H, msg string) {
	c.JSON(httpstatus, gin.H{"code": code, "data": data, "msg": msg})
}

//Success 常用的请求成功返回
func Success(c *gin.Context, data gin.H, msg string) {
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": data, "msg": msg})
}

//Fail 常用的请求失败返回
func Fail(c *gin.Context, data gin.H, msg string) {
	c.JSON(http.StatusOK, gin.H{"code": 400, "data": data, "msg": msg})
}
