package resetcodevalidate

import (
	"time"

	er "github.com/group13/blog/domain/errors"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	ijwt "github.com/group13/blog/usecase/common/i_jwt"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

type Handler struct {
	userrepo   irepo.User
	jwtService ijwt.Service
}

var _ icmd.IHandler[Command, string] = &Handler{}

func New(userrepo irepo.User, jwtService ijwt.Service) *Handler {
	return &Handler{userrepo: userrepo,
		jwtService: jwtService,
	}
}

func (h *Handler) Handle(cmd Command) (string, error) {
	user, err := h.userrepo.FindById(cmd.Id)
	if err != nil {
		return "", er.NewUnauthorized(err.Error())
	}

	resetCode := user.ResetCode()
	if resetCode == nil || resetCode.Code != cmd.Code {
		return "", er.NewUnauthorized("invalid code")
	}

	exprTime := user.ResetCode().Expr
	if exprTime.After(time.Now()) {
		return "", er.NewUnauthorized("time expierd")
	}

	if err = user.UpdateResetCode(nil); err != nil {
		return "", err
	}

	if err = h.userrepo.Save(user); err != nil {
		return "", err
	}

	token, err := h.jwtService.Generate(user, ijwt.Reset)
	if err != nil {
		return "", err
	}

	return token, nil
}
