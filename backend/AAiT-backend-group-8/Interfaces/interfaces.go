package Infrastructure

type IUserRepository interface {
	RegisterNewUser(domain.User) error
	GetSingleUser(email string) *domain.User, error
	PromoteUser(userID primitive.ObjectID) error
	// ForgetPassword()
	// DeleteRefreshToken(primitive.ObjectID) error
	DemoteUser(userID primitive.ObjectID) error
	DeleteUser(userID primitive.ObjectID) error
	UpdateUser(user domain.User) error	
}

type IUserUseCase interface {
	RegisterNewUser(user domain.User) error
	Login(email, password string ) string error
	GetSingleUser(email string) *domain.User
	PromoteUser(userID string ) error
	DemoteUser (userID string) error
	UpdateUser (userID string ,user *domain.User) error
}




	