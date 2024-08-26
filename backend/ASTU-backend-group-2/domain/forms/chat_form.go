package forms

type CreateChatForm struct {
	UserID string `json:"user_id,omitempty" binding:"required,hexadecimal"`
	Title  string `json:"title,omitempty" binding:"required,min=5,max=30"`
}

type MessageForm struct {
	Message string `json:"message,omitempty" binding:"required,min=5,max=300"`
}
