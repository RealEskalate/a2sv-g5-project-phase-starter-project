package usecases

import repositories "blogs/Repositories"

type SignupUseCase struct {
	SignupRepository repositories.SignupRepository
}

func NewSignupUseCase(SignupRepository repositories.SignupRepository) *SignupUseCase {
	return &SignupUseCase{SignupRepository: SignupRepository}
}

func (u *SignupUseCase) Create() {
	
}