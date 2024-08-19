package usercommand

import (
	er "github.com/group13/blog/domain/errors"
	ihash "github.com/group13/blog/domain/i_hash"
	result "github.com/group13/blog/usecases_sof/user/result"
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




func (h *SignUpHandler) HandleLogin(command *LoginCommand) (*result.LoginInResult, error) {
	user, err := h.repo.FindByUsername(command.username)
	if err != nil {
		return nil, er.NewNotFound("user not found.")
	}

	if ok := h.repo.MatchPassword(command.password, user.PasswordHash(), h.hashService); !ok {

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

