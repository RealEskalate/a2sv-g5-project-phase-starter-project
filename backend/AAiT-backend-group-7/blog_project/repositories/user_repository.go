package repositories

import (
	"blog_project/domain"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) domain.IUserRepository {
	return &userRepository{collection: collection}
}

func (userRepo *userRepository) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	cursor, err := userRepo.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []domain.User
	for cursor.Next(ctx) {
		var user domain.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (userRepo *userRepository) GetUserByID(ctx context.Context, id int) (domain.User, error) {
	var user domain.User
	result := userRepo.collection.FindOne(ctx, bson.M{"id": id})

	if err := result.Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, err
	}

	return user, nil
}

func (userRepo *userRepository) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	result, err := userRepo.collection.InsertOne(ctx, user)
	if err != nil {
		return domain.User{}, err
	}

	if id, ok := result.InsertedID.(int); ok {
		user.ID = id
	} else {
		return domain.User{}, errors.New("failed to convert inserted ID to int")
	}

	return user, nil
}

func (userRepo *userRepository) UpdateUser(ctx context.Context, id int, user domain.User) (domain.User, error) {
	var updatedUser domain.User
	result := userRepo.collection.FindOneAndUpdate(
		ctx,
		bson.M{"id": id},
		bson.M{"$set": user},
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	)

	if err := result.Decode(&updatedUser); err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, err
	}

	return updatedUser, nil
}

func (userRepo *userRepository) DeleteUser(ctx context.Context, id int) error {
	result, err := userRepo.collection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("user not found")
	}

	return nil
}

func (userRepo *userRepository) SearchByUsername(ctx context.Context, username string) (domain.User, error) {
	var user domain.User
	err := userRepo.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, err
	}

	return user, nil
}

func (userRepo *userRepository) SearchByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User
	err := userRepo.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, err
	}

	return user, nil
}

func (userRepo *userRepository) AddBlog(ctx context.Context, userID int, blog domain.Blog) (domain.User, error) {
	var user domain.User
	err := userRepo.collection.FindOne(ctx, bson.M{"id": userID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, err
	}

	user.Blogs = append(user.Blogs, blog.ID)
	_, err = userRepo.collection.UpdateOne(ctx, bson.M{"id": userID}, bson.M{"$set": bson.M{"blogs": user.Blogs}})
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (userRepo *userRepository) StoreRefreshToken(ctx context.Context, userID int, refreshToken string) error {
	_, err := userRepo.collection.UpdateOne(ctx, bson.M{"id": userID}, bson.M{"$set": bson.M{"refresh_token": refreshToken}})
	return err
}

func (userRepo *userRepository) ValidateRefreshToken(ctx context.Context, userID int, refreshToken string) (bool, error) {
	var user domain.User
	err := userRepo.collection.FindOne(ctx, bson.M{"id": userID, "refresh_token": refreshToken}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
