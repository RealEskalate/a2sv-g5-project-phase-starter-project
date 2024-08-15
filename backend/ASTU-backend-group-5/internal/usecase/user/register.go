package user

import(
	"blogApp/internal/repository"
	"blogApp/pkg/password"
)

type UserUsecase struct{
	repo repository.UserRepo
}

func NewUserUsecase(Repo *repository.UserRepo) *Register{
	return &Users{
		repo : &Repo,
	}
}

func (s *UserUsecase) AddUser(user *domain.User) (*domain.User, error){
	email := user.Email
	username := user.UserName
	user.ID = primitive.NewObjectId()

	e, err := s.repo.GetUserByEmail(email)


	if e, _ := s.repo.GetUserByEmail(email); e != nil{
		return nil, errors.New("user in this email already exist")
	}
	
	if u, _ := s.repo.GetUserByUserName(username); u != nil{
		return nil, errors.New("user in this username already exist")
	}

	hashedPassword, err := hash.EncryptPassword(username)

	if err != nil{
		return nil, err
	}

	if s.repo.NumberOfUsers() == 0{
		user.Role = "owner"
	}else{
		user.Role = "user"
	}

	user.Password = hashedPassword
	 
	return s.repo.CreateUser(user)
}
