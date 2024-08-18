package iuser

type IUserRepository interface {
	Save(string, bool) error
}
