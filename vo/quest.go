package vo

//CreateQuestRequest c
type CreateQuestRequest struct {
	Title       string `json:"title" binding:"required,max=20"`
	Description string `json:"description" binding:"required,max=50"`
	Content     string `json:"content" binding:"required"`
}
