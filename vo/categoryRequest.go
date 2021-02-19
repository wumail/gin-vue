package vo

type CreateCategoryRequest struct {
	Name string `Json:"name" binding:"required"`
}
