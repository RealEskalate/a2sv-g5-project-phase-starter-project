package usecase

import "github.com/RealEskalate/blogpost/domain"

type UserUseCase struct {
	UserRepo domain.User_Registery_interface
}

func (usecase *UserUseCase) GetOneUser(id string) (domain.ResponseUser, error) {
	user,err := usecase.UserRepo.GetUserDocumentByID(id)
	if err != nil {
		return domain.ResponseUser{},err
	}
	response_user := domain.CreateResponseUser(user)
	return response_user,nil
}

func (usecase *UserUseCase) GetUsers() ([]domain.ResponseUser, error) {
	users,err := usecase.UserRepo.GetUserDocuments()
	if err != nil {
		return []domain.ResponseUser{},err
	}
	responses_users:= []domain.ResponseUser{}

	for _,user := range users {
		responses_users = append(responses_users , domain.CreateResponseUser(user))
	}
	return responses_users,nil
}

func (usecase *UserUseCase) UpdateUser(id string, user domain.UpdateUser) (domain.ResponseUser, error) {
	new_user,err := usecase.UserRepo.UpdateUserDocument(id, user)
	if err != nil {
		return domain.ResponseUser{},err
	}
	return domain.CreateResponseUser(new_user),nil
}


func (usecase *UserUseCase) DeleteUser(id string) error {
	return usecase.UserRepo.DeleteUserDocument(id)
}

func (usecase *UserUseCase) LogIn(user domain.LogINUser) (domain.ResponseUser, error) {
	logged_user,err := usecase.UserRepo.LogIn(user)
	if err != nil {
		return domain.ResponseUser{},err
	}
	return domain.CreateResponseUser(logged_user),nil
}

func (usecase *UserUseCase) Register(user domain.RegisterUser) (domain.ResponseUser, error) {
	new_user,err := usecase.UserRepo.Register(user)
	if err != nil {
		return domain.ResponseUser{},err
	}
	return domain.CreateResponseUser(new_user),nil
}

func (usecase *UserUseCase) FilterUser(filter map[string]string) ([]domain.ResponseUser, error) {
	users,err := usecase.UserRepo.FilterUserDocument(filter)
	if err != nil {
		return []domain.ResponseUser{},err
	}

	response_users := []domain.ResponseUser{}

	for _,user := range users {
		response_users = append(response_users , domain.CreateResponseUser(user))
	}

	return response_users,nil
}