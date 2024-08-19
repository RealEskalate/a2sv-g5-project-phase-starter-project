package usercommand

import (
	er "github.com/group13/blog/domain/errors"
	ihash "github.com/group13/blog/domain/i_hash"
	result "github.com/group13/blog/usecases_sof/user/result"
	icommand "github.com/group13/blog/usecases_sof/utils/command"
	iemail "github.com/group13/blog/usecases_sof/utils/i_email"
	ijwt "github.com/group13/blog/usecases_sof/utils/i_jwt"
	irepository "github.com/group13/blog/usecases_sof/utils/i_repo"
)

type LoginHandler struct {
	repo        irepository.UserRepository
	jwtService  ijwt.Services
	hashService ihash.Service
	emailService iemail.Service 
}

type LoginConfig struct {
	UserRepo irepository.UserRepository 
	jwtService   ijwt.Services 
	HashService  ihash.Service 
	emailService iemail.Service 
	
}

func NewLoginHandler(config LoginConfig) *LoginHandler {
	return &LoginHandler{
		repo:        config.UserRepo,
		jwtService:  config.jwtService,
		hashService: config.HashService,
		emailService: config.emailService,
	}
}

// Ensure Handler implements icmd.IHandler
var _ icommand.Ihandler[*LoginCommand, *result.LoginInResult] = &LoginHandler{}


func (h *LoginHandler) Handle(command *LoginCommand) (*result.LoginInResult, error) {
	user, err := h.repo.FindByUsername(command.username)
	if err != nil {
		return nil, er.NewNotFound("user not found.")
	}

	ok, err := h.hashService.Match(command.password, user.PasswordHash())
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, er.NewValidation("password is incorrect.")
	}

	user.MakeActive()
	if err := h.repo.Save(user); err != nil {
		return nil, err
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

