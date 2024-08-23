package middleware

import (
	"net/http"
	"strings"

	"github.com/RealEskalate/blogpost/config"
	"github.com/RealEskalate/blogpost/database"
	"github.com/RealEskalate/blogpost/domain"
	"github.com/RealEskalate/blogpost/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func RoleBasedAuth(protected bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		//get token with claims
		auth := c.GetHeader("Authorization")

		var claims = domain.UserClaims{}
		authSplit := strings.Split(auth, " ")
		_, err := jwt.ParseWithClaims(authSplit[1], &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("accessblahblah"), nil
		})

		if err != nil {
			c.IndentedJSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		usr, _ := c.Get("user")
		user, _ := usr.(domain.ResponseUser)

		if user.Is_Admin{
			c.Set("filter", bson.M{})
		} else {
			if protected {
				c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "you need to be an admin"})
				c.Abort()
				return
			}

			path := c.Request.URL.Path
			idx := c.Param("id")
        	if strings.Contains(path, "user") && idx != "" && idx != user.ID {
				c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
				c.Abort()
				return
			}else if strings.Contains(path, "blog") && idx != "" && c.Request.Method != "GET"{
				SC := config.ServerConnection{}
				SC.Connect_could()
				coll := &database.MongoCollection{Collection: SC.Client.Database("BlogPost").Collection("Blogs")}
				BR := repository.NewBlogRepository(coll)
				blog, _ := BR.GetOneBlogDocument(idx)
				if blog.Owner.ID != user.ID{
					c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
					c.Abort()
					return
				}
			
			}
			c.Set("filter", bson.M{"user._id": claims.ID})
		}

		c.Next()
	}

}