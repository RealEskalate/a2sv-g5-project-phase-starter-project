package service

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
)

type userService struct {
	userRepository interfaces.UserRepository
}

func NewUserService(userRepository interfaces.UserRepository) interfaces.UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (service *userService) CreateUser(user *entities.User) (*entities.User, error){
	return service.userRepository.CreateUser(user)
}

func (service *userService) FindUserByEmail(email string) (*entities.User, error){
	return service.userRepository.FindUserByEmail(email)
}

func (service *userService) FindUserById(userId string) (*entities.User, error){
	return service.userRepository.FindUserById(userId)
}

func (service *userService) UpdateUser(user *entities.User) (*entities.User, error){
	return service.userRepository.UpdateUser(user)
}

func (service *userService) DeleteUser(userId string) error{
	return service.userRepository.DeleteUser(userId)
}



func (service *userService) PromoteUserToAdmin(userId string) error{
	user, err := service.userRepository.FindUserById(userId)
	if err != nil {
		return err
	}

	user.Role = "admin"
	_, err = service.userRepository.UpdateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (service *userService) DemoteUserToRegular(userId string) error{
	user, err := service.userRepository.FindUserById(userId)
	if err != nil {
		return err
	}

	user.Role = "regular"
	_, err = service.userRepository.UpdateUser(user)
	if err != nil {
		return err
	}

	return nil
}
