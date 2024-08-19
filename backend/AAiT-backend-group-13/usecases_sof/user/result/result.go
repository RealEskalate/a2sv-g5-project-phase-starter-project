package result

import "github.com/google/uuid"

type SignUpResult struct {
	ID       uuid.UUID 
	Username string    
	FirstName string 
	LastName string    
	IsAdmin  bool

}


func NewSignUpResult(ID uuid.UUID, Username string, firstName string, lastName string, IsAdmin bool) SignUpResult {
	return SignUpResult{
		ID:        ID,
		Username:  Username,
		FirstName: firstName,
		LastName:  lastName,
		IsAdmin:   IsAdmin,
	}
}


type LoginInResult struct {	
	Token string 
	Refreshtoekn string 
}

func NewLoginInResult(token string, refreshtoekn string) LoginInResult {
	return LoginInResult{
		Token: token,
		Refreshtoekn: refreshtoekn,
	}
}