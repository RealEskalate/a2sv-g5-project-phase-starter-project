package usecase

import (
	infrastructure "astu-backend-g1/Infrastructure"
	"astu-backend-g1/domain"
	"crypto/rand"
	"math/big"
)

type userUsecase struct {
	userRepository domain.UserRepository
}

func NewUserUsecase(u domain.UserRepository) (domain.UserUsecase, error) {
	return &userUsecase{userRepository: u}, nil
}

func (useCase *userUsecase) Get() ([]domain.User, error) {
	return useCase.userRepository.Get(domain.UserFilterOption{})
}

func (useCase *userUsecase) GetByID(userID string) (domain.User, error) {
	filter := domain.UserFilter{UserId: userID}
	opts := domain.UserFilterOption{Filter: filter}
	users, err := useCase.userRepository.Get(opts)
	return users[0], err
}

func (useCase *userUsecase) GetByUsername(username string) (domain.User, error) {
	filter := domain.UserFilter{Username: username}
	opts := domain.UserFilterOption{Filter: filter}
	users, err := useCase.userRepository.Get(opts)
	return users[0], err
}
func (useCase *userUsecase) AccountVerification(uemail string, confirmationToken string) (string, error) {
	filter := domain.UserFilter{Email: uemail}
	opts := domain.UserFilterOption{Filter: filter}
	users, err := useCase.userRepository.Get(opts)
	if users[0].VerifyToken == confirmationToken {
		users[0].IsActive = true
		useCase.userRepository.Update(users[0].ID, users[0])
		token, err := infrastructure.Genratetoken(&users[0], users[0].Password)
		if err != nil {
			return "", err
		}
		return token, nil

	}
	return "", err
}

func (useCase *userUsecase) GetByEmail(email string) (domain.User, error) {
	filter := domain.UserFilter{Email: email}
	opts := domain.UserFilterOption{Filter: filter}
	users, err := useCase.userRepository.Get(opts)
	return users[0], err
}

func (useCase *userUsecase) Create(u *domain.User) (domain.User, error) {
	u.Password, _ = infrastructure.PasswordHasher(u.Password)
	u.IsActive = false
	nUser, err := useCase.userRepository.Create(u)
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	confirmationToken := make([]byte, 64)
	charsetLength := big.NewInt(int64(len(charset)))

	for i := 0; i < 64; i++ {
		num, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return nUser, err
		}
		confirmationToken[i] = charset[num.Int64()]
	}
	err = infrastructure.SendEmail(u.Email, "Registration Confirmation", "This sign up Confirmation email to verify: ", string(confirmationToken))
	if err != nil {
		return nUser, err
	}
	return nUser, err
}

func (useCase *userUsecase) Update(userId string, updateData domain.User) (domain.User, error) {
	return useCase.userRepository.Update(userId, updateData)
}

func (useCase *userUsecase) Delete(userId string) error {
	return useCase.userRepository.Delete(userId)
}
