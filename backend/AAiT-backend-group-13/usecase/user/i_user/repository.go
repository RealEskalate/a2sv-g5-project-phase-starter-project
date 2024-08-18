package iuser

type IUserRepository interface {
	Save(username string) error
}
