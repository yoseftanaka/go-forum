package dto

type CreatePostRequest struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	UserID   uint   `json:"user_id"`
	Password string `json:"password"`
}
