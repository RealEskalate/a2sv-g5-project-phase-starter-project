package infrastructure

import domain "aait-backend-group4/Domain"

// CalculatePopularity calculates the popularity score of a feedback based on various factors.
// It takes a pointer to a Feedback object as input and returns a float64 value representing the popularity score.
// The popularity score is calculated using the following formula:
//
//	val := (float64(fb.Likes) * 2) + (float64(fb.View_count) * 1.2) + (float64(len(fb.Comments)) * 1.5) - float64(fb.Dislikes)
//	where fb.Likes represents the number of likes, fb.View_count represents the number of views,
//	len(fb.Comments) represents the number of comments, and fb.Dislikes represents the number of dislikes.
//
// The calculated popularity score is returned as a float64 value.
func CalculatePopularity(fb *domain.Feedback) float64 {
	val := (float64(fb.Likes) * 2) + (float64(fb.View_count) * 1.2) +
		(float64(len(fb.Comments)) * 1.5) - float64(fb.Dislikes)
	return val

}
