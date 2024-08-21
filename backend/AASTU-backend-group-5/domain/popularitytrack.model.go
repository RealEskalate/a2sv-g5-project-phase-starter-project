package domain

import (
	"errors"
)

type SortBy string

const (
	SortByLikeCount    SortBy = "likeCount"
	SortByCommentCount SortBy = "commentCount"
	SortByPublishDate  SortBy = "createdAt"
	SortByEngagement   SortBy = "engagement"
	SortByDislikeCount SortBy = "dislikeCount"
)

type SortOrder int

const (
	SortOrderAscending  SortOrder = 1
	SortOrderDescending SortOrder = -1
)

var ErrInvalidSortBy = errors.New("invalid sort by value")
