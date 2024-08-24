package entities

import mongopagination "github.com/gobeam/mongo-go-pagination"

type PaginatedResponse struct {
	Data     interface{}                    `json:"data"`
	MetaData mongopagination.PaginationData `json:"metadata"`
}
