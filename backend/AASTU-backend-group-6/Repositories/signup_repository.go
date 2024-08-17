package repositories

import (
	domain "blogs/Domain"
	"blogs/mongo"
	"context"

	// "go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

type signupRepository struct {
	database   mongo.Database
	collection string
}

// Create implements domain.SignupRepository.
func (s *signupRepository) Create(ctx context.Context, user domain.User) (domain.User, error) {
	collection := s.database.Collection(s.collection)

	_, err := collection.InsertOne(ctx, user)

	if err != nil {
		return domain.User{}, err
	}
	return user , nil
}

// the function that sets the otp code to the db 
func (r *signupRepository) SetOTP(c context.Context , email string , otp string) (error) { 
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

func (r *signupRepository) VerifyOTP(c context.Context , email string , otp string) error { 
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
func (r *signupRepository) VerifyUser(c context.Context ,  user domain.User) (domain.User , error) { 
	collection := r.database.Collection(r.collection)
	
	update := bson.M{"$set": user}
	_, err := collection.UpdateOne(c , bson.M{"email": user.Email} , update)
	if err != nil {
		return domain.User{} , err
}
	return user , nil

}
