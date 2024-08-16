# Blog Starter Project


## Entities 

### User

* id
* email - unique, required, valid
* password - strength, required
* fullName - required
* bio
* imageUrl
* isAdmin - boolean, default - false
* refreshToken


### Blog

* id
* Title - required
* Content - required
* Author -  ref(userId)
* Tags - required 
* CreatedAt
* UpdatedAt
* ViewCount


### Like

* id
* BlogId - ref(blogId)
* UserId - ref(userId)
* isLike - true if like, false if dislike


### Comment

* id
* BlogId - ref(blogId)
* UserId - ref(userId)
* Comment - required


## routes

### Auth

* /register - registers and creates new user by gathering required fields
* /login
* /logout
* /identify - to identify email when forgotten password or other identification by sending verification emails
* /verify - accepting and verifying verification emails and granting access or changing password


### User

* /user/update/:id - update profile
* /user/promote/?id?makeAdmin  - promote or demote user with id based on makeAdmin bool


### Blog

* /blogs - get all blogs - also filter based on query parameters if there is one 
* /blogs/:id - get one blog
* /blogs - add blog
* /blogs/:id - update blog
* /blogs/:id - delete blog


### Like

* /like?isLike - like or dislike based on isLike value
* /unlike - delete the like


### Comment

* /comments/:blogid - get one blogs comment
* /comments - add comment
* /comments - update comment