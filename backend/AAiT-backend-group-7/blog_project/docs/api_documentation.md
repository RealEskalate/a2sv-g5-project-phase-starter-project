# Blog Platform API Documentation

This documentation outlines the API endpoints for managing blogs and users on the Blog Platform. The platform supports various operations such as creating, retrieving, updating, and deleting blogs, as well as advanced functionalities like commenting, liking, disliking, and AI-generated content recommendations.

## Base URL

```
http://localhost:8080
```

## Authentication

The API uses JWT for authentication. Most endpoints require a valid JWT token included in the `Authorization` header of the request.

---

## Blogs

### 1. Get All Blogs

**Endpoint:**  
`GET /blogs`

**Query Parameters:**

- `sort` (optional): Defines the sort order of the blogs. Default is `DESC` for descending order. Acceptable values: `ASC` (ascending), `DESC` (descending).
- `page` (optional): Specifies the page number for pagination. Default is `1`.
- `limit` (optional): Determines the number of blogs returned per page. Default is `10`.

**Description:**  
Retrieves a list of blogs with options for sorting and pagination. The response includes metadata such as total pages and total number of blogs.

**Response:**

- `200 OK`: Successfully retrieved the list of blogs.
- `400 Bad Request`: Invalid query parameters provided.

---

### 2. Create a Blog

**Endpoint:**  
`POST /blogs`

**Description:**  
Creates a new blog entry. The user must be authenticated to perform this action.

**Request Body:**

- `title`: The title of the blog (required).
- `content`: The content of the blog (required).
- `author_id`: The unique ID of the author creating the blog (required).

**Response:**

- `200 OK`: The blog was successfully created.
- `400 Bad Request`: The request body contains invalid or missing fields.

---

### 3. Update a Blog

**Endpoint:**  
`PUT /blogs/:id`

**Path Parameters:**

- `id`: The unique ID of the blog to be updated.

**Description:**  
Updates an existing blog. Only the fields provided in the request body will be updated. The user must be the author of the blog or an admin.

**Request Body:**

- Fields to be updated, such as `title` and `content`.

**Response:**

- `200 OK`: The blog was successfully updated.
- `400 Bad Request`: Invalid ID or request body provided.

---

### 4. Delete a Blog

**Endpoint:**  
`DELETE /blogs/:id`

**Path Parameters:**

- `id`: The unique ID of the blog to be deleted.

**Description:**  
Deletes an existing blog. The user must be the author of the blog or have admin privileges.

**Response:**

- `200 OK`: The blog was successfully deleted.
- `400 Bad Request`: Invalid ID provided.

---

### 5. Add a Comment to a Blog

**Endpoint:**  
`POST /blogs/:blog_id/:author_id/comment`

**Path Parameters:**

- `blog_id`: The unique ID of the blog where the comment will be added.
- `author_id`: The unique ID of the user making the comment.

**Request Body:**

- `content`: The content of the comment (required).

**Description:**  
Adds a comment to a specified blog. The user must be authenticated to perform this action.

**Response:**

- `200 OK`: The comment was successfully added to the blog.
- `400 Bad Request`: Invalid request body or parameters provided.

---

### 6. Like a Blog

**Endpoint:**  
`POST /blogs/:blog_id/:author_id/like`

**Path Parameters:**

- `blog_id`: The unique ID of the blog to be liked.
- `author_id`: The unique ID of the user liking the blog.

**Description:**  
Registers a like for the specified blog. The system ensures that the user has not already liked the blog to prevent duplicate likes.

**Response:**

- `200 OK`: The like was successfully registered.
- `400 Bad Request`: Invalid parameters provided.

---

### 7. Dislike a Blog

**Endpoint:**  
`POST /blogs/:blog_id/:author_id/dislike`

**Path Parameters:**

- `blog_id`: The unique ID of the blog to be disliked.
- `author_id`: The unique ID of the user disliking the blog.

**Description:**  
Registers a dislike for the specified blog. The system ensures that the user has not already disliked the blog to prevent duplicate dislikes.

**Response:**

- `200 OK`: The dislike was successfully registered.
- `400 Bad Request`: Invalid parameters provided.

---

### 8. Search Blogs

**Endpoint:**  
`POST /blogs/search`

**Query Parameters:**

- `author` (optional): Filters blogs by author name.
- `tags` (optional): Filters blogs by tags. Multiple tags can be provided.
- `title` (optional): Filters blogs by title.

**Description:**  
Searches for blogs based on the provided criteria, such as author, tags, or title. The results are filtered and returned to the user.

**Response:**

- `200 OK`: Successfully retrieved the list of blogs that match the search criteria.
- `400 Bad Request`: Invalid query parameters provided.

---

### 9. AI Blog Content Recommendation

**Endpoint:**  
`POST /blogs/GenerateContent`

**Request Body:**

- `content`: The text content that the AI will analyze to provide recommendations (required).

**Description:**  
Generates content recommendations or suggestions based on the provided text using AI capabilities. This feature assists users in enhancing or expanding their blog content.

**Response:**

- `200 OK`: Successfully generated AI content recommendations.
- `400 Bad Request`: Invalid request body provided.

---

## Users

### 1. Create User

**Endpoint:**  
`POST /users/`

**Description:**  
Creates a new user account on the platform.

**Request Body:**

- `username` (string): The desired username for the new user (required).
- `password` (string): The password for the new user (required).
- `email` (string): The email address of the new user (required).
- `role` (string): The role of the new user (e.g., `user`, `admin`) (required).
- `bio` (string): A brief biography for the new user (optional).
- `phone` (string): The phone number of the new user (optional).
- `profile_pic` (file): The profile picture of the new user (optional).

**Response:**

- `200 OK`: Successfully created the user. Returns the user object.
- `400 Bad Request`: The request payload is invalid or there was an error uploading the profile picture.

---

### 2. User Login

**Endpoint:**  
`POST /users/login`

**Description:**  
Authenticates a user and returns access and refresh tokens.

**Request Body:**

- `username` (string): The username of the user (required).
- `password` (string): The password of the user (required).

**Response:**

- `200 OK`: Successfully authenticated. Returns the access and refresh tokens.
- `400 Bad Request`: The username or password provided is incorrect.

---

### 3. Forgot Password

**Endpoint:**  
`POST /users/forget-password/:email`

**Description:**  
Sends a password reset link to the specified email address.

**Path Parameters:**

- `email` (string): The email address associated with the user's account (required).

**Response:**

- `200 OK`: The password reset link has been sent to the specified email.
- `400 Bad Request`: There was an error processing the request, such as an invalid email address.

---

### 4. Reset Password

**Endpoint:**  
`POST /users/reset-password/:username/:password`

**Description:**  
Resets the password for a user identified by their username.

**Path Parameters:**

- `username` (string): The username of the user (required).
- `password` (string): The new password to be set for the user (required).

**Response:**

- `200 OK`: The password has been successfully reset.
- `400 Bad Request`: There was an error processing the request, such as an invalid username or password.

---

### 5. Logout

**Endpoint:**  
`POST /users/logout`

**Description:**  
Logs out the user by invalidating the provided access token.

**Request Body:**

- `token` (string): The access token to be invalidated (required).

**Response:**

- `200 OK`: Successfully logged out.
- `400 Bad Request`: There was an error processing the request, such as an invalid token.

---

### 6. Get Users

**Endpoint:**  
`GET /users/`

**Description:**  
Retrieves a list of all users on the platform.

**Response:**

- `200 OK`: Successfully retrieved the list of users.
- `400 Bad Request`: There was an error processing the request.

---

### 7. Get User

**Endpoint:**  
`GET /users/:id`

**Description:**  
Retrieves details for a specific user identified by their ID.

**Path Parameters:**

- `id` (int): The unique ID of the user (required).

**Response:**

- `200 OK`: Successfully retrieved the user object.
- `400 Bad Request`: There was an error processing the request, such as an invalid user ID.

---

### 8. Update User

**Endpoint:**  
`PUT /users/:id`

**Description:**  
Updates details for a specific user identified by their ID.

**Path Parameters:**

- `id` (int): The unique ID of the user to be updated (required).

**Request Body:**

- JSON object representing the fields to update, such as `username`, `email`, `bio`, etc.

**Response:**

- `200 OK`: Successfully updated the user object.
- `400 Bad Request`: There was an error processing the request, such as an invalid ID or invalid fields.

---

### 9. Delete User

**Endpoint:**  
`DELETE /users/:id`

**Description:**  
Deletes a user account identified by their ID.

**Path Parameters:**

- `id` (int): The unique ID of the user to be deleted (required).

**Response:**

- `200 OK`: Successfully deleted the user.
- `400 Bad Request`: There was an error processing the request, such as an invalid user ID.

---

### 10. Promote User

**Endpoint:**  
`POST /users/promote/:id`

**Description:**  
Promotes a user to a higher role, such as from `user` to `admin`.

**Path Parameters:**

- `id` (int): The unique ID of the user to be promoted (required).

**Response:**

- `200 OK`: Successfully promoted the user.
- `400 Bad Request`: There was an error processing the request, such as an invalid user ID.

---

### 11. Demote User

**Endpoint:**  
`POST /users/demote/:id`

**Description:**  
Demotes a user to a lower role.

**Path Parameters:**

- `id` (int): The unique ID of the user to be demoted (required).

**Response:**

- `200 OK`: Successfully demoted the user.
- `400 Bad Request`: There was an error processing the request, such as an invalid user ID.

---

### 12. Refresh Token

**Endpoint:**  
`POST /users/refresh-token`

**Description:**  
Generates a new access token using a refresh token.

**Request Body:**

- `refresh_token` (string): The refresh token used to generate a new access token (required).

**Response:**

- `200 OK`: Successfully generated a new access token.
- `400 Bad Request`: There was an error processing the request, such as an invalid or expired refresh token.
