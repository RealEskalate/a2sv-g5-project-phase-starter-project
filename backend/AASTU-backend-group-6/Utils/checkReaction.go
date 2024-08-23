package utils

import (
	domain "blogs/Domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func IsAlreadyReacted(user domain.User, user_id primitive.ObjectID) (bool, bool) {
	var liked, disLike bool
	for _, like := range user.LikedPostsID {
		if like == user_id {
			liked = true
			break
		}
	}
	for _, dislike := range user.DisLikePostsID {
		if dislike == user_id {
			disLike = true
			break
		}
	}
	return liked, disLike
}
