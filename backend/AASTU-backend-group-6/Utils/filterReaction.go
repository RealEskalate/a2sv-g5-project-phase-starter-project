package utils

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FilterReactionBlog(ids []primitive.ObjectID, reaction bool, isLiked bool, isDisliked bool) (primitive.M, primitive.M) {
	var update bson.M
	filter := bson.M{"_id": ids[1]}
	if reaction{
		if !isLiked && !isDisliked{
			update = bson.M{"$push": bson.M{"likes": ids[0]}}
		}else if isDisliked{
			update = bson.M{"$push": bson.M{"likes": ids[0]}, "$pull": bson.M{"dislikes": ids[0]}}
		}
    } else{
		if !isLiked && !isDisliked{
			update = bson.M{"$push": bson.M{"dislikes": ids[0]}}
		}else if isLiked{
			update = bson.M{"$push": bson.M{"dislikes": ids[0]}, "$pull": bson.M{"likes": ids[0]}}
		}
	}
	return filter, update
}

func FilterReactionUser(ids []primitive.ObjectID, reaction bool, isLiked bool, isDisliked bool) (primitive.M, primitive.M) {
	var update bson.M
	filter := bson.M{"_id": ids[0], "posts._id" : ids[2]}
	if reaction{
		if !isLiked && !isDisliked{
			update = bson.M{"$push": bson.M{"posts.$.likes": ids[1]}}
		}else if isDisliked{
			update = bson.M{"$push": bson.M{"posts.$.likes": ids[1]}, "$pull": bson.M{"posts.$.dislikes": ids[1]}}
		}
    } else{
		if !isLiked && !isDisliked{
			update = bson.M{"$push": bson.M{"posts.$.dislikes": ids[1]}}
		}else if isLiked{
			update = bson.M{"$push": bson.M{"posts.$.dislikes": ids[1]}, "$pull": bson.M{"posts.$.likes": ids[1]}}
		}
	}
	fmt.Println(filter, update)
	return filter, update
}