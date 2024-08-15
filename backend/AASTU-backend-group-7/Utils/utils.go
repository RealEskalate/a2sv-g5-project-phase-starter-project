package Utils

import (
	"encoding/json"
	// "errors"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ReadJSON(ctx *gin.Context) (map[string]interface{}, error) {
	var jsonData map[string]interface{}
	err := json.NewDecoder(ctx.Request.Body).Decode(&jsonData)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func ObjectIdToString(objID primitive.ObjectID) string {
	return primitive.ObjectID.Hex(objID)
}

func StringToObjectId(str string) (primitive.ObjectID, error) {
	objID, err := primitive.ObjectIDFromHex(str)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return objID, nil
}

// func ExtractUser(c *gin.Context) (Domain.OmitedUser, error) {
// 	userID, ok := c.Get("user_id")
// 	if !ok {
// 		return Domain.OmitedUser{}, errors.New("Failed to retrieve user ID")
// 	}
// 	UserobjectID, err := primitive.ObjectIDFromHex(userID.(string))
// 	if err != nil {
// 		return Domain.OmitedUser{}, errors.New("invalid user ID")
// 	}
// 	is_admin, ok := c.Get("is_admin")
// 	if !ok {

// 		return Domain.OmitedUser{}, errors.New("Failed to retrieve role")
// 	}

// 	return Domain.OmitedUser{
// 		ID: UserobjectID,
// 		Is_Admin: is_admin.(bool),
// 	}, nil
// }
