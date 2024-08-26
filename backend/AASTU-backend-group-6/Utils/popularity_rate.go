package utils

import "strings"

func PopularityRate(rateType string) int64 {
	var total int64
	if strings.ToLower(rateType) == "comment" {
		total = 3
	} else if strings.ToLower(rateType) == "reply" {
		total = 2
	} else if strings.ToLower(rateType) == "like" {
		total = 2
	} else if strings.ToLower(rateType) == "view" {
		total = 1
	} else if strings.ToLower(rateType) == "dislike" {
		total -= 2
	}
	return total
}
