package chat

type CreateChatForm struct {
	UserID string `json:"user_id,omitempty" validate:"required,hexadecimal"`
	Title string `json:"title,omitempty" validate:"required,min=2,max=30"`
}

