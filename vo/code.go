package vo

type CreateCodeRequest struct {
	Title string `json:"title" binding:"required"`
	Code  string `json:"code" binding:"required"`
}
