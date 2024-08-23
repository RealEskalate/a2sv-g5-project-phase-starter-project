package utils

import (
	domain "blogs/Domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func IsAlreadyReacted(post *domain.Blog, user_id primitive.ObjectID) (bool, bool) {
	// var liked, disLike bool
	// for _, like := range post.Likes {
	// 	if like == user_id {
	// 		liked = true
	// 		break
	// 	}
	// }
	// for _, dislike := range post.DisLikes {
	// 	if dislike == user_id {
	// 		disLike = true
	// 		break
	// 	}
	// }
	return true, true
}
