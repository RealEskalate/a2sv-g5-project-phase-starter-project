package repositories

import (
	domain "blogs/Domain"
	"blogs/mongo"
	"context"
)

type SignupRepository struct {
	database mongo.Database
	collection string	
}
func NewSignupRepository(database mongo.Database , collection string) domain.SignupRepository {
	return &SignupRepository{database: database , 
							collection: collection}
}

func (r *SignupRepository) Create(user domain.User) (domain.User , error) { 
		collection := r.database.Collection(r.collection)

		_, err := collection.InsertOne(context.Background() , user)

		if err != nil {
			return domain.User{} , err
}

		return user , nil


}

func (r * SignupRepository) FindUserByEmail(c context.Context , email string) (domain.User , error) {
	collection := r.database.Collection(r.collection)
	var user domain.User
	err := collection.FindOne(c, domain.User{Email: email}).Decode(&user)
	if err != nil {
		return domain.User{} , err
	}
	return user , nil
}