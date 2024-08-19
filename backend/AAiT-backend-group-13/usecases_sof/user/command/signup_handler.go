package usercommand

import (
	er "github.com/group13/blog/domain/errors"
	ihash "github.com/group13/blog/domain/i_hash"
	usermodel "github.com/group13/blog/domain/models/user"
	result "github.com/group13/blog/usecases_sof/user/result"
	iemail "github.com/group13/blog/usecases_sof/utils/i_email"
	ijwt "github.com/group13/blog/usecases_sof/utils/i_jwt"
	irepository "github.com/group13/blog/usecases_sof/utils/i_repo"
)

type SignUpHandler struct {
	repo        irepository.UserRepository
	jwtService  ijwt.Services
	hashService ihash.Service
	emailService iemail.Service 
}

type SignUpConfig struct {
	UserRepo irepository.UserRepository 
	jwtService   ijwt.Services 
	HashService  ihash.Service 
	emailService iemail.Service 
	
}
// Ensure Handler implements icmd.IHandler
var _ icommand.IHandler[*Command, *blogmodel.Blog] = &Handler{}

func NewSignUpHandler(config SignUpConfig) *SignUpHandler {
	return &SignUpHandler{
		repo:        config.UserRepo,
		jwtService:  config.jwtService,
		hashService: config.HashService,
		emailService: config.emailService,
	}
}


func (h *SignUpHandler) Handle(command *signUpCommand) (*result.SignUpResult, error) {
	cfg := &usermodel.Config{
		Username:       command.username,
		Email:          command.email,
		PlainPassword:  command.password,
	}

	user, err := usermodel.New(*cfg)
	if err != nil {
		return nil, err
	}

	res,err := h.repo.FindByUsername(user.Username())
	if res != nil {
		return nil, er.NewConflict("username taken") 
	}

	if err != nil {
		return nil, err
	}

	res, err = h.repo.FindByEmail(user.Email())
	if res != nil {
		return nil, er.NewConflict("email already exists.") 
	}

	if err != nil {
		return nil, err
	}

	

	password := user.PasswordHash()
	user.UpdatePassword(password, h.hashService)
	validationLink, err := h.repo.GenerateValidationLink(*user)
	if err != nil {
		return nil, err
	}

	mails := []string{user.Email()}

	mail := iemail.NewSignUpEmail("", mails, validationLink) 

	if err := h.emailService.Send(mail); err != nil {
		return nil, err
	}

	if err := h.repo.Save(user); err != nil {
		return nil, err
	}

	return &result.SignUpResult{
		ID:        user.ID(),
		Username:  user.Username(),
		FirstName: user.FirstName(),
		LastName:  user.LastName(),
		IsAdmin:   user.IsAdmin(),
	}, nil
}
