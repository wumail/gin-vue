package controller

import "github.com/gin-gonic/gin"

//RestController s
type RestController interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Show(c *gin.Context)
	Delete(c *gin.Context)
}
