package domain

// import (
// 	"context"

// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// type Admin struct {
// 	ID        primitive.ObjectID  `bson:"_id,omitempty" json:"id,omitempty"`
// 	Username  string              `bson:"username" json:"username"`
// 	Email     string              `bson:"email" json:"email"`
// 	Password  string              `bson:"password" json:"password"`
// 	Bio       string              `bson:"bio,omitempty" json:"bio,omitempty"`
// 	Role      Role                `bson:"role" json:"role"`
// 	CreatedAt primitive.Timestamp `bson:"createdAt" json:"createdAt"`
// 	UpdatedAt primitive.Timestamp `bson:"updatedAt" json:"updatedAt"`
// }

// type AdminUsecase interface {
// 	// Admin Profile
// 	GetProfile(ctx context.Context, adminID string) (User, error)
// 	UpdateProfile(ctx context.Context, admin User) error
// 	UploadImage(ctx context.Context, adminID string, imagePath string) error
// 	Logout(ctx context.Context, adminID string) error
// 	ResetPassword(ctx context.Context, token, newPassword string) error
// 	RefreshToken(ctx context.Context, token string) (string, error)

// 	// User Management
// 	GetUsers(ctx context.Context) ([]User, error)
// 	GetUser(ctx context.Context, userID string) (User, error)
// 	DeleteUser(ctx context.Context, userID string) error
// 	UpdateUserRole(ctx context.Context, userID, role string) error

// 	// Blog Management

// }

// type AdminRepository interface {
// 	// Admin Profile
// 	GetProfile(ctx context.Context, adminID string) (User, error)
// 	UpdateProfile(ctx context.Context, admin User) error
// 	UploadImage(ctx context.Context, adminID string, imagePath string) error
// 	Logout(ctx context.Context, adminID string) error
// 	ResetPassword(ctx context.Context, token, newPassword string) error
// 	RefreshToken(ctx context.Context, token string) (string, error)

// 	// User Management
// 	GetUsers(ctx context.Context) ([]User, error)
// 	GetUser(ctx context.Context, userID string) (User, error)
// 	DeleteUser(ctx context.Context, userID string) error
// 	UpdateUserRole(ctx context.Context, userID, role string) error

// 	// Blog Management
// }
