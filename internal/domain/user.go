package domain

const UserCollection = "users"

type User struct {
	ID       string `json:"id" bson:"_id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	IsActive bool   `json:"is_active" bson:"is_active"`
	IsAdmin  bool   `json:"is_admin" bson:"is_admin"`
}

type UserRepository interface {
	Create(User) (User, error)
	Update(string, User) (User, error)
	Delete(string) error
	Get() ([]User, error)
	GetByID(string) (User, error)
	GetByEmail(string) (User, error)
}

type UserUsecases interface {
	Register(User) (User, error)
	VerifyEmail(token, email string) error
	Login(password, email string) (User, error)
	GetProfile(userID string) (User, error)
	ForgetPassword(email string) (string, error)
	ResetPassword(token, newPassword string) (string, error)
}
