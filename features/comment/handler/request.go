package handler

type AddCommentReq struct {
	UserID  uint `json:"user_id" form:"user_id"`
	Content uint `json:"content" form:"content"`
}
