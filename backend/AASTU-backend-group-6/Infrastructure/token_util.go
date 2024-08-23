package infrastructure

import (
	"fmt"
	"time"

	domain "blogs/Domain"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateToken(claims jwt.Claims, secret string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}

// func CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
// 	exp := time.Now().Add(time.Hour * time.Duration(expiry))
// 	claimsRefresh := &domain.JwtCustomRefreshClaims{
// 		ID: user.ID.Hex(),
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(exp), // Convert expiration time to *jwt.NumericDate
// 		},
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
// 	rt, err := token.SignedString([]byte(secret))
// 	if err != nil {
// 		return "", err
// 	}
// 	return rt, err
// }

func IsAuthorized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractIDFromToken(requestToken string, secret string) (string, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	return claims["id"].(string), nil
}


func ExtractFromToken(requestToken string, secret string) (domain.User, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return domain.User{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return domain.User{}, fmt.Errorf("invalid token")
	}

	userMap, ok := claims["user"].(map[string]interface{})
	if !ok {
		return domain.User{}, fmt.Errorf("invalid user claims")
	}

	var user domain.User
	err = copier.Copy(&user, userMap)
	if err != nil {
		return domain.User{}, err
	}

	if id, ok := userMap["id"].(string); ok && id != "" {
		objectID, _ := primitive.ObjectIDFromHex(id)
		user.ID = objectID
	}


	if resetPasswordExpires, ok := userMap["reset_password_expires"].(string); ok && resetPasswordExpires != "" {
		user.ResetPasswordExpires, _ = time.Parse(time.RFC3339, resetPasswordExpires)
	}

	if postsIDs, ok := userMap["posts_id"].([]interface{}); ok {
		user.PostsID = convertToObjectIDArray(postsIDs)
	}

	if commentsIDs, ok := userMap["comments_id"].([]interface{}); ok {
		user.CommentsID = convertToObjectIDArray(commentsIDs)
	}

	if likedPostsIDs, ok := userMap["liked_posts_id"].([]interface{}); ok {
		user.LikedPostsID = convertToObjectIDArray(likedPostsIDs)
	}

	if dislikedPostsIDs, ok := userMap["disliked_posts_id"].([]interface{}); ok {
		user.DisLikePostsID = convertToObjectIDArray(dislikedPostsIDs)
	}

	// Set other fields, handling nil values
	if bio, ok := userMap["bio"].(string); ok {
		user.Bio = bio
	}
	if contact, ok := userMap["contact"].(string); ok {
		user.Contact = contact
	}
	if email, ok := userMap["email"].(string); ok {
		user.Email = email
	}
	if fullName, ok := userMap["full_name"].(string); ok {
		user.Full_Name = fullName
	}
	if googleID, ok := userMap["googleId"].(string); ok {
		user.GoogleID = googleID
	}
	if profileImage, ok := userMap["profile_image"].(string); ok {
		user.Profile_image_url = profileImage
	}
	if role, ok := userMap["roles"].(string); ok {
		user.Role = role
	}
	if resetPasswordToken, ok := userMap["reset_password_token"].(string); ok {
		user.ResetPasswordToken = resetPasswordToken
	}
	if password, ok := userMap["password"].(string); ok {
		user.Password = password
	}
	if username, ok := userMap["username"].(string); ok {
		user.Username = username
	}
	
fmt.Println(user)
	return user, nil
}

// Helper function to convert an array of interface{} to an array of ObjectID
func convertToObjectIDArray(input []interface{}) []primitive.ObjectID {
	var result []primitive.ObjectID
	for _, v := range input {
		if strID, ok := v.(string); ok {
			objectID, _ := primitive.ObjectIDFromHex(strID)
			result = append(result, objectID)
		}
	}
	return result
}
func ExtractFromTokenAuthClaim(requestToken string, secret string) (domain.JwtCustomClaims, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
	  if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	  }
	  return []byte(secret), nil
	})
  
	if err != nil {
	  return domain.JwtCustomClaims{}, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
  
	if !ok && !token.Valid {
	  return domain.JwtCustomClaims{}, fmt.Errorf("invalid token")
	}
	return domain.JwtCustomClaims{
	  ID: claims["id"].(string),
	}, nil
  }
  