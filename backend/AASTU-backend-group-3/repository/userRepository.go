package repository

import (
	"context"
	"errors"
	"group3-blogApi/domain"
	"group3-blogApi/infrastracture"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserRepositoryImpl struct {
	collection *mongo.Collection
}

func NewUserRepositoryImpl(coll *mongo.Collection) domain.UserRepository {
	return &UserRepositoryImpl{collection: coll}
}



func (ur *UserRepositoryImpl) Login(user *domain.User) (*domain.User, error) {
	var existingUser domain.User
	err := ur.collection.FindOne(context.Background(), map[string]string{"email": user.Email}).Decode(&existingUser)
	if err != nil {
		return &domain.User{}, err
	}
	return &existingUser, nil
	
}

func (ur *UserRepositoryImpl) GetUserByID(id string) (domain.User, error) {
	var user domain.User
	objID, err:= primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.User{}, err
	}

	err = ur.collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (ur *UserRepositoryImpl) DeleteRefreshToken(user *domain.User, token string) error {
	objID, err := primitive.ObjectIDFromHex(user.ID.Hex())
    if err != nil {
        return err
    }
    _, err = ur.collection.UpdateOne(
        context.Background(),
        bson.M{"_id": objID},
        bson.M{"$pull": bson.M{"refresh_tokens": bson.M{"token": token}}},
    )
    return err
}

func (ur *UserRepositoryImpl) UpdateUser(user *domain.User) error {
	
	_, err := ur.collection.UpdateOne(context.Background(), map[string]string{"email": user.Email}, bson.M{"$set": user})
	return err
}

func (ur *UserRepositoryImpl) DeleteAllRefreshTokens(user *domain.User) error {
	_, err := ur.collection.UpdateOne(context.Background(), map[string]string{"username": user.Username}, bson.M{"$set": bson.M{"refresh_tokens": []domain.RefreshToken{}}})
	return err
}


func (ur *UserRepositoryImpl) Register(user domain.User) error {

	// isUserExist := ur.collection.FindOne(context.Background(), map[string]string{"username": user.Username}).Err()
	// if isUserExist == nil {
	// 	return errors.New("user already exists")
	// }
	_, err := ur.collection.InsertOne(context.Background(), user)
	return err
}

func (ur *UserRepositoryImpl) GetUserByUsernameOrEmail(username, email string) (domain.User, error) {
	var user domain.User
	err := ur.collection.FindOne(context.Background(),  bson.M{"username": username, "email": email}).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}


func (ur *UserRepositoryImpl) AccountActivation(token string, email string) error {
	
	var user domain.User
	err := ur.collection.FindOne(context.Background(), map[string]string{"activation_token": token}).Decode(&user)
	if err != nil {
		return errors.New("invalid token or user not found")
	}

	if time.Since(user.TokenCreatedAt) > 24*time.Hour {
		return errors.New("token has expired")
	}

	

	// err = ur.collection.FindOneAndUpdate(context.Background(), bson.M{"_id": newID}, bson.M{"$set": user}).Decode(&updatedUser)
	_, err = ur.collection.UpdateOne(context.Background(), bson.M{"email": email}, bson.M{"$set": bson.M{"is_active": true}, "$unset": bson.M{"activation_token": ""}, "$currentDate": bson.M{"updated_at": true}})
	if err != nil {
		return errors.New("failed to activate account")
	}

	return nil
	

}










// rest password


func (ur *UserRepositoryImpl) GetUserByEmail(email string) (domain.User, error) {
	var user domain.User
	err := ur.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (ur *UserRepositoryImpl) GetUserByResetToken(token string) (domain.User, error) {
	var user domain.User
	// fmt.Println(token,"***************-----------------")	
	err := ur.collection.FindOne(context.Background(), bson.M{"password_reset_token": token}).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}




// password Hashing

func(uc *UserRepositoryImpl )HashPasswordRepo(password string) (string, error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (uc *UserRepositoryImpl) CheckPasswordHashRepo(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	
	return err == nil
}


////////////////////////////////////////////////////

func (ur *UserRepositoryImpl) SendReminderEmail() error {
    threeDaysAgo := time.Now().Add(-3 * 24 * time.Hour) // Adjust the threshold to 3 days ago

    filter := bson.M{
        "is_active": false,
        "token_created_at": bson.M{"$lt": threeDaysAgo, "$gte": time.Now().Add(-7 * 24 * time.Hour)},
    }

    // Find the users who need a reminder
    var users []domain.User
	cur, err := ur.collection.Find(context.Background(), filter)
	if err != nil {
		return err
	}
	err = cur.All(context.Background(), &users)
	if err != nil {
		return err
	}


    for _, user := range users {
        // Send activation email
        err := infrastracture.SendActivationEmail(user.Email, user.ActivationToken)
        if err != nil {
            log.Println("Failed to send reminder email to:", user.Email, err)
            continue
        }
    }
    return nil
}

// Delete users who have been inactive for more than 7 days
func (ur *UserRepositoryImpl) DeleteInActiveUser() error {
    sevenDaysAgo := time.Now().Add(-7 * 24 * time.Hour)

    filter := bson.M{
        "is_active": false,
        "token_created_at": bson.M{"$lt": sevenDaysAgo},
    }

    result, err := ur.collection.DeleteMany(context.Background(), filter)
    if err != nil {
        return err
    }

    log.Printf("Deleted %d inactive users", result.DeletedCount)
    return nil
}

// Schedule the process to send reminders and delete inactive users
func (ur *UserRepositoryImpl) ScheduleDeleteAndReminderForInActiveUser() {
    ticker := time.NewTicker(24 * time.Hour)
    go func() {
        for range ticker.C {
            // Send reminder emails to users who haven't activated within 3 days
            err := ur.SendReminderEmail()
            if err != nil {
                log.Println("Error sending reminder emails:", err)
            }

            // Delete users who have been inactive for more than 7 days
            err = ur.DeleteInActiveUser()
            if err != nil {
                log.Println("Error deleting inactive users:", err)
            }
        }
    }()
}