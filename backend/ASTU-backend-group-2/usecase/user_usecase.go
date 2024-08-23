package usecase

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain"
	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
)

type userUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (uu *userUsecase) CreateUser(c context.Context, user *domain.User) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.CreateUser(ctx, user)
}

func (uu *userUsecase) GetUserByEmail(c context.Context, email string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.GetUserByEmail(ctx, email)
}

func (uu *userUsecase) GetUserById(c context.Context, userId string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.GetUserById(ctx, userId)
}
func (uu *userUsecase) GetUsers(c context.Context, userFilter domain.UserFilter) (*[]domain.UserOut, mongopagination.PaginationData, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	filter := UserFilterOption(userFilter)
	users, meta, err := uu.userRepository.GetUsers(ctx, filter, userFilter)

	if err != nil {
		return nil, mongopagination.PaginationData{}, err
	}

	// map users to userout
	res := make([]domain.UserOut, 0)

	for _, user := range *users {
		res = append(res, *user.ToUserOut())
	}

	return &res, meta, nil
}

func (uu *userUsecase) UpdateUser(c context.Context, userID string, updatedUser *domain.UserUpdate) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	return uu.userRepository.UpdateUser(ctx, userID, updatedUser)
}

func (uu *userUsecase) ActivateUser(c context.Context, userID string) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	_, err := uu.userRepository.ActivateUser(ctx, userID)
	return err
}

func (uu *userUsecase) DeleteUser(c context.Context, userID string) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.DeleteUser(ctx, userID)
}

func (uu *userUsecase) IsUserActive(c context.Context, userID string) (bool, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.IsUserActive(ctx, userID)
}

func (uu *userUsecase) IsOwner(c context.Context) (bool, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.IsOwner(ctx)
}

func (uu *userUsecase) ResetUserPassword(c context.Context, userID string, resetPassword *domain.ResetPasswordRequest) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.ResetUserPassword(ctx, userID, resetPassword)
}

func (uu *userUsecase) UpdateUserPassword(c context.Context, userID string, updatePassword *domain.UpdatePassword) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.UpdateUserPassword(ctx, userID, updatePassword)
}

func (uu *userUsecase) PromoteUserToAdmin(c context.Context, userID string) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.PromoteUserToAdmin(ctx, userID)
}

func (uu *userUsecase) DemoteAdminToUser(c context.Context, userID string) error {
	user, err := uu.GetUserById(c, userID)
	if err != nil {
		return err
	}
	if user.IsOwner {
		return errors.New("cannot demote owner")
	}
	return uu.userRepository.DemoteAdminToUser(c, userID)
}
func (uu *userUsecase) UpdateProfilePicture(c context.Context, userID string, filename string) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	_, err := uu.GetUserById(c, userID)
	if err != nil {
		return err
	}
	return uu.userRepository.UpdateProfilePicture(ctx, userID, filename)
}

func UserFilterOption(filter domain.UserFilter) bson.M {

	query := bson.M{
		"$match": bson.M{},
	}
	semiquery := query["$match"].(bson.M)

	// Email filter
	if filter.Email != "" {
		semiquery["email"] = bson.M{"$regex": filter.Email, "$options": "i"}
	}

	// filter.Role
	if filter.Role != "" {
		semiquery["role"] = filter.Role
	}

	// Active filter
	if filter.Active != "" {
		semiquery["active"] = filter.Active == "true"
	}

	// Bio filter
	if filter.Bio != "" {
		semiquery["bio"] = bson.M{"$regex": filter.Bio, "$options": "i"} // case-insensitive search
	}

	// First name filter
	if filter.FirstName != "" {
		semiquery["first_name"] = bson.M{"$regex": filter.FirstName, "$options": "i"} // case-insensitive search
	}

	// Last name filter
	if filter.LastName != "" {
		semiquery["last_name"] = bson.M{"$regex": filter.LastName, "$options": "i"} // case-insensitive search
	}

	// Is owner filter
	if filter.IsOwner != "" {
		semiquery["is_owner"] = filter.IsOwner == "true"
	}

	// Is admin filter
	if filter.Role != "" {
		semiquery["role"] = filter.Role
	}

	// Date range filter
	if !filter.DateFrom.IsZero() && !filter.DateTo.IsZero() {
		semiquery["created_at"] = bson.M{
			"$gte": filter.DateFrom,
			"$lte": filter.DateTo,
		}
	} else if !filter.DateFrom.IsZero() {
		semiquery["created_at"] = bson.M{"$gte": filter.DateFrom}
	} else if !filter.DateTo.IsZero() {
		semiquery["created_at"] = bson.M{"$lte": filter.DateTo}
	}

	log.Println(query)
	return query

}
