package usecase

import "astu-backend-g1/domain"

type userUsecase struct {
	userRepository domain.UserRepository
}

func NewUserUsecase() (domain.UserUsecase, error) {
	return &userUsecase{}, nil
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

func (useCase *userUsecase) GetByEmail(email string) (domain.User, error) {
	filter := domain.UserFilter{Email: email}
	opts := domain.UserFilterOption{Filter: filter}
	users, err := useCase.userRepository.Get(opts)
	return users[0], err
}

func (useCase *userUsecase) Create(u *domain.User) (domain.User, error) {
	return useCase.userRepository.Create(u)
}

func (useCase *userUsecase) Update(userId string, updateData domain.User) (domain.User, error) {
	return useCase.userRepository.Update(userId, updateData)
}

func (useCase *userUsecase) Delete(userId string) error {
	return useCase.userRepository.Delete(userId)
}
