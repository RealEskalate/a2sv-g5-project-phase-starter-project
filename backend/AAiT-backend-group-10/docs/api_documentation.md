# Auth API Documentation

## Overview
This document outlines the endpoints for user authentication, including registration, login, token refresh, and password management.

### Base URL
``` localhost:8080 ```

## Endpoints

### 1. Register User
**Endpoint:** `POST /register`

**Description:** 
The POST request to /register endpoint is used to register a new user. The request should include the fullname, email, and password in the raw request body.

**Request Body:**
```json
{
    "fullname": "samrawit dawit",
    "email": "samri@gmail.com",
    "password": "123"
}
```
**Response:**
``` json
{
    "id": "264b2261-2054-4c99-9a69-b1906f9a0498",
    "full_name": "samrawit dawit",
    "email": "samri@gmail.com",
    "bio": "",
    "image_url": ""
}
```
**Example cURL:**
```bash
curl --location 'localhost:8080/register' \
--data-raw '{
    "fullname": "samrawit dawit",
    "email": "samri@gmail.com",
    "password": "123"
}'
```

### 2. Login
**Endpoint:** `POST /login`

**Description:** 
The POST /login endpoint is used to authenticate users and obtain access and refresh tokens. The request should include the user's email and password in the request body.

**Request Body:**
```json
{
    "email": "samri@gmail.com",
    "password": "123"
}

```
**Response:**
``` json
{
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}

```
**Example cURL:**
```bash
curl --location 'localhost:8080/login' \
--data-raw '{
    "email": "samri@gmail.com",
    "password": "123"
}'

```

### 3. Refresh Token
**Endpoint:** `POST /refresh-token`

**Description:** 

This endpoint is used to refresh an access token using a refresh token. The HTTP POST request should be made to localhost:8080/refresh-token with a payload containing the refresh token in the raw request body.

**Request Body:**
```json
{
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}


```
**Response:**
``` json
{
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}

```
**Example cURL:**
```bash
curl --location 'localhost:8080/refresh-token' \
--data '{
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}'

```
### 4. Forgot Password
**Endpoint:** `POST /forgot-password`

**Description:** 
The POST /forgot-password endpoint is used to initiate the process of resetting a user's password. It expects a JSON payload with the user's email address in the request body.

**Request Body:**
```json
{
    "email": "abe@gmail.com"
}



```
**Example cURL:**
```bash
curl --location 'localhost:8080/forgot-password' \
--data-raw '{
    "email": "abe@gmail.com"
}'
```

### 3. Reset Password
**Endpoint:** `POST /reset-password`

**Description:** 
This endpoint is used to reset a user's password.

**Request Body:**
```json
{
 "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
 "new_password": "456"
}


```
**Example cURL:**
```bash
curl --location 'localhost:8080/reset-password' \
--data-raw '{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "new_password": "456"
}'
```

# Blog-API Documentation

## POST /blogs

### CreateBlog

**URL**: `localhost:8080/blogs`  
**Method**: `POST`  
**Description**: Add a new blog post.

**Request Body**:

- `title` (string, required): The title of the blog post.
- `content` (string, required): The content of the blog post.
- `author` (string, required): The author of the blog post.
- `tags` (array of strings, required): The tags associated with the blog post.

**Response**:

- Status Code: `201 Created`
- Body (JSON):

```json
{
    "id": "unique-id",
    "title": "title",
    "content": "content",
    "author": "author-id",
    "tags": ["tag1", "tag2"],
    "createdAt": "date-time",
    "updatedAt": "date-time",
    "viewCount": 0
}
```
**Request Headers**:
- ` Authorization: Bearer <token> `

**Example request**:
```bash
curl --location 'localhost:8080/blogs' \
--header 'Authorization: Bearer <token>' \
--data '{
    "title": "tenth blog",
    "content": "blog 10",
    "author": "author-id",
    "tags": ["Rock", "EDM"]
}'
```
**Example response**:
```json
{
    "id": "91a392c8-c8f3-4489-a149-6151f7c92ec4",
    "title": "tenth blog",
    "content": "blog 10",
    "author": "author-id",
    "tags": ["Rock", "EDM"],
    "createdAt": "2024-08-18T12:42:18.575Z",
    "updatedAt": "2024-08-18T12:42:18.575Z",
    "viewCount": 0
}
```

### GetAllBlogs

**URL**: `localhost:8080/blogs`  
**Method**: `GET`  
**Description**: Retrieve a list of all blog posts.

**Response**:

- Status Code: `200 Ok`
- Body (JSON):

```json
[
    {
        "id": "id",
        "title": "title",
        "content": "content",
        "author": "author-id",
        "tags": ["tag1", "tag2"],
        "createdAt": "date-time",
        "updatedAt": "date-time",
        "viewCount": 0
    }
]

```
**Request Headers**:
- ` Authorization: Bearer <token> `

**Example request**:
```bash
curl --location 'localhost:8080/blogs' \
--header 'Authorization: Bearer <token>'

```
**Example response**:
```json
[
    {
        "id": "c7bed03d-76bd-438e-9c19-e28544454ab5",
        "title": "first blog",
        "content": "blog 1",
        "author": "author-id",
        "tags": ["Sport", "Entertainment"],
        "createdAt": "2024-08-18T12:27:46.231Z",
        "updatedAt": "2024-08-18T12:27:46.231Z",
        "viewCount": 1000
    },
    {
        "id": "2f5b54c1-9827-4b5a-98e5-778d76842564",
        "title": "second blog",
        "content": "blog 2",
        "author": "author-id",
        "tags": ["Sport", "News"],
        "createdAt": "2024-08-18T12:34:13.07Z",
        "updatedAt": "2024-08-18T12:34:13.07Z",
        "viewCount": 333
    }
]

```

### GetByID

**URL**: `localhost:8080/blogs:id`  
**Method**: `GET`  
**Description**: Retrieve a specific blog post by its ID.

**Path Parameter**:
- `id` (string): The unique identifier of the blog post.

**Response**:

- Status Code: `200 Ok`
- Body (JSON):

```json
{
    "id": "id",
    "title": "title",
    "content": "content",
    "author": "author-id",
    "tags": ["tag1", "tag2"],
    "createdAt": "date-time",
    "updatedAt": "date-time",
    "viewCount": 0
}


```
**Request Headers**:
- ` Authorization: Bearer <token> `

**Example request**:
```bash
curl --location 'localhost:8080/blogs/<id>' \
--header 'Authorization: Bearer <token>'

```
**Example response**:
```json
{
    "id": "c7bed03d-76bd-438e-9c19-e28544454ab5",
    "title": "first blog",
    "content": "blog 1",
    "author": "author-id",
    "tags": ["Sport", "Entertainment"],
    "createdAt": "2024-08-18T12:27:46.231Z",
    "updatedAt": "2024-08-18T12:27:46.231Z",
    "viewCount": 1000
}
```


### UpdateBlog

**URL**: `localhost:8080/blogs:id`  
**Method**: `PUT`  
**Description**:  Update a specific blog by its ID.

**Request Body**:

- `title` (string): The updated title of the blog.
- `content` (string): The updated content of the blog.
- `author` (string): The updated author of the blog.
- `tags` (array of strings) The updated tags associated with the blog.post.

**Response**:

- Status Code: `200 OK`
- Body (JSON):

```json
{
    "message": "Blog updated successfully"
}

```
**Request Headers**:
- ` Authorization: Bearer <token> `

**Path Parameter**:
`id` (string): The unique identifier of the blog post.

**Example request**:
```bash
curl --location --request PUT 'localhost:8080/blogs/<id>' \
--header 'Authorization: Bearer <token>' \
--data '{
    "title": "first blog update",
    "content": "blog 1",
    "author": "author-id",
    "tags": ["Sport", "Entertainment"]
}'

```
**Example response**:
```json
{
    "message": "Blog updated successfully"
}

```
### AddView

**URL**: `localhost:8080/blogs/:id/view`  
**Method**: `PATCH`  
**Description**:  Update the view count of a specific blog post.

**Request Body**:

- `id` (string, required): The unique identifier of the blog post.

**Response**:

- Status Code: `200 OK`
- Body (JSON):

```json
{
    "message": "View added successfully"
}

```
**Request Headers**:
- ` Authorization: Bearer <token> `

**Path Parameter**:
`id` (string): The unique identifier of the blog post.

**Example request**:
```bash
curl --location --request PATCH 'localhost:8080/blogs/<id>/view' \
--header 'Authorization: Bearer <token>'


```
**Example response**:
```json
{
    "message": "View added successfully"
}

```
### DeleteBlog

**URL**: `localhost:8080/blogs/:id`  
**Method**: `DELETE`  
**Description**:  Delete a specific blog post by its ID.

**Response**:

- Status Code: `200 OK`
- Body (JSON):

```json
{
    "message": "Blog deleted successfully"
}

```
**Request Headers**:
- ` Authorization: Bearer <token> `

**Path Parameter**:
`id` (string): The unique identifier of the blog post.

**Example request**:
```bash
curl --location --request DELETE 'localhost:8080/blogs/<id>' \
--header 'Authorization: Bearer <token>'

```
**Example response**:
```json
{
    "message": "Blog deleted successfully"
}

```
### SearchBlog

**URL**: `localhost:8080/blogs/search`  
**Method**: `GET`  
**Description**:  Retrieve a list of blogs based on search criteria such as page number, limit, title, and author.

**Response**:

- Status Code: `200 OK`
- Body (JSON):

```json
{
    "message": "Blog deleted successfully"
}

```
**Request Headers**:
- ` Authorization: Bearer <token> `

**Path Parameter**:
`id` (string): The unique identifier of the blog post.

**Example request**:
```bash
curl --location --request DELETE 'localhost:8080/blogs/<id>' \
--header 'Authorization: Bearer <token>'

```
**Example response**:
```json
{
    "message": "Blog deleted successfully"
}

```

### DeleteBlog

**URL**: `localhost:8080/blogs/:id`  
**Method**: `DELETE`  
**Description**:  Delete a specific blog post by its ID.

**Query Parameters**:

- `page` (integer): The page number for paginated results.
- `limit` (integer):The maximum number of blogs to be returned per page.
- `title` (string): The title to be used as a search filter.
- `author` (string): The author's unique identifier to filter blogs by author.

**Response**:

- Status Code: `200 OK`
- Body (JSON):

```json
{
    "blogs": [
        {
            "id": "id",
            "title": "title",
            "content": "",
            "author": "",
            "tags": [""],
            "createdAt": "",
            "updatedAt": "",
            "viewCount": 0
        }
    ],
    "currentPage": 0,
    "totalCount": 0,
    "totalPages": 0
}

```
**Request Headers**:
- ` Authorization: Bearer <token> `

**Query Parameter**:
`author` : seifu.

**Example request**:
```bash
curl --location 'localhost:8080/blogs/search?page=1&limit=2&title=nth&author=a5a06932-deec-4d73-3333-45191b69f6ff'

```
**Example response**:
```json
{
    "blogs": [
        {
            "id": "148c4b8e-c567-42b1-a36a-8166ea9359a5",
            "title": "seventh blog",
            "content": "blog 7",
            "author": "a5a06932-deec-4d73-3333-45191b69f6ff",
            "tags": [
                "Tech"
            ],
            "createdAt": "2024-08-18T12:35:34.982Z",
            "updatedAt": "2024-08-18T12:35:34.982Z",
            "viewCount": 11
        }
    ],
    "currentPage": 1,
    "totalCount": 1,
    "totalPages": 1
}

```

### GenerateContent
**URL**: `localhost:8080/blogs/generate`  
**Method**: `POST`  
**Description**:  Generates .

**Request Headers**:
- ` Authorization: Bearer <token> `

**Request body**:
```json
{
    "topic": "Tech Trends",
    "keywords": ["AI", "Blockchain", "Quantum Computing"]
}
```
### SuggestImprovement
**URL**: `localhost:8080/blogs/suggest`  
**Method**: `POST`  
**Description**:  Suggests.. .

**Request Headers**:
- ` Authorization: Bearer <token> `

**Request body**:
```json
json
{
    "content": "blog about my day, I went to school then I ate lunch then get back to school again then went home"
}
```


# Comment-API Documentation

## POST /comment

### AddComment

**URL**: `localhost:8080/comment`  
**Method**: `POST`  
**Description**: Submit a comment on a specific blog.

**Request Body**:

- `blog_id` (string, required): The ID of the blog on which the comment is being submitted.
- `comment` (string, required): The content of the comment being submitted.

**Response**:

- Status Code: `201 Created`

**Request Headers**:
- `Authorization: Bearer <token>`
- `Content-Type: application/json`

**Example request**:
```bash
curl --location 'localhost:8080/comment' \
--header 'Authorization: Bearer <token>' \
--header 'Content-Type: application/json' \
--data '{
    "blog_id": "22e4f98c-7eb0-401e-8007-9315b5b9e743",
    "comment": "good one"
}'
```
### GetBlogcomments

**URL**: `localhost:8080/comment/:blog_id`  
**Method**: `GET`  
**Description**: Retrieve the comments for a specific blog post.

**Path Variables**:

- `blog_id` (string, required): The ID of the blog on which the comment is being submitted.


**Response**:

- Status Code: `200 OK` 
- Body (JSON):
```json
{
    "comments": [
        {
            "id": "string",
            "text": "string",
            "author": "string",
            "created_at": "2024-08-18T12:35:34.982Z"
        }
    ]
}

```
**Request Headers**:
- `Authorization: Bearer <token>`

### EditComment

**URL**: `localhost:8080/comment/:id`  
**Method**: `PUT`  
**Description**: Update a specific comment associated with a blog.

**Path Variables**:

- `id` (string, required): The ID of the comment to be updated.

**Request Body**:

- `blog_id` (string): The ID of the blog associated with the comment.
- `comment` (string): The updated content of the comment.

**Response**:

- Status Code: `200 OK`
- Body (JSON): Represents the structure of the updated comment object.

**Request Headers**:
- `Authorization: Bearer <token>`
- `Content-Type: application/json`

**Example request**:
```bash
curl --location --request PUT 'localhost:8080/comment/72b44200-8460-4b57-b4d6-190800d6d213' \
--header 'Authorization: Bearer <token>' \
--header 'Content-Type: application/json' \
--data '{
    "blog_id": "22e4f98c-7eb0-401e-8007-9315b5b9e743",
    "comment": "that'\''s a good one"
}'

```

### DeleteComment

**URL**: `localhost:8080/comment/:id`  
**Method**: `Delete`  
**Description**:  Delete a specific comment by its ID.

**Path Variables**:

- `id` (string, required): The ID of the comment to be updated.


**Response**:

- Status Code: `200 OK`
- Status Code: `404 Not Found`
- Status Code: `500 Internal Server Error`

**Request Headers**:
- `Authorization: Bearer <token>`
- `Content-Type: application/json`

# Like-API Documentation

## PUT /like

### ReactBlog

**URL**: `localhost:8080/like`  
**Method**: `PUT`  
**Description**: Update the like status for a specific blog.

**Request Body**:

- `blog_id` (string, required): The ID of the blog for which the like status is being updated.
- `is_like` (boolean, required): Indicates whether the blog is being liked (`true`) or unliked (`false`).

**Response**:

- Status Code: `200 OK` (example)
- Body (JSON):

```json
{
    "message": "Like added successfully"
}
```

**Request Headers**:
- `Authorization: Bearer <token>`
- `Content-Type: application/json`

**Example request**:
```bash
curl --location --request PUT 'localhost:8080/like' \
--header 'Authorization: Bearer <token>' \
--header 'Content-Type: application/json' \
--data '{
    "blog_id": "22e4f98c-7eb0-401e-8007-9315b5b9e743",
    "is_like": false
}'

```

## DELETE /like

### Unlike

**URL**: `localhost:8080/like`  
**Method**: `DELETE`  
**Description**: Remove a "like" from a specific blog post.

**Request Body**:

- `blog_id` (string, required): The ID of the blog for which the like status is being updated.


**Response**:

- Status Code: `200 OK` (example)
- Body (JSON): 
```json
{
    "message": "Like deleted successfully"
}
```

**Request Headers**:
- `Authorization: Bearer <token>`
- `Content-Type: application/json`

**Example request**:
```bash
curl --location --request DELETE 'localhost:8080/like' \
--header 'Authorization: Bearer <token>' \
--header 'Content-Type: application/json' \
--data '{
    "blog_id": "22e4f98c-7eb0-401e-8007-9315b5b9e743"
}'
```

# User-API Documentation

## PATCH /users/promote

### PromoteUser

**URL**: `localhost:8080/users/promote`  
**Method**: `PATCH`  
**Description**: Promote or demote a user to/from an admin role.

**Request Body**:

- `id` (string, required): The ID of the user to be promoted or demoted.
- `is_promote` (boolean, required): `true` to promote the user to admin, `false` to demote.

**Response**:

- Status Code: `200 OK`
- Body (JSON): 

```json
```

**Request Headers**:
- `Authorization: Bearer <token>`
- `Content-Type: application/json`

**Example request**:
```bash
curl --location --request PATCH 'localhost:8080/users/promote' \
--header 'Authorization: Bearer <token>' \
--header 'Content-Type: application/json' \
--data '{
    "id": "85400eef-85e8-479a-8408-c7f4355670c9",
    "is_promote": true
}'
```
## PUT /users

### UpdateProfile

**URL**: `localhost:8080/users/:id`  
**Method**: `PUT`  
**Description**: Update the profile information of a specific user.

**Request Body**:

- `bio` (string, required): The updated bio of the user.

**Response**:

- Status Code: `200 OK`
- Body (JSON): 

```json
{
  "type": "object",
  "properties": {
    "status": {
      "type": "string"
    },
    "message": {
      "type": "string"
    }
  }
}
```

**Request Headers**:
- `Authorization: Bearer <token>`
- `Content-Type: application/json`

**Example request**:
```bash
curl --location --request PUT 'localhost:8080/users/85400eef-85e8-479a-8408-c7f4355670c9' \
--header 'Authorization: Bearer <token>' \
--header 'Content-Type: application/json' \
--data '{
    "bio": "HUMAN"
}'
```

## POST /upload-image

### Upload Profile Picture

**URL**: `localhost:8080/upload-image`  
**Method**: `POST`  
**Description**:  Upload a profile picture for a user.

**Request Body**:

- `profile_pic` (file, required): The image file to be uploaded.

**Response**:

- Status Code: `200 OK`
- Body (JSON): 

```json
{
    "message": "Profile picture uploaded successfully",
    "profile_pic_url": "https://res.cloudinary.com/db7bvqukt/image/upload/v1724273743/profile_pics/264b2261-2054-4c99-9a69-b1906f9a0498-Screenshot%20from%202024-08-16%2009-45-18.png.png"
}

```

**Request Headers**:
- `Authorization: Bearer <token>`
- `Content-Type: application/json`

**Example request**:
```bash
curl --location --request POST 'localhost:8080/upload-image' \
--header 'Authorization: Bearer <token>' \
--form 'profile_pic=@"/home/yordanos/Pictures/Screenshots/Screenshot from 2024-08-16 09-45-18.png"'

```
