package repository_test

import (
	// "blogs/bootstrap"
	"blogs/bootstrap"
	"blogs/domain"
	"blogs/mocks"
	"blogs/repository"
	"context"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
)

type BlogRepositoryTestSuite struct {
    suite.Suite
    client *mongo.Client
    db *mongo.Database
    blogCollection *mongo.Collection
    viewCollection *mongo.Collection
    likeCollection    *mongo.Collection
	commentCollection *mongo.Collection
    cache *mocks.Cache
    repository repository.BlogRepository
}    		

func (suite *BlogRepositoryTestSuite) SetupSuite() {
	// Initialize environment variables and MongoDB connection
	// bootstrap.InitEnv()
	// env, err := bootstrap.GetEnv("MONGO_URI")
	// if err != nil {
	// 	suite.T().Fatalf("Failed to get MongoDB URI from environment: %v", err)
	// }
	client, err := bootstrap.ConnectToMongoDB("mongodb+srv://nathnaeldes:12345678n@cluster0.w8bpdtf.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")
	if err != nil {
		suite.T().Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Initialize MongoDB collections
	suite.client = client
	suite.db = suite.client.Database("test")
	suite.blogCollection = suite.db.Collection("blogs_test")
	suite.viewCollection = suite.db.Collection("views_test")
	suite.likeCollection = suite.db.Collection("likes_test")
	suite.commentCollection = suite.db.Collection("comments_test")

    suite.cache = &mocks.Cache{}
    repoository := repository.NewBlogRepository(suite.db, suite.cache).(*repository.BlogRepository)
    suite.repository = *repoository
}


func (suite *BlogRepositoryTestSuite) TestInsertBlog() {
    // Initialize a blog
    blog := &domain.Blog{
        Title:         "Test Blog",
        Content:       "Test Content",
        Author:       "Test Author",
        Tags:          []string{"test", "blog"},
        CreatedAt:     time.Now(),
        LastUpdatedAt: time.Now(),
        ViewsCount:    0,
        LikesCount:    0,
        CommentsCount: 0,
    }

    // Insert the blog
    result, err := suite.repository.InsertBlog(blog)
    log.Println(result,"result")
    suite.Require().NoError(err)
    suite.Require().NotNil(result)

    
}

func (suite *BlogRepositoryTestSuite) TestGetBlogByID() {
    // Initialize a blog
    blog := &domain.Blog{
        Title:         "Test Blog",
        Content:       "Test Content",
        Author:       "Test Author",
        Tags:          []string{"test", "blog"},
        CreatedAt:     time.Now(),
        LastUpdatedAt: time.Now(),
        ViewsCount:    0,
        LikesCount:    0,
        CommentsCount: 0,
    }

    // Insert the blog
    result, err := suite.repository.InsertBlog(blog)
    suite.Require().NoError(err)
    suite.Require().NotNil(result)

    // Get the blog by ID
    insertedBlog, err := suite.repository.GetBlogByID(result.ID.Hex())
    suite.Require().NoError(err)
    suite.Require().NotNil(insertedBlog)
    
    
}

func (suite *BlogRepositoryTestSuite) TestUpdateBlogByID() {
    // Initialize a blog
    blog := &domain.Blog{
        Title:         "Test Blog",
        Content:       "Test Content",
        Author:       "Test Author",
        Tags:          []string{"test", "blog"},
        CreatedAt:     time.Now(),
        LastUpdatedAt: time.Now(),
        ViewsCount:    0,
        LikesCount:    0,
        CommentsCount: 0,
    }

    // Insert the blog
    result, err := suite.repository.InsertBlog(blog)
    suite.Require().NoError(err)
    suite.Require().NotNil(result)
    
    // Update the blog
    result.Title = "Updated Blog"
    result.Content = "Updated Content"
    result.Author = "Updated Author"
    result.Tags = []string{"updated", "blog"}
    result.LastUpdatedAt = time.Now()
    err = suite.repository.UpdateBlogByID(result.ID.Hex(), result)
    suite.Require().NoError(err)

    
}

func (suite *BlogRepositoryTestSuite) TestDeleteBlogByID() {
    // Initialize a blog
    blog := &domain.Blog{
        Title:         "Test Blog",
        Content:       "Test Content",
        Author:       "Test Author",
        Tags:          []string{"test", "blog"},
        CreatedAt:     time.Now(),
        LastUpdatedAt: time.Now(),
        ViewsCount:    0,
        LikesCount:    0,
        CommentsCount: 0,
    }

    // Insert the blog
    result, err := suite.repository.InsertBlog(blog)
    suite.Require().NoError(err)
    suite.Require().NotNil(result)

    // Delete the blog
    err = suite.repository.DeleteBlogByID(result.ID.Hex())
    suite.Require().NoError(err)

}

func (suite *BlogRepositoryTestSuite) TestSearchBlog() {
    // Initialize a blog
    blog := &domain.Blog{
        Title:         "Test Blog",
        Content:       "Test Content",
        Author:       "Test Author",
        Tags:          []string{"test", "blog"},
        CreatedAt:     time.Now(),
        LastUpdatedAt: time.Now(),
        ViewsCount:    0,
        LikesCount:    0,
        CommentsCount: 0,
    }

    // Insert the blog
    result, err := suite.repository.InsertBlog(blog)
    suite.Require().NoError(err)
    suite.Require().NotNil(result)

    // Search for the blog
    //check if the block is in mock cache if it is not in the cache it will ret
    suite.cache.On("GetCache", "search:Test Blog::").Return(nil, nil)


    blogs, err := suite.repository.SearchBlog("Test Blog", "", nil)
    suite.Require().NoError(err)
    suite.Require().Len(blogs, 1)
    
   
}

func (suite *BlogRepositoryTestSuite) TestFilterBlog() {
    // Initialize a blog
    blog := &domain.Blog{
        Title:         "Test Blog",
        Content:       "Test Content",
        Author:       "Test Author",
        Tags:          []string{"test", "blog"},
        CreatedAt:     time.Now(),
        LastUpdatedAt: time.Now(),
        ViewsCount:    0,
        LikesCount:    0,
        CommentsCount: 0,
    }

    // Insert the blog
    result, err := suite.repository.InsertBlog(blog)
    suite.Require().NoError(err)
    suite.Require().NotNil(result)

    // Filter the blog
    blogs, err := suite.repository.FilterBlog([]string{"test"}, time.Now().Add(-time.Hour), time.Now().Add(time.Hour))
    suite.Require().NoError(err)
    suite.Require().Len(blogs, 1)


    
}

func (suite *BlogRepositoryTestSuite) TestAddView() {
    // Initialize a blog
    blog := &domain.Blog{
        Title:         "Test Blog",
        Content:       "Test Content",
        Author:       "Test Author",
        Tags:          []string{"test", "blog"},
        CreatedAt:     time.Now(),
        LastUpdatedAt: time.Now(),
        ViewsCount:    0,
        LikesCount:    0,
        CommentsCount: 0,
    }

    // Insert the blog
    result, err := suite.repository.InsertBlog(blog)
    suite.Require().NoError(err)
    suite.Require().NotNil(result)

    // Initialize a view
    view := &domain.View{
        BlogID: result.ID,
        User:   "Test User",
    }

    // Add the view
    err = suite.repository.AddView([]*domain.View{view})
    suite.Require().NoError(err)

   

}

func (suite *BlogRepositoryTestSuite) TestAddLike() {
    // Initialize a blog
    blog := &domain.Blog{
        Title:         "Test Blog",
        Content:       "Test Content",
        Author:       "Test Author",
        Tags:          []string{"test", "blog"},
        CreatedAt:     time.Now(),
        LastUpdatedAt: time.Now(),
        ViewsCount:    0,
        LikesCount:    0,
        CommentsCount: 0,
    }

    // Insert the blog
    result, err := suite.repository.InsertBlog(blog)
    suite.Require().NoError(err)
    suite.Require().NotNil(result)

    // Initialize a like
    like := &domain.Like{
        BlogID: result.ID,
        User:   "Test User",
        Like:   true,
    }

    // Add the like
    err = suite.repository.AddLike(like)
    suite.Require().NoError(err)

}

func (suite *BlogRepositoryTestSuite) TestUpdateLike() {
    // Initialize a blog
    blog := &domain.Blog{
        Title:         "Test Blog",
        Content:       "Test Content",
        Author:       "Test Author",
        Tags:          []string{"test", "blog"},
        CreatedAt:     time.Now(),
        LastUpdatedAt: time.Now(),
        ViewsCount:    0,
        LikesCount:    0,
        CommentsCount: 0,
    }

    // Insert the blog
    result, err := suite.repository.InsertBlog(blog)
    suite.Require().NoError(err)
    suite.Require().NotNil(result)

    // Initialize a like
    like := &domain.Like{
        BlogID: result.ID,
        User:   "Test User",
        Like:   true,
    }

    // Add the like
    err = suite.repository.AddLike(like)
    suite.Require().NoError(err)
    
    // Update the like

    like.Like = false
    err = suite.repository.UpdateLike(like)
    suite.Require().NoError(err)

    
}

func (suite *BlogRepositoryTestSuite) TestAddComment() {
    // Initialize a blog
    blog := &domain.Blog{
        Title:         "Test Blog",
        Content:       "Test Content",
        Author:      "Test Author",
        Tags:          []string{"test", "blog"},
        CreatedAt:     time.Now(),
        LastUpdatedAt: time.Now(),
        ViewsCount:    0,
        LikesCount:    0,
        CommentsCount: 0,
    }

    // Insert the blog
    result, err := suite.repository.InsertBlog(blog)
    suite.Require().NoError(err)
    suite.Require().NotNil(result)

    // Initialize a comment
    comment := &domain.Comment{
        BlogID:  result.ID,
        Author: "Test Author",
        Content: "Test Content",   
        Date: time.Now(),
    }

    // Add the comment
    err = suite.repository.AddComment(comment)
    suite.Require().NoError(err)


}

func (suite *BlogRepositoryTestSuite) TestGetBlogsByPopularity() {
    // Initialize blogs
    blogs := []*domain.Blog{
        {
            Title:         "Test Blog 1",
            Content:       "Test Content 1",
            Author:      "Test Author 1",
            Tags:          []string{"test", "blog"},
            CreatedAt:     time.Now(),
            LastUpdatedAt: time.Now(),
            ViewsCount:    10,
            LikesCount:    10,
            CommentsCount: 10,
        },
        {
            Title:         "Test Blog 2",
            Content:       "Test Content 2",
            Author:     "Test Author 2",
            Tags:          []string{"test", "blog"},
            CreatedAt:     time.Now(),
            LastUpdatedAt: time.Now(),
            ViewsCount:    20,
            LikesCount:    20,
            CommentsCount: 20,
        },
        {
            Title:         "Test Blog 3",
            Content:       "Test Content 3",
            Author: "Test Author 3",
            Tags:          []string{"test", "blog"},
            CreatedAt:     time.Now(),
            LastUpdatedAt: time.Now(),
            ViewsCount:    30,
            LikesCount:    30,
            CommentsCount: 30,
        },
    }

    // Insert the blogs
    for _, blog := range blogs {
        _, err := suite.repository.InsertBlog(blog)
        suite.Require().NoError(err)
    }

    // Get the blogs by popularity
    popularBlogs, err := suite.repository.GetBlogsByPopularity(0, 10, false)
    suite.Require().NoError(err)
    suite.Require().Len(popularBlogs, 3)

    
    // Get the blogs by popularity in reverse order
    popularBlogs, err = suite.repository.GetBlogsByPopularity(0, 10, true)
    suite.Require().NoError(err)
    suite.Require().Len(popularBlogs, 3)

    
}

func (suite *BlogRepositoryTestSuite) TestGetBlogsByRecent() {
    // Initialize blogs
    blogs := []*domain.Blog{
        {
            Title:         "Test Blog 1",
            Content:       "Test Content 1",
            Author: "Test Author 1",
            Tags:          []string{"test", "blog"},
            CreatedAt:     time.Now().Add(-time.Hour),
            LastUpdatedAt: time.Now().Add(-time.Hour),
            ViewsCount:    10,
            LikesCount:    10,
            CommentsCount: 10,
        },
        {
            Title:         "Test Blog 2",
            Content:       "Test Content 2",
            Author: "Test Author 2",
            Tags:          []string{"test", "blog"},
            CreatedAt:     time.Now(),
            LastUpdatedAt: time.Now(),
            ViewsCount:    20,
            LikesCount:    20,
            CommentsCount: 20,
        },
        {
            Title:         "Test Blog 3",
            Content:       "Test Content 3",
            Author: "Test Author 3",
            Tags:          []string{"test", "blog"},
            CreatedAt:     time.Now().Add(time.Hour),
            LastUpdatedAt: time.Now().Add(time.Hour),
            ViewsCount:    30,
            LikesCount:    30,
            CommentsCount: 30,
        },
    }

    // Insert the blogs
    for _, blog := range blogs {
        _, err := suite.repository.InsertBlog(blog)
        suite.Require().NoError(err)
    }

    // Get the blogs by recent
    recentBlogs, err := suite.repository.GetBlogsByRecent(0, 10, false)
    suite.Require().NoError(err)
    suite.Require().Len(recentBlogs, 3)

    

    // Get the blogs by recent in reverse order
    recentBlogs, err = suite.repository.GetBlogsByRecent(0, 10, true)
    suite.Require().NoError(err)
    suite.Require().Len(recentBlogs, 3)

   

}

func (suite *BlogRepositoryTestSuite) TestGetLikebyAuthorAndBlogID() {
    // Initialize a blog
    blog := &domain.Blog{
        Title:         "Test Blog",
        Content:       "Test Content",
        Author: "Test Author",
        Tags:          []string{"test", "blog"},
        CreatedAt:     time.Now(),
        LastUpdatedAt: time.Now(),
        ViewsCount:    0,
        LikesCount:    0,
        CommentsCount: 0,
    }

    // Insert the blog
    result, err := suite.repository.InsertBlog(blog)
    suite.Require().NoError(err)
    suite.Require().NotNil(result)

    // Initialize a like
    like := &domain.Like{
        BlogID: result.ID,
        User:   "Test User",
        Like:   true,
    }

    // Add the like
    err = suite.repository.AddLike(like)
    suite.Require().NoError(err)

    // Get the like by author and blog ID
    insertedLike, err := suite.repository.GetLikebyAuthorAndBlogID(result.ID.Hex(), "Test User")
    suite.Require().NoError(err)
    suite.Require().NotNil(insertedLike)

}

func (suite *BlogRepositoryTestSuite) TestGetBlogComments() {
    // Initialize a blog
    blog := &domain.Blog{
        Title:         "Test Blog",
        Content:       "Test Content",
        Author: "Test Author",
        Tags:          []string{"test", "blog"},
        CreatedAt:     time.Now(),
        LastUpdatedAt: time.Now(),
        ViewsCount:    0,
        LikesCount:    0,
        CommentsCount: 0,
    }

    // Insert the blog
    result, err := suite.repository.InsertBlog(blog)
    suite.Require().NoError(err)
    suite.Require().NotNil(result)

    // Initialize a comment
    comment := &domain.Comment{
        BlogID:  result.ID,
        Author: "Test Author",
        Content: "Test Content",
        Date: time.Now(),
    }

    // Add the comment
    err = suite.repository.AddComment(comment)
    suite.Require().NoError(err)

    // Get the comments of the blog
    comments, err := suite.repository.GetBlogComments(result.ID.Hex())
    suite.Require().NoError(err)
    suite.Require().Len(comments, 1)

    
}

func (suite *BlogRepositoryTestSuite) TestGetBlogLikes() {
    // Initialize a blog
    blog := &domain.Blog{
        Title:         "Test Blog",
        Content:       "Test Content",
        Author: "Test Author",
        Tags:          []string{"test", "blog"},
        CreatedAt:     time.Now(),
        LastUpdatedAt: time.Now(),
        ViewsCount:    0,
        LikesCount:    0,
        CommentsCount: 0,
    }

    // Insert the blog
    result, err := suite.repository.InsertBlog(blog)
    suite.Require().NoError(err)
    suite.Require().NotNil(result)

    // Initialize a like
    like := &domain.Like{
        BlogID: result.ID,
        User:   "Test User",
        Like:   true,
    }

    // Add the like
    err = suite.repository.AddLike(like)
    suite.Require().NoError(err)

    // Get the likes of the blog
    likes, err := suite.repository.GetBlogLikes(result.ID.Hex())
    suite.Require().NoError(err)
    suite.Require().Len(likes, 1)

   

}

func (suite *BlogRepositoryTestSuite) TestIncrmentBlogViews() {
    // Initialize a blog
    blog := &domain.Blog{
        Title:         "Test Blog",
        Content:       "Test Content",
        Author: "Test Author",
        Tags:          []string{"test", "blog"},
        CreatedAt:     time.Now(),
        LastUpdatedAt: time.Now(),
        ViewsCount:    0,
        LikesCount:    0,
        CommentsCount: 0,
    }

    // Insert the blog
    result, err := suite.repository.InsertBlog(blog)
    suite.Require().NoError(err)
    suite.Require().NotNil(result)

    // Increment the views of the blog
    err = suite.repository.IncrmentBlogViews(result.ID.Hex())
    suite.Require().NoError(err)

    
}

func (suite *BlogRepositoryTestSuite) TestIncrmentBlogLikes() {
    // Initialize a blog
    blog := &domain.Blog{
        Title:         "Test Blog",
        Content:       "Test Content",
        Author: "Test Author",
        Tags:          []string{"test", "blog"},
        CreatedAt:     time.Now(),
        LastUpdatedAt: time.Now(),
        ViewsCount:    0,
        LikesCount:    0,
        CommentsCount: 0,
    }

    // Insert the blog
    result, err := suite.repository.InsertBlog(blog)
    suite.Require().NoError(err)
    suite.Require().NotNil(result)

    // Increment the likes of the blog
    err = suite.repository.IncrmentBlogLikes(result.ID.Hex())
    suite.Require().NoError(err)

}

func (suite *BlogRepositoryTestSuite) TestIncrmentBlogComments() {
    // Initialize a blog
    blog := &domain.Blog{
        Title:         "Test Blog",
        Content:       "Test Content",
        Author: "Test Author",
        Tags:          []string{"test", "blog"},
        CreatedAt:     time.Now(),
        LastUpdatedAt: time.Now(),
        ViewsCount:    0,
        LikesCount:    0,
        CommentsCount: 0,
    }

    // Insert the blog
    result, err := suite.repository.InsertBlog(blog)
    suite.Require().NoError(err)
    suite.Require().NotNil(result)

    // Increment the comments of the blog
    err = suite.repository.IncrmentBlogComments(result.ID.Hex())
    suite.Require().NoError(err)

    
}

func (suite *BlogRepositoryTestSuite) TestDecrementBlogLikes() {
    // Initialize a blog
    blog := &domain.Blog{
        Title:         "Test Blog",
        Content:       "Test Content",
        Author: "Test Author",
        Tags:          []string{"test", "blog"},
        CreatedAt:     time.Now(),
        LastUpdatedAt: time.Now(),
        ViewsCount:    0,
        LikesCount:    1,
        CommentsCount: 0,
    }

    // Insert the blog
    result, err := suite.repository.InsertBlog(blog)
    suite.Require().NoError(err)
    suite.Require().NotNil(result)


    // Decrement the likes of the blog
    err = suite.repository.DecrementBlogLikes(result.ID.Hex())
    suite.Require().NoError(err)

    
}

//TearDownTest is called after each test in the suite
func (suite *BlogRepositoryTestSuite) TearDownTest() {
    // Drop test collections
    err := suite.blogCollection.Drop(context.Background())
    if err != nil {
        suite.T().Fatalf("Failed to drop blogs collection: %v", err)
    }
    err = suite.viewCollection.Drop(context.Background())
    if err != nil {
        suite.T().Fatalf("Failed to drop views collection: %v", err)
    }
    err = suite.likeCollection.Drop(context.Background())
    if err != nil {
        suite.T().Fatalf("Failed to drop likes collection: %v", err)
    }
    err = suite.commentCollection.Drop(context.Background())
    if err != nil {
        suite.T().Fatalf("Failed to drop comments collection: %v", err)
    }
}

//TeardownSuite is called after all tests in the suite have run
func (suite *BlogRepositoryTestSuite) TearDownSuite() {
   
    // Disconnect from MongoDB
    err := suite.client.Disconnect(context.Background())
    if err != nil {
        suite.T().Fatalf("Failed to disconnect from MongoDB: %v", err)
    }
}

func TestBlogRepositoryTestSuite(t *testing.T) {
    suite.Run(t, new(BlogRepositoryTestSuite))
}











    



    