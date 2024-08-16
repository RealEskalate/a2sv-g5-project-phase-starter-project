package domain

import "time"

type User struct {
  FirstName  string    `json:"firstname"`
  LastName   string    `json:"lastname"`
  Bio        string    `json:"bio"`
  Avatar     string    `json:"avatar"`
  UserName   string    `json:"username" binding:"required"`
  Password   string    `json:"password" binding:"required"`
  Email      string    `json:"email" binding:"required"`
  Role       string    `json:"role" `
  Address    string    `json:"address"`
  JoinedDate time.Time `json:"joinedDate"`
}

type Token struct {
  Username  string    `json:"username" bson:"username"`
  ExpiresAt time.Time `json:"expires_at" bson:"expires_at"`
}

type UserRepository interface {
  CheckUsernameAndEmail(username, email string) error
  Register(user *User) error
  GetUserByUsernameorEmail(usernameoremail string) (*User, error)
  UpdateProfile(usernameoremail string, user *User) error
  Resetpassword(usernameoremail string, password string) error
  InsertToken(token *Token) error
  GetTokenByUsername(username string) (*Token, error)
}

type UserUsecase interface {
  RegisterUser(user *User) error
  LoginUser(usernameoremail, password string) (string, error)
  UpdateProfile(usernameoremail string, user *User) error
  ResetPassword(usernameoremail, password string) error
  LogoutUser(username string) error
  ForgotPassword(email string) error
}
