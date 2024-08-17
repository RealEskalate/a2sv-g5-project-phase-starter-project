package repositories

import (
	domain "blogs/Domain"
	"blogs/mongo"
	"context"
	"fmt"

	// "go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

type SignupRepository struct {
	database mongo.Database
	collection string	
}
func NewSignupRepository(database mongo.Database , collection string) domain.SignupRepository {
	return &SignupRepository{database: database , 
							collection: collection}
}

func (r *SignupRepository) Create(c context.Context ,user domain.User) (domain.User , error) { 
	collection := r.database.Collection(r.collection)
	_, err := collection.InsertOne(c, user)

	if err != nil {
		return domain.User{} , err
	}

	return user , nil
}

func (r * SignupRepository) FindUserByEmail(c context.Context , email string) (domain.User , error) {
	collection := r.database.Collection(r.collection)
	var user domain.User
	err := collection.FindOne(c, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return domain.User{} , err
	}
	return user , nil
}

// the function that sets the otp code to the db 
func (r *SignupRepository) SetOTP(c context.Context , email string , otp string) (error) { 
	collection := r.database.Collection(r.collection)
	// data we put in
	
	data := bson.M{"$set": bson.M{"otp": otp}}  // Corrected the key
	_ , err := collection.UpdateOne(c, bson.M{"email": email}, data)
	if err != nil {
		return err
	}

	return nil
}


// verify OTP code

func (r *SignupRepository) VerifyOTP(c context.Context , email string , otp string) error { 
	collection := r.database.Collection(r.collection)
	// data we put in getting it and checking from the current otp 
	var user domain.User
	err := collection.FindOne(c , bson.M{otp: otp}).Decode(&user)

	if err != nil { 
		return err
	}
	return nil
}

// update user
func (r *SignupRepository) VerifyUser(c context.Context ,  user domain.User) (domain.User , error) { 
	collection := r.database.Collection(r.collection)
	
	update := bson.M{"$set": user}
	_, err := collection.UpdateOne(c , bson.M{"email": user.Email} , update)
	if err != nil {
		return domain.User{} , err
}
	return user , nil

}