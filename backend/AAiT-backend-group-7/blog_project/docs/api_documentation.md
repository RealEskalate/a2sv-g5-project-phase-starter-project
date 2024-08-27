### Blog Application API

This project is a blog platform API built with Go using the Gin web framework. It allows users to create, manage, and interact with blog posts. The platform supports user authentication, blog management, and interactions like commenting, liking, and disliking posts. Additionally, the system integrates AI-based content recommendations.

### Running the Project

To run the project, follow these steps:

1. Start MongoDB:
   ```bash
   mongod
   ```
2. (Optional) Start Redis if your setup includes caching or token management with Redis:
   ```bash
   redis-server
   ```
3. Navigate to the `delivery` directory:
   ```bash
   cd delivery
   ```
4. Run the application:
   ```bash
   go run .
   ```

### API Documentation

#### **Users API**

**Base Path:** `/users`

1. **Get All Users**
   - **Endpoint:** `GET /users`
   - **Description:** Retrieves a list of all users.
   - **Response:** JSON array of user objects.

2. **Get User by ID**
   - **Endpoint:** `GET /users/:id`
   - **Description:** Retrieves details of a specific user by ID.
   - **Parameters:**
     - `id` (path): User ID.
   - **Response:** JSON object of the user.

3. **Create User**
   - **Endpoint:** `POST /users`
   - **Description:** Creates a new user.
   - **Request Type:** `multipart/form-data`
   - **Form Data Fields:**
     - `username`: (string) Required.
     - `password`: (string) Required.
     - `email`: (string) Required.
     - `bio`: (string) Optional.
     - `phone`: (string) Optional.
     - `profile_pic`: (file) Optional.
   - **Response:** JSON object of the newly created user.

4. **Update User**
   - **Endpoint:** `PUT /users/:id`
   - **Description:** Updates the details of an existing user.
   - **Parameters:**
     - `id` (path): User ID.
   - **Request Type:** `multipart/form-data`
   - **Form Data Fields:**
     - `username`: (string) Optional.
     - `password`: (string) Optional.
     - `email`: (string) Optional.
     - `bio`: (string) Optional.
     - `phone`: (string) Optional.
     - `profile_pic`: (file) Optional.
   - **Response:** JSON object of the updated user.

5. **Delete User**
   - **Endpoint:** `DELETE /users/:id`
   - **Description:** Deletes a user by ID.
   - **Parameters:**
     - `id` (path): User ID.
   - **Response:** JSON message confirming deletion.

6. **Login**
   - **Endpoint:** `POST /users/login`
   - **Description:** Authenticates a user and returns an access token and a refresh token.
   - **Request Type:** `application/json`
   - **Request Body Fields:**
     - `username`: (string) Required.
     - `password`: (string) Required.
   - **Response:** JSON object containing `access_token`, `refresh_token`, and user details.

7. **Logout**
   - **Endpoint:** `POST /users/logout`
   - **Description:** Logs out a user by invalidating the provided token.
   - **Request Type:** `application/json`
   - **Request Body Fields:**
     - `token`: (string) Required.
   - **Response:** JSON message confirming logout.

8. **Forget Password**
   - **Endpoint:** `GET /users/forget-password`
   - **Description:** Sends a password reset link to the user's email.
   - **Query Parameters:**
     - `email`: (string) Required.
   - **Response:** JSON message confirming the email was sent.

9. **Reset Password**
   - **Endpoint:** `PUT /users/reset-password/:username/:password`
   - **Description:** Resets the user's password.
   - **Parameters:**
     - `username` (path): User's username.
     - `password` (path): New password.
   - **Response:** JSON message confirming the password reset.

10. **Promote User**
    - **Endpoint:** `PUT /users/promote/:id`
    - **Description:** Promotes a user to a higher role.
    - **Parameters:**
      - `id` (path): User ID.
    - **Response:** JSON message confirming promotion.

11. **Demote User**
    - **Endpoint:** `PUT /users/demote/:id`
    - **Description:** Demotes a user to a lower role.
    - **Parameters:**
      - `id` (path): User ID.
    - **Response:** JSON message confirming demotion.

12. **Refresh Token**
    - **Endpoint:** `POST /users/refresh-token`
    - **Description:** Refreshes the access token using a valid refresh token.
    - **Request Type:** `application/json`
    - **Request Body Fields:**
      - `refresh_token`: (string) Required.
    - **Response:** JSON object with the new access token.

#### **Blogs API**

**Base Path:** `/blogs`

1. **Get All Blogs**
   - **Endpoint:** `GET /blogs`
   - **Description:** Retrieves a paginated list of all blogs.
   - **Query Parameters:**
     - `page`: (int) Optional. Default is 1.
     - `limit`: (int) Optional. Default is 10.
     - `sortOrder`: (string) Optional. Can be `asc` or `desc`.
   - **Response:** JSON array of blog objects.

2. **Get Blog by ID**
   - **Endpoint:** `GET /blogs/:id`
   - **Description:** Retrieves details of a specific blog by ID.
   - **Parameters:**
     - `id` (path): Blog ID.
   - **Response:** JSON object of the blog.

3. **Create Blog**
   - **Endpoint:** `POST /blogs`
   - **Description:** Creates a new blog.
   - **Request Type:** `application/json`
   - **Request Body Fields:**
     - `title`: (string) Optional.
     - `content`: (string) Optional.
     - `tags`: (array of strings) Optional.
   - **Response:** JSON object of the newly created blog.

4. **Update Blog**
   - **Endpoint:** `PUT /blogs/:id`
   - **Description:** Updates an existing blog.
   - **Parameters:**
     - `id` (path): Blog ID.
   - **Request Type:** `application/json`
   - **Request Body Fields:**
     - `title`: (string) Optional.
     - `content`: (string) Optional.
     - `tags`: (array of strings) Optional.
   - **Response:** JSON object of the updated blog.

5. **Delete Blog**
   - **Endpoint:** `DELETE /blogs/:id`
   - **Description:** Deletes a blog by ID.
   - **Parameters:**
     - `id` (path): Blog ID.
   - **Response:** JSON message confirming deletion.

6. **Add Comment**
   - **Endpoint:** `POST /blogs/comments/:id`
   - **Description:** Adds a comment to a blog.
   - **Parameters:**
     - `id` (path): Blog ID.
   - **Request Type:** `application/json`
   - **Request Body Fields:**
     - `content`: (string) Required.
   - **Response:** JSON object of the blog with the new comment.

7. **Like Blog**
   - **Endpoint:** `POST /blogs/like/:id`
   - **Description:** Likes a blog.
   - **Parameters:**
     - `id` (path): Blog ID.
   - **Response:** JSON object of the blog with the updated likes.

8. **Dislike Blog**
   - **Endpoint:** `POST /blogs/dislike/:id`
   - **Description:** Dislikes a blog.
   - **Parameters:**
     - `id` (path): Blog ID.
   - **Response:** JSON object of the blog with the updated dislikes.

9. **Search Blogs**
   - **Endpoint:** `GET /blogs/search`
   - **Description:** Searches blogs by title, tags, or author.
   - **Query Parameters:**
     - `title`: (string) Optional.
     - `tags`: (array of strings) Optional.
     - `author`: (string) Optional.
   - **Response:** JSON array of blogs matching the search criteria.

10. **AI Recommendation**
    - **Endpoint:** `POST /blogs/recommendation`
    - **Description:** Generates AI-based content recommendations based on the given content.
    - **Request Type:** `application/json`
    - **Request Body Fields:**
      - `content`: (string) Required.
    - **Response:** JSON object with the AI-generated recommendation.

