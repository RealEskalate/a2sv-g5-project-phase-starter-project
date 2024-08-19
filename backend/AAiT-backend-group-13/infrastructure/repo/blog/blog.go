package blogrepo

import (
	"context"
	"time"

	"github.com/google/uuid"
	er "github.com/group13/blog/domain/errors"
	blogmodel "github.com/group13/blog/domain/models/blog"
	irepo "github.com/group13/blog/usecase/common/i_repo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Repo handles the persistence of user models.
type Repo struct {
	collection *mongo.Collection
}

// Ensure Repo implements irepo.User.
var _ irepo.Blog = &Repo{}

// NewRepo creates a new BlogRepo with the given MongoDB client, database name, and collection name.
func NewRepo(client *mongo.Client, dbName, collectionName string) *Repo {
	collection := client.Database(dbName).Collection(collectionName)
	return &Repo{
		collection: collection,
	
	}
}

func (b Repo) Save(blog *blogmodel.Blog) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": blog.ID()}
	update := bson.M{
		"$set": bson.M{
			"title":        blog.Title(),
			"content":      blog.Content(),
			"tags":         blog.Tags(),
			"createdDate":  blog.CreatedDate(),
			"updatedDate":  time.Now(),
			"userid":       blog.UserId(),
			"likeCount":    blog.LikeCount(),
			"disLikeCount": blog.DisLikeCount(),
			"commentCount": blog.CommentCount(),
		},
	}

	opts := options.Update().SetUpsert(true)
	_, err := b.collection.UpdateOne(ctx, filter, update, opts)

	if err != nil {
		return er.NewUnexpected(err.Error())
	}

	return nil

}

func (b Repo) Delete(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	result, err := b.collection.DeleteOne(ctx, filter)

	if err != nil {
		return er.NewUnexpected(err.Error())
	}
	if result.DeletedCount == 0 {
		return er.BlogNotFound
	}
	return nil
}

func (b Repo) GetAll() ([]*blogmodel.Blog, error) {
	return nil, nil
}

func (r Repo) GetSingle(id uuid.UUID) (*blogmodel.Blog, error) {
	return nil, nil
}
