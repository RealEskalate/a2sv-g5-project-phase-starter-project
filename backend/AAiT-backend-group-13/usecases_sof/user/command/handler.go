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

type handler struct {
	repo        irepository.UserRepository
	jwtService  ijwt.Services
	hashService ihash.Service
	emailService iemail.Service 
}

type Config struct {
	UserRepo irepository.UserRepository 
	jwtService   ijwt.Services 
	HashService  ihash.Service 
	emailService iemail.Service 
	
}

func NewHandler(config Config) *handler {
	return &handler{
		repo:        config.UserRepo,
		jwtService:  config.jwtService,
		hashService: config.HashService,
		emailService: config.emailService,
	}
}


func (h *handler) HandleSignup(command *signUpCommand) (*result.SignUpResult, error) {
	cfg := &usermodel.Config{
		Username:       command.username,
		Email:          command.email,
		PlainPassword:  command.password,
	}

	user, err := usermodel.New(*cfg)
	if err != nil {
		return nil, err
	}

	if err := h.repo.CheckUsernameAvailability(user.Username()); err != nil {
		return nil, err
	}

	if err := h.repo.CheckEmailAvailability(user.Email()); err != nil {
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

func (h *handler) HandleLogin(command *LoginCommand) (*result.LoginInResult, error) {
	user, err := h.repo.FindByUsername(command.username)
	if err != nil {
		return nil, er.NewNotFound("user not found.")
	}

	if ok := h.repo.MatchPassword(command.password, user.PasswordHash(), h.hashService); !ok {

		return nil, er.NewValidation("password is incorrect.")
	}

	token, err := h.jwtService.Generate(user, ijwt.Access)
	if err != nil {
		return nil, err
	}

	refreshtoken, err := h.jwtService.Generate(user, ijwt.Refresh)

	if err != nil {	
		return nil, err
	}

	return &result.LoginInResult{
		Token:       token,
		Refreshtoekn: refreshtoken,
	}, nil
	
}

