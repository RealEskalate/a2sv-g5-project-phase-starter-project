package Domain

type User struct{
	Id string `json:"id" bson:"id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Email string `json:"email" bson:"email"`
	PostsIDs []string `json:"posts_id" bson:"posts_id"`
	ProfilePicture string `json:"profile_picture" bson:"profile_picture"`
	Bio string `json:"bio" bson:"bio"`
	Gender string `json:"gender" bson:"gender"`
	Role string `json:"role" bson:"role"`
	IsAdmin bool `json:"is_admin" bson:"is_admin"`
	IsActive bool `json:"is_active" bson:"is_active"`
	Adress string `json:"adress" bson:"adress"`
}



type  RegisterInput struct{
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Email string `json:"email" bson:"email"`
	ProfilePicture string `json:"profile_picture" bson:"profile_picture"`
	Bio string `json:"bio" bson:"bio"`
	Gender string `json:"gender" bson:"gender"`
	Adress string `json:"adress" bson:"adress"`
}

type LoginInput struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	
}

type UpdateUserInput struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	ProfilePicture string `json:"profile_picture" bson:"profile_picture"`
	Bio string `json:"bio" bson:"bio"`
	Adress string `json:"adress" bson:"adress"`
}


type ForgetPasswordInput struct {
	Email string `json:"email" bson:"email"`
}

type ResetPassworInput struct {
	NewPassword string `json:"password" bson:"password"`
}