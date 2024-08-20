package logincommand

import (
	er "github.com/group13/blog/domain/errors"
	ihash "github.com/group13/blog/domain/i_hash"
	// icmd "github.com/group13/blog/usecase/usecase/common/cqrs/command"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	iemail "github.com/group13/blog/usecase/common/i_email"
	ijwt "github.com/group13/blog/usecase/common/i_jwt"
	irepo "github.com/group13/blog/usecase/common/i_repo"
	result "github.com/group13/blog/usecase/user/result"
)

type LoginHandler struct {
	repo         irepo.UserRepository
	jwtService   ijwt.Service
	hashService  ihash.Service
	emailService iemail.Service
}

type LoginConfig struct {
	UserRepo     irepo.UserRepository
	JwtService   ijwt.Service
	HashService  ihash.Service
	EmailService iemail.Service
}

func NewLoginHandler(config LoginConfig) *LoginHandler {
	return &LoginHandler{
		repo:         config.UserRepo,
		jwtService:   config.JwtService,
		hashService:  config.HashService,
		emailService: config.EmailService,
	}
}

// Ensure Handler implements icmd.IHandler

var _ icmd.IHandler[*LoginCommand, *result.LoginInResult] = &LoginHandler{}

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
		Token:        token,
		Refreshtoekn: refreshtoken,
	}, nil

}
