package utils

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FilterReactionBlog(ids []primitive.ObjectID, reaction bool, isLiked bool, isDisliked bool) (primitive.M, primitive.M) {
	var update bson.M
	filter := bson.M{"_id": ids[1]}

	if reaction {
		if !isLiked && !isDisliked {
			update = bson.M{"$inc": bson.M{"like_count": PopularityRate("like")}}
		} else if isLiked {
			update = bson.M{"$inc": bson.M{"like_count": -PopularityRate("like")}}
		} else if isDisliked {
			update = bson.M{
				"$inc": bson.M{
					"like_count":    PopularityRate("like"),
					"dislike_count": -PopularityRate("dislike"),
				},
			}
		}
	} else {
		if !isLiked && !isDisliked {
			update = bson.M{"$inc": bson.M{"dislike_count": PopularityRate("dislike")}}
		} else if isDisliked {
			update = bson.M{"$inc": bson.M{"dislike_count": -PopularityRate("dislike")}}
		} else if isLiked {
			update = bson.M{
				"$inc": bson.M{
					"dislike_count": PopularityRate("dislike"),
					"like_count":    -PopularityRate("like"),
				},
			}
		}
	}

	// if reaction {
	// 	if !isLiked && !isDisliked {
	// 		update = bson.M{"$push": bson.M{"likes": ids[0]}}
	// 	} else if isLiked {
	// 		update = bson.M{"$pull": bson.M{"likes": ids[0]}}
	// 	} else if isDisliked {
	// 		update = bson.M{"$push": bson.M{"likes": ids[0]}, "$pull": bson.M{"dislikes": ids[0]}}
	// 	}
	// } else {
	// 	if !isLiked && !isDisliked {
	// 		update = bson.M{"$push": bson.M{"dislikes": ids[0]}}
	// 	} else if isDisliked {
	// 		update = bson.M{"$pull": bson.M{"dislikes": ids[0]}}
	// 	} else if isLiked {
	// 		update = bson.M{"$push": bson.M{"dislikes": ids[0]}, "$pull": bson.M{"likes": ids[0]}}
	// 	}
	// }
	return filter, update
}

func FilterReactionUser(ids []primitive.ObjectID, reaction bool, isLiked bool, isDisliked bool) (primitive.M, primitive.M) {
	var update bson.M
	filter := bson.M{"_id": ids[0]}

	if reaction {
		if !isLiked && !isDisliked {
			// User is liking a post for the first time
			update = bson.M{"$push": bson.M{"liked_posts_id": ids[1]}}
		} else if isLiked {
			// User is un-liking a post
			update = bson.M{"$pull": bson.M{"liked_posts_id": ids[1]}}
		} else if isDisliked {
			// User is switching from disliking to liking
			update = bson.M{
				"$push": bson.M{"liked_posts_id": ids[1]},
				"$pull": bson.M{"disliked_posts_id": ids[1]},
			}
		}
	} else {
		if !isLiked && !isDisliked {
			// User is disliking a post for the first time
			update = bson.M{"$push": bson.M{"disliked_posts_id": ids[1]}}
		} else if isDisliked {
			// User is un-disliking a post
			update = bson.M{"$pull": bson.M{"disliked_posts_id": ids[1]}}
		} else if isLiked {
			// User is switching from liking to disliking
			update = bson.M{
				"$push": bson.M{"disliked_posts_id": ids[1]},
				"$pull": bson.M{"liked_posts_id": ids[1]},
			}
		}
	}

	fmt.Println(filter, update)
	return filter, update

}
