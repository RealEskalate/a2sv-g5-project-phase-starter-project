package Repository

import (
	"ASTU-backend-group-3/Blog_manager/Domain"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *blogRepository) AddReport(blogID string, report Domain.Report) error {

	blog, err := r.GetBlogByID(blogID)
	if err != nil {
		return err
	}

	blog.Reports = append(blog.Reports, report)
	_, err = r.collection.UpdateOne(context.TODO(), bson.M{"id": blogID}, bson.M{"$set": bson.M{"reports": blog.Reports}})
	if err != nil {
		return err
	}

	return nil
}


func (r *blogRepository) IsUserReported(blogID string, userID string) (bool) {
    blog, err := r.GetBlogByID(blogID)
    if err != nil {
        return false
    }
	fmt.Println(userID)

    for _, report := range blog.Reports {
        if report.UserId == userID {
			fmt.Println(report.UserId)
            return true
        }
    }

    return false
}
