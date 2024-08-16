package Repositories

import "blogapp/Domain"

type blogrepository struct {
	collection Domain.Collection
}

func NewBlogRepository(_collection Domain.Collection) *blogrepository {
	return &blogrepository{

		collection: _collection,
	}

}
