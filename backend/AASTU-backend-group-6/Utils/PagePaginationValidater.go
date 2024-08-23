package utils

import (
	"errors"
	"strconv"
)

func PagePaginationValidater(pageNo, pageSize string) (int64, int64, error) {
	if pageNo == "" || pageSize == "" {
		return 0, 0, errors.New("invalid pageNo or pageSize")
	}

	PageNo, err := strconv.ParseInt(pageNo, 10, 64)
	if err != nil {
		return 0, 0, errors.New("invalid pageNo or pageSize")
	}
	PageSize, err := strconv.ParseInt(pageSize, 10, 64)
	if err != nil {
		return 0, 0, errors.New("invalid pageNo or pageSize")
	}
	return PageNo, PageSize, nil
}
