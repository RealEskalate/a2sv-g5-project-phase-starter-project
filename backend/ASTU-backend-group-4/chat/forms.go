package chat

type CreateChatForm struct {
	UserID string `json:"user_id,omitempty" validate:"required,hexadecimal"`
	Title  string `json:"title,omitempty" validate:"required,min=2,max=30"`
}

type TextForm struct {
	Text string `json:"text,omitempty" validate:"required,min=2,max=30"`
}

type DefaultChatForm struct {
	ChatID string `json:"chat_id,omitempty" validate:"required,hexadecimal"`
	UserID string `json:"user_id,omitempty" validate:"required,hexadecimal"`
}

type UserIDForm struct {
	UserID string `json:"user_id,omitempty" validate:"required,hexadecimal"`
}
