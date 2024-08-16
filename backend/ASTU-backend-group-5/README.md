# API Endpoints

## Auth Endpoints

### 1. User Registration
- **POST** `/api/v1/auth/register`
  - Registers a new user with email, password, and profile details.
###### Request:
```json
{
  "username": "john_doe",
  "email": "john@example.com",
  "password": "StrongPass123",
  
}
```
###### Response:
```
{
  "message": "User registered successfully. Please check your email to activate your account."
}
```


### 2. Login
- **POST** `/api/v1/auth/login`
  - Authenticates the user with username/email and password, returns JWT tokens.
###### Request:
```json
{
  "email": "john@example.com",
  "password": "StrongPass123"
}
  ```
###### Response:

```json
{
  "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "refreshToken": "dGVzdF9yZWZyZXNoX3Rva2Vu..."
}


```

### 3. Token Refresh
- **POST** `/api/v1/auth/token/refresh`
  - Refreshes the access token using the refresh token.

### 4. Forgot Password
- **POST** `/api/v1/auth/password-reset/request`
  - Sends a password reset link via email.
###### Request:
```json
{
  "email": "john@example.com"
}

```
###### Response:
```json

{
  "message": "Password reset link has been sent to your email."
}

```

### 5. Reset Password
- **POST** `/api/v1/auth/password/reset/confirm?token="huifehfiehfjiehfeimkcsn"`
  - Resets the user's password using the reset token.
##### Request:
```json
{
  "password": "Abc123!#()"
}
```
###### Response:
```json
{
  "message": "Password changed successfully."
}
```
### 6. Logout
- **POST** `/api/v1/auth/logout`
  - Invalidates the user's access and refresh tokens.
##### Response

```json
{
  "message": "User logged out successfully."
}


```

## Account Endpoints

### 1. Update User Profile
- **PUT** `/api/v1/accounts/me`
  - Updates the user's profile details such as bio, profile picture, and contact information.


## User Management Endpoints

### 1. Promote/Demote User
- **PUT** `/api/v1/users/:id/promote`
  - Promotes or demotes a user based on the request (Admin only).






## Blog Management Endpoints

### 1. Create Blog Post
- **POST** `/api/v1/blogs`
  - Creates a new blog post with title, content, tags, etc.

### 2. Retrieve Blog Posts (with pagination and sorting)
- **GET** `/api/v1/blogs`
  - Retrieves a paginated list of blog posts, can be sorted by popularity or recent posts.

### 3. Retrieve Single Blog Post
- **GET** `/api/v1/blogs/:id`
  - Retrieves a specific blog post by ID.

### 4. Update Blog Post
- **PUT** `/api/v1/blogs/:id`
  - Updates an existing blog post (User must be the author).

### 5. Delete Blog Post
- **DELETE** `/api/v1/blogs/:id`
  - Deletes a blog post (User must be the author or an Admin).

### 6. Search Blogs
- **GET** `/api/v1/blogs/search`
  - Searches for blog posts based on title, author name, or both.

### 7. Track Blog Popularity
- **POST** `/api/v1/blogs/:id/track`
  - Tracks popularity metrics like views, likes, and comments.

### 8. Filter Blogs
- **GET** `/api/v1/blogs/filter`
  - Filters blogs by tags, date, or popularity.

## AI Integration Endpoints

### 1. AI-Generated Content Suggestions
- **POST** `/api/v1/blogs/ai/suggestions`
  - Generates content suggestions or enhancements based on user-provided keywords or topics.
