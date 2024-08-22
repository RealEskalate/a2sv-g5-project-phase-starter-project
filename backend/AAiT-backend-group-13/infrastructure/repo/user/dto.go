package userrepo

import (
	"time"

	"github.com/google/uuid"
	"github.com/group13/blog/domain/models"
)

type ResetCodeDto struct {
	CodeHash string    `bson:"code"`
	Expr     time.Time `bson:"expr"`
}
type UserDTO struct {
	ID        uuid.UUID     `bson:"_id"`
	Username  string        `bson:"username"`
	Password  string        `bson:"passwordHash"`
	Email     string        `bson:"email"`
	IsActive  bool          `bson:"isActive"`
	IsAdmin   bool          `bson:"isAdmin"`
	FirstName string        `bson:"firstName"`
	LastName  string        `bson:"lastName"`
	CreatedAt time.Time     `bson:"createdAt"`
	UpdatedAt time.Time     `bson:"updatedAt"`
	ResetCode *ResetCodeDto `bson:"resetCode"`
}

func FromUser(u *models.User) *UserDTO {
	var resetCodeDto ResetCodeDto

	if u.ResetCode() != nil {
		resetCodeDto = ResetCodeDto{
			CodeHash: u.ResetCode().CodeHash,
			Expr:     u.ResetCode().Expr,
		}
	}

	return &UserDTO{
		ID:        u.ID(),
		Username:  u.Username(),
		Password:  u.PasswordHash(),
		Email:     u.Email(),
		IsActive:  u.IsActive(),
		IsAdmin:   u.IsAdmin(),
		FirstName: u.FirstName(),
		LastName:  u.LastName(),
		CreatedAt: u.CreatedAt(),
		UpdatedAt: u.UpdatedAt(),
		ResetCode: &resetCodeDto,
	}
}

func ToUser(dto *UserDTO) *models.User {
	var resetCode models.ResetCode

	if dto.ResetCode != nil {
		resetCode = models.ResetCode{
			CodeHash: dto.ResetCode.CodeHash,
			Expr:     dto.ResetCode.Expr,
		}

	}

	user := models.MapUser(models.MapUserConfig{
		Id:             dto.ID,
		FirstName:      dto.FirstName,
		LastName:       dto.LastName,
		Username:       dto.Username,
		IsAdmin:        dto.IsAdmin,
		IsActive:       dto.IsActive,
		Email:          dto.Email,
		HashedPassword: dto.Password,
		CreatedAt:      dto.CreatedAt,
		UpdatedAt:      dto.UpdatedAt,
		ResetCode:      &resetCode,
	})
	return user
}

