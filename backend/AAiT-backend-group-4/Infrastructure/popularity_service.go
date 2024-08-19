package infrastructure

import domain "aait-backend-group4/Domain"

func CalculatePopularity(fb *domain.Feedback) float64 {
	val := (float64(fb.Likes) * 2) + (float64(fb.View_count) * 1.2) +
		(float64(len(fb.Comments)) * 1.5) - float64(fb.Dislikes)
	return val

}
