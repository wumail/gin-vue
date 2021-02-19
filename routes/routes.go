package routes

import (
	"main/controller"
	"main/middleware"

	"github.com/gin-gonic/gin"
)

//CollectRouter 路由
func CollectRouter(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORMiddleware(), middleware.RecoveryMiddleware())
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.GET("/info", middleware.AuthMiddleware(), controller.Info)

	categoryRoutes := r.Group("/categories")
	categoryController := controller.NewCategoryController()
	categoryRoutes.POST("", categoryController.Create)
	categoryRoutes.PUT("/:id", categoryController.Update)
	categoryRoutes.GET("/:id", categoryController.Show)
	categoryRoutes.DELETE("/:id", categoryController.Delete)
	// categoryRoutes.PATCH()

	postRoutes := r.Group("/posts")
	postController := controller.NewPostController()
	postRoutes.POST("", middleware.AuthMiddleware(), postController.Create)
	postRoutes.PUT("/:id", middleware.AuthMiddleware(), postController.Update)
	postRoutes.GET("/:id", postController.Show)
	postRoutes.DELETE("/:id", middleware.AuthMiddleware(), postController.Delete)
	postRoutes.POST("/page/list", postController.PageList)
	postRoutes.POST("/mypost/list", middleware.AuthMiddleware(), postController.MyPosts)

	questRoutes := r.Group("/quests")
	questController := controller.NewQuestController()
	questRoutes.POST("", middleware.AuthMiddleware(), questController.Create)
	questRoutes.PUT("/:id", middleware.AuthMiddleware(), questController.Update)
	questRoutes.GET("/:id", questController.Show)
	questRoutes.DELETE("/:id", middleware.AuthMiddleware(), questController.Delete)
	questRoutes.POST("/list", questController.QuestList)
	// questRoutes.POST("/mypost/list", middleware.AuthMiddleware(), questController.MyQuests)

	codeRoutes := r.Group("/codes")
	codeRoutes.Use(middleware.AuthMiddleware())
	codeController := controller.NewCodeController()
	codeRoutes.POST("", codeController.Create)
	codeRoutes.PUT("/:id", codeController.Update)
	codeRoutes.GET("/:id", codeController.Show)
	codeRoutes.DELETE("/:id", codeController.Delete)
	codeRoutes.POST("/list", codeController.CodeList)

	return r
}
