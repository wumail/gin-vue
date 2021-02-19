package middleware

import (
	"main/common"
	"main/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//AuthMiddleware 中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//从header中获得Authrization信息
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足:您应该先登录或者注册"}) //权限不足:缺少Bearer header,您应该先登录或者注册
			c.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足:您应该重新登录"}) //权限不足:token无效，您应该重新登录
			c.Abort()
			return
		}

		//验证通过后获取claims中的userID
		userID := claims.UserID
		db := common.GetDB()
		var user model.User
		db.First(&user, userID)

		//用户
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足:用户不存在"})
			c.Abort()
			return
		}

		//用户存在将user信息写入上下文
		c.Set("user", user)
		c.Next()

	}
}
