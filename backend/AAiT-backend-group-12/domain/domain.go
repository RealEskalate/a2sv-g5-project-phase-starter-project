package domain

import "github.com/gin-gonic/gin"

/*
Defines the names of the collections in the DB
*/
const (
	CollectionUsers = "users"
	CollectionBlogs = "blogs"
)

type Response gin.H
