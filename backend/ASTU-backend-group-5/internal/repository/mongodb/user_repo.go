package mongodb


import(
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
	"blogApp/internal/domain"
	"blogApp/internal/config"
	"blogApp/pkg/mongo"
	"errors"
)

type MongoUserRepo struct{
	Collection *mongo.Collection
}


func NewMongoUserRepo() (*MongoUserRepo, error){
	conn := config.Load()
	url := conn.MONGO_URI

	if url == ""{
		return nil, fmt.Errorf("error loading the url from environment")
	}

	client, err := mongo.NewMongoStorage(url)

	if err != nil {
		return nil, fmt.Errorf("error creating MongoDB client: %w", err)
	}

	
	NewUsercollection := client.Database("blogDb").Collection("users")

	if NewTaskcollection == nil {
		return nil, fmt.Errorf("failed to get tasks collection")
	}

	return &MongoTaskRepo{
		Collection: NewTaskcollection,
	}, nil
}

func (s *MongoUserRepo) CreateUser(data *domain.User) (data, error) {
	user, err := s.Collection.InsertOne(context.TODO(), data)
	return user, err
}

func (repo *MongoUserRepo) GetUserByEmail(email string) (*domain.User, error) {
    filter := bson.M{"email": email}
    
    var user domain.User

    err := repo.Collection.FindOne(context.TODO(), filter).Decode(&user)
    if err != nil {
        if errors.Is(err, mongo.ErrNoDocuments) {
            return nil, errors.New("user not found")
        }
        return nil, err
    }

    return &user, nil
}

func (repo *MongoUserRepo) GetUserByEmail(username string) (*User, error) {
    filter := bson.M{"username": username}
    var user User
    err := repo.Collection.FindOne(context.TODO(), filter).Decode(&user)

    if err != nil {
        if errors.Is(err, mongo.ErrNoDocuments) {
            return nil, errors.New("user not found")
        }
        return nil, err
    }

    return &user, nil
}


func (s *MongoUserRepo) NumberOfUsers() (int64, error){
	return s.Collection.CountDocuments(context.TODO(), bson.D{})
}