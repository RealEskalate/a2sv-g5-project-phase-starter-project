package auth

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsActive bool   `json:"isactive"`
	IsAdmin  bool   `json:"isadmin"`
}

type Token struct {
	ID          string
	TokenString string
}

type AuthRepository interface {
	CreateUser(user User) error // return the user id
	UpdateUser(id string, user User) (User, error)
	GetUserByUsername(username string) (User, error)
	GetUserByEmail(email string) (User, error)
	RegisterToken(token string) error
	DeleteUser(id string) error
	GetToken(token string) (Token, error)
	DeleteToken(token string) error
}

type AuthServices interface {
	Login(user User)
	RegisterUser(user User)
	UpdateProfile(user User)
	Activate(userID string, token string)
	Logout(userID string)
	GenerateToken(user User) (string, error)
}

// Register(User) error
// Login(User) error
// Activate(User) error
