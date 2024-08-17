### 1. **Signup**

- **Endpoint:** `POST /api/v1/auth/signup`
- **Request Body:**
  ```json
  {
    "firstName": "string",
    "lastName": "string",
    "userName": "string",
    "email": "string",
    "password": "string"
  }
  ```
- **Response:**
  - **Status:** `201 Created`
  - **Body:**
    ```json
    {
      "id": "00000000-0000-0000-0000-000000000000",
      "username": "string",
      "firstName": "string",
      "lastName": "string",
      "email": "string",
      "isAdmin": true
    }
    ```
  - **Headers:**
    ```
    Set-Cookie: accessToken=<token_value>; HttpOnly; Secure
    Set-Cookie: refreshToken=<token_value>; HttpOnly; Secure
    ```

### 2. **Login**

- **Endpoint:** `POST /api/v1/auth/login`
- **Request Body:**
  ```json
  {
    "username": "string",
    "password": "string"
  }
  ```
- **Response:**
  - **Status:** `200 OK`
  - **Body:**
    ```json
    {
      "id": "00000000-0000-0000-0000-000000000000",
      "username": "string",
      "firstName": "string",
      "lastName": "string",
      "email": "string",
      "isAdmin": true
    }
    ```
  - **Headers:**
    ```
    Set-Cookie: accessToken=<token_value>; HttpOnly; Secure
    Set-Cookie: refreshToken=<token_value>; HttpOnly; Secure
    ```

### 3. **Forgot Password**

- **Endpoint:** `POST /api/v1/auth/forgot-password`
- **Request Body:**
  ```json
  {
    "email": "string"
  }
  ```
- **Response:**
  - **Status:** `204 No Content`

### 4. **Reset Password**

- **Endpoint:** `POST /api/v1/auth/reset-password`
- **Request Body:**
  ```json
  {
    "resetToken": "string",
    "newPassword": "string"
  }
  ```
- **Response:**
  - **Status:** `204 No Content`

### 5. **Logout**

- **Endpoint:** `POST /api/v1/auth/logout`
- **Request Body:** `None`
- **Response:**
  - **Status:** `204 No Content`
  - **Headers:**
    ```
    Set-Cookie: accessToken=; HttpOnly; Secure
    Set-Cookie: refreshToken=; HttpOnly; Secure
    ```

### 6. **Promotion**

- **Endpoint:** `POST /api/v1/users/{username}/promote`
- **Response:**
  - **Status:** `204 No Content`

### 7. **Demotion**

- **Endpoint:** `POST /api/v1/users/{username}/demote`
- **Response:**
  - **Status:** `204 No Content`

## Blog Management

### 8. **Create Blog**

- **Endpoint:** `POST /api/v1/blogs`
- **Request Body:**
  ```json
  {
    "title": "string",
    "content": "string",
    "tags": ["string"],
    "author": "string"
  }
  ```
- **Response:**
  - **Status:** `201 Created`
  - **Body:**
    ```json
    {
      "id": "00000000-0000-0000-0000-000000000000",
      "title": "string",
      "content": "string",
      "tags": ["string"],
      "createdAt": "2024-08-16T00:00:00Z",
      "updatedAt": "2024-08-16T00:00:00Z",
      "likeCount": 0,
      "dislikeCount": 0,
      "viewCount": 0,
      "commentCount": 0
    }
    ```

### 9. **Get All Blogs**

- **Endpoint:** `GET /api/v1/blogs?cursor={base64_string}&limit={number}&sortField={field}&filterField={field}&filterValue={value}&sortOrder={asc|desc}`
- **Response:**
  - **Status:** `200 OK`
  - **Body:**
    ```json
    {
      "blogs": [
        {
          "id": "00000000-0000-0000-0000-000000000000",
          "title": "string",
          "content": "string",
          "tags": ["string"],
          "createdAt": "2024-08-16T00:00:00Z",
          "updatedAt": "2024-08-16T00:00:00Z",
          "likeCount": 0,
          "dislikeCount": 0,
          "viewCount": 0,
          "commentCount": 0,
          "author": "string"
        }
      ],
      "cursor": "base64_string"
    }
    ```

### 10. **Update Blog**

- **Endpoint:** `PATCH /api/v1/blogs/{id}`
- **Request Body:**
  ```json
  {
    "title": "string",
    "content": "string",
    "tags": ["string"]
  }
  ```
- **Response:**
  - **Status:** `204 No Content`

### 11. **Delete Blog**

- **Endpoint:** `DELETE /api/v1/blogs/{id}`
- **Response:**
  - **Status:** `204 No Content`

### 12. **Popularity Tracking**

- **Endpoint:** `POST /api/v1/blogs/{id}/react`
- **Request Body:**
  ```json
  {
    "reactionType": "like" | "dislike" | "view"
  }
  ```
- **Response:**
  - **Status:** `204 No Content`

## User Profile

### 13. **Edit Profile**

- **Endpoint:** `PATCH /api/v1/profile`
- **Request Body:**
  ```json
  {
    "firstName": "string",
    "lastName": "string",
    "email": "string",
    "userName": "string"
  }
  ```
- **Response:**
  - **Status:** `204 No Content`

## Comment Management

### 14. **Add Comment**

- **Endpoint:** `POST /api/v1/blogs/{id}/comments`
- **Request Body:**
  ```json
  {
    "content": "string"
  }
  ```
- **Response:**
  - **Status:** `201 Created`
  - **Body:**
    ```json
    {
      "id": "00000000-0000-0000-0000-000000000000",
      "content": "string",
      "author": "string",
      "createdAt": "2024-08-16T00:00:00Z"
    }
    ```

### 15. **Get Comments**

- **Endpoint:** `GET /api/v1/blogs/{id}/comments`
- **Response:**
  - **Status:** `200 OK`
  - **Body:**
    ```json
    {
      "comments": [
        {
          "id": "00000000-0000-0000-0000-000000000000",
          "content": "string",
          "author": "string",
          "createdAt": "2024-08-16T00:00:00Z"
        }
      ]
    }
    ```
