## System Overview

## Introduction
This project is a backend API designed for a blog platform. The platform allows users to manage blog posts through core CRUD operations (Create, Read, Update, Delete), and supports advanced features like user authentication, authorization, and AI-based content suggestions. The API is built with scalability, security, and performance in mind, ensuring a robust and responsive user experience.

## System Overview
The system is a comprehensive backend API designed for a blog platform, enabling users to manage their blog posts with core CRUD operations. It includes robust user management features such as registration, login, password recovery, and role-based access control, allowing differentiated access levels for regular users and admins.

Users can search and filter blog posts based on various criteria, including tags, dates, and popularity, while the system tracks metrics like views, likes, and comments to measure post popularity. Advanced AI integration is included to assist users in generating content suggestions or enhancements based on keywords or topics.

The API is built with scalability and performance in mind, leveraging Go’s concurrency model to handle multiple requests efficiently. Security is also a priority, with JWT-based authentication, secure password handling, and OAuth2 integration for third-party authentication. The system is configured for high performance and reliability, ensuring a smooth and responsive user experience even under heavy loads.

## Key Features

### 1. User Management
- **Registration:** Allows users to sign up with email, password, and profile details.
- **Login:** Authenticated access using secure credentials.
- **Forgot Password:** Enables users to reset their password securely.
- **User Promotion and Demotion:** Admins can promote users to admin status or demote them to regular users.
- **Logout:** Securely logs users out, invalidating their session tokens.

### 2. Blog Management
- **Blog Creation:** Users can create and publish new blog posts.
- **Blog Retrieval:** Provides paginated access to blog posts, along with popularity metrics.
- **Blog Update:** Allows users to update their own blog posts.
- **Blog Deletion:** Enables users or admins to delete blog posts.
- **Blog Search:** Users can search for blog posts by title or author.
- **Blog Popularity Tracking:** Tracks metrics like views, likes, and comments.

### 3. AI Integration
- **Content Generation:** Provides AI-based content suggestions based on user-provided keywords or topics.

### 4. Profile Management
- **Profile Update:** Users can update their profile details such as bio, profile picture, and contact information.







## How the System Works

### Initialization:
The `main` package initializes the application by loading configuration settings, setting up the router, and starting the HTTP server.

### Handling Requests:
When a user sends an HTTP request, the `router` directs the request to the appropriate `controller` based on the endpoint and HTTP method.

### Processing Logic:
The `controller` processes the request by interacting with the `usecases`. For example, the `BlogController` might call `FilterBlogs` from the `Usecases` layer to retrieve blog posts based on the provided filters.

### Business Rules:
The `usecases` layer executes the specific business logic required for the operation, interacting with the `domain` entities and `repository` layer to perform tasks such as saving or retrieving data.

### Data Access:
The `repository` layer handles the data operations, interfacing with the database systems to persist or fetch data as needed.

### External Services:
The `infrastructure` layer manages interactions with external services, such as sending emails or handling password resets, which are required to support various application functionalities.

### Response:
Once the `usecases` layer completes the processing, the `controller` formats the response and sends it back to the user through the HTTP response.

## Configuration and Environment

- Configuration settings and environment variables are managed through the `.env` file, ensuring that sensitive information and environment-specific configurations are handled securely and efficiently.

### .env File Includes:
- Database connection strings
- SMTP configuration for sending emails
- API keys and secrets
- JWT secret key
- Other environment-specific variables





# UserController API Documentation

## Endpoints

### 1. **Register**
   - **Description:** Registers a new user.
   - **Method:** POST
   - **Endpoint:** `/users/register`
   - **Request Body:**
     ```json
     {
       "username": "newuser",
       "password": "password123",
       "email": "newuser@example.com"
     }
     ```
   - **Response:**
     - **Success:**
       - **Status Code:** `201 Created`
       - **Example:**
         ```json
         {
           "user": {
             "username": "newuser",
             "email": "newuser@example.com"
           }
         }
         ```
     - **Error:**
       - **Status Code:** `400 Bad Request`
       - **Example:**
         ```json
         {
           "error": "Invalid input"
         }
         ```

### 2. **Update User**
   - **Description:** Updates user information.
   - **Method:** PUT
   - **Endpoint:** `/users/{username}`
   - **Request Body:**
     ```json
     {
       "email": "updatedemail@example.com",
       "password": "newpassword123"
     }
     ```
   - **Response:**
     - **Success:**
       - **Status Code:** `200 OK`
       - **Example:**
         ```json
         {
           "message": "User updated successfully"
         }
         ```
     - **Error:**
       - **Status Code:** `400 Bad Request`
       - **Example:**
         ```json
         {
           "error": "Invalid input"
         }
         ```

### 3. **Delete User**
   - **Description:** Deletes a user by username.
   - **Method:** DELETE
   - **Endpoint:** `/users/{username}`
   - **Response:**
     - **Success:**
       - **Status Code:** `200 OK`
       - **Example:**
         ```json
         {
           "message": "User deleted successfully"
         }
         ```
     - **Error:**
       - **Status Code:** `400 Bad Request`
       - **Example:**
         ```json
         {
           "error": "error"
         }
         ```

### 4. **Login**
   - **Description:** Logs in a user and generates an access token.
   - **Method:** POST
   - **Endpoint:** `/users/login`
   - **Request Body:**
     ```json
     {
       "username": "existinguser",
       "password": "password123"
     }
     ```
   - **Response:**
     - **Success:**
       - **Status Code:** `200 OK`
       - **Example:**
         ```json
         {
           "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImV4aXN0aW5ndXNlcjEiLCJyb2xlIjoidXNlcjEifQ.SiSvwA1GFbCVbyzxk4zRQl3Pne89U8IgLlDL8_p_m2c"
         }
         ```
     - **Error:**
       - **Status Code:** `400 Bad Request`
       - **Example:**
         ```json
         {
           "error": "Invalid input"
         }
         ```

### 5. **Refresh Token**
   - **Description:** Refreshes the access token using the refresh token.
   - **Method:** POST
   - **Endpoint:** `/users/refresh`
   - **Response:**
     - **Success:**
       - **Status Code:** `200 OK`
       - **Example:**
         ```json
         {
           "access_token": "new_access_token_here"
         }
         ```
     - **Error:**
       - **Status Code:** `400 Bad Request`
       - **Example:**
         ```json
         {
           "error": "Refresh token not found"
         }
         ```
       - **Status Code:** `401 Unauthorized`
       - **Example:**
         ```json
         {
           "error": "Invalid or expired token"
         }
         ```
       - **Status Code:** `401 Unauthorized`
       - **Example:**
         ```json
         {
           "error": "Failed to get username from token"
         }
         ```
       - **Status Code:** `401 Unauthorized`
       - **Example:**
         ```json
         {
           "error": "Invalid token claims"
         }
         ```


### 6. **Forgot Password**
   - **Description:** Initiates the password reset process by sending a reset token.
   - **Method:** POST
   - **Endpoint:** `/users/forgot-password`
   - **Request Body:**
     ```json
     {
       "username": "userToReset"
     }
     ```
   - **Response:**
     - **Success:**
       - **Status Code:** `200 OK`
       - **Example:**
         ```json
         {
           "token": "reset_token_here"
         }
         ```
     - **Error:**
       - **Status Code:** `400 Bad Request`
       - **Example:**
         ```json
         {
           "error": "Failed to generate reset token"
         }
         ```

### 7. **Reset Password**
   - **Description:** Resets the password using the reset token.
   - **Method:** POST
   - **Endpoint:** `/users/reset-password/{token}`
   - **Request Body:**
     ```json
     {
       "new_password": "newpassword123"
     }
     ```
   - **Response:**
     - **Success:**
       - **Status Code:** `200 OK`
       - **Example:**
         ```json
         {
           "access_token": "new_access_token"
         }
         ```
     - **Error:**
       - **Status Code:** `400 Bad Request`
       - **Example:**
         ```json
         {
           "error": "Failed to reset password"
         }
         ```

### 8. **Change Password**
   - **Description:** Changes the user's password.
   - **Method:** PUT
   - **Endpoint:** `/users/change-password`
   - **Request Body:**
     ```json
     {
       "username": "userToChange",
       "new_password": "newpassword123"
     }
     ```
   - **Response:**
     - **Success:**
       - **Status Code:** `200 OK`
       - **Example:**
         ```json
         {
           "message": "Password changed successfully"
         }
         ```
     - **Error:**
       - **Status Code:** `400 Bad Request`
       - **Example:**
         ```json
         {
           "error": "Invalid input"
         }
         ```

### 9. **Logout**
   - **Description:** Logs out a user by invalidating their session.
   - **Method:** POST
   - **Endpoint:** `/users/logout/{username}`
   - **Request Header:**
     - `Authorization`: Bearer token
   - **Response:**
     - **Success:**
       - **Status Code:** `200 OK`
       - **Example:**
         ```json
         {
           "message": "User logged out successfully"
         }
         ```
     - **Error:**
       - **Status Code:** `400 Bad Request`
       - **Example:**
         ```json
         {
           "error": "Failed to log out"
         }
         ```

This Markdown file provides a detailed overview of each endpoint in the `UserController`, including request and response formats, status codes, and error messages. Adjust as needed to match your actual API implementation.


# BlogController API Documentation

## Endpoints

### 1. **Create Blog**
   - **Description:** Creates a new blog post.
   - **Method:** POST
   - **Endpoint:** `/blogs`
   - **Request Body:**
     ```json
     {
       "title": "New Blog Title",
       "content": "Blog content goes here",
       "tags": ["tech", "programming"]
     }
     ```
   - **Response:**
     - **Success:**
       - **Status Code:** `200 OK`
       - **Example:**
         ```json
         {
           "data": {
             "id": "1",
             "title": "New Blog Title",
             "content": "Blog content goes here",
             "tags": ["tag1", "tag2"],
             "author": "username",
             "createdAt": "2024-08-19T00:00:00Z",
             "updatedAt": "2024-08-19T00:00:00Z",
             "viewCount": 0,
             "likes": [],
             "dislikes": [],
             "comments": []
           }
         }
         ```
     - **Error:**
       - **Status Code:** `400 Bad Request`
       - **Example:**
         ```json
         {
           "error": "Invalid input data."
         }
         ```
     - **Error:**
       - **Status Code:** `500 Internal Server Error.`
       ```json
        {
            "error": "Internal Server Error."
        }
        ```

### 2. **Retrieve Blogs**
   - **Description:** Retrieves a list of blogs with pagination and sorting.
   - **Method:** GET
   - **Endpoint:** `/blogs`
   - **Query Parameters:**
     - `page` (optional): Page number (default is 1)
     - `pageSize` (optional): Number of items per page (default is 20)
     - `sortBy` (optional): Sort by field (default is "date")
   - **Response:**
     - **Success:**
       - **Status Code:** `200 OK`
       - **Example:**
         ```json
         {
           "data": [
             {
               "id": "1",
               "title": "Blog Title 1",
               "content": "Blog content 1",
               "tags": ["tag1"],
               "author": "author1",
               "createdAt": "2024-08-19T00:00:00Z",
               "updatedAt": "2024-08-19T00:00:00Z",
               "viewCount": 10,
               "likes": ["user1"],
               "dislikes": [],
               "comments": []
             }
           ],
           "totalPages": 1,
           "currentPage": 1,
           "totalPosts": 1
         }
         ```
     - **Error:**
       - **Status Code:** `500 Internal Server Error`
       - **Example:**
         ```json
         {
           "error": "Error while retrieving blogs."
         }
         ```

### 3. **Delete Blog by ID**
   - **Description:** Deletes a blog post by its ID.
   - **Method:** DELETE
   - **Endpoint:** `/blogs/{id}`
   - **Response:**
     - **Success:**
       - **Status Code:** `200 OK`
       - **Example:**
         ```json
         {
           "message": "Blog post deleted successfully"
         }
         ```
     - **Error:**
       - **Status Code:** `404 Not Found`
       - **Example:**
         ```json
         {
           "error": "Blog not found with the provided ID."
         }
         ```
       - **Status Code:** `403 Forbidden`
       - **Example:**
         ```json
         {
           "error": "You are not authorized to delete this blog."
         }
         ```
       - **Status Code:** `500 Internal Server Error`
       - **Example:**
         ```json
         {
           "error": "Failed to delete blog."
         }
         ```

### 4. **Search Blogs**
   - **Description:** Searches for blogs based on title, author, and tags.
   - **Method:** GET
   - **Endpoint:** `/blogs/search`
   - **Query Parameters:**
     - `title` (optional): Title of the blog
     - `author` (optional): Author of the blog
     - `tags` (optional): Tags associated with the blog (multiple values allowed)
   - **Response:**
     - **Success:**
       - **Status Code:** `200 OK`
       - **Example:**
         ```json
         {
           "data": [
             {
               "id": "1",
               "title": "Search Result Blog Title",
               "content": "Search result blog content",
               "tags": ["tag1"],
               "author": "author1",
               "createdAt": "2024-08-19T00:00:00Z",
               "updatedAt": "2024-08-19T00:00:00Z",
               "viewCount": 5,
               "likes": ["user1"],
               "dislikes": [],
               "comments": []
             }
           ]
         }
         ```
     - **Error:**
       - **Status Code:** `500 Internal Server Error`
       - **Example:**
         ```json
         {
           "error": "Error while searching for blogs."
         }
         ```

### 5. **Update Blog**
   - **Description:** Updates an existing blog post by its ID.
   - **Method:** PUT
   - **Endpoint:** `/blogs/{id}`
   - **Request Body:**
     ```json
     {
       "title": "Updated Blog Title",
       "content": "Updated blog content",
       "tags": ["updatedTag1"]
     }
     ```
   - **Response:**
     - **Success:**
       - **Status Code:** `200 OK`
       - **Example:**
         ```json
         {
           "data": {
             "id": "1",
             "title": "Updated Blog Title",
             "content": "Updated blog content",
             "tags": ["updatedTag1"],
             "author": "username",
             "createdAt": "2024-08-19T00:00:00Z",
             "updatedAt": "2024-08-19T00:00:00Z",
             "viewCount": 10,
             "likes": ["user1"],
             "dislikes": [],
             "comments": []
           }
         }
         ```
     - **Error:**
       - **Status Code:** `400 Bad Request`
       - **Example:**
         ```json
         {
           "error": "Invalid input data."
         }
         ```
       - **Status Code:** `401 Unauthorized`
       - **Example:**
         ```json
         {
           "error": "User not found in context."
         }
         ```
       - **Status Code:** `500 Internal Server Error`
       - **Example:**
         ```json
         {
           "error": "Error while updating the blog."
         }
         ```
### 6. **Increment View Count**
   - **Description:** Increments the view count of a blog post by its ID.
   - **Method:** POST
   - **Endpoint:** `/blogs/:id/views`
   - **Response:**
     - **Success:**
       - **Status Code:** `200 OK`
       - **Example:**
         ```json
         {
           "message": "View count updated"
         }
         ```
     - **Error:**
       - **Status Code:** `500 Internal Server Error`
       - **Example:**
         ```json
         {
           "error": "Failed to increment view count"
         }
         ```

### 7. **Toggle Like**
   - **Description:** Toggles the like status for a blog post by its ID.
   - **Method:** POST
   - **Endpoint:** `/blogs/:id/like`
   - **Response:**
     - **Success:**
       - **Status Code:** `200 OK`
       - **Example:**
         ```json
         {
           "message": "Like toggled"
         }
         ```
     - **Error:**
       - **Status Code:** `500 Internal Server Error`
       - **Example:**
         ```json
         {
           "error": "Failed to toggle like"
         }
         ```

### 8. **Toggle Dislike**
   - **Description:** Toggles the dislike status for a blog post by its ID.
   - **Method:** POST
   - **Endpoint:** `/blogs/:id/dislike`
   - **Response:**
     - **Success:**
       - **Status Code:** `200 OK`
       - **Example:**
         ```json
         {
           "message": "Dislike toggled"
         }
         ```
     - **Error:**
       - **Status Code:** `500 Internal Server Error`
       - **Example:**
         ```json
         {
           "error": "Failed to toggle dislike"
         }
         ```

### 9. **Add Comment**
   - **Description:** Adds a comment to a blog post by its ID.
   - **Method:** POST
   - **Endpoint:** `/blogs/:id/comments`
   - **Request Body:**
     - **Content-Type:** `application/json`
     - **Body:**
       ```json
       {
         "content": "This is a comment"
       }
       ```
   - **Response:**
     - **Success:**
       - **Status Code:** `200 OK`
       - **Example:**
         ```json
         {
           "message": "Comment added"
         }
         ```
     - **Error:**
       - **Status Code:** `400 Bad Request`
       - **Example:**
         ```json
         {
           "error": "Invalid input"
         }
         ```
       - **Status Code:** `500 Internal Server Error`
       - **Example:**
         ```json
         {
           "error": "Failed to add comment"
         }
  
  ## Endpoints

### 10. **Filter Blogs**
   - **Description:** Filters and retrieves blog posts based on tags, date range, and sorting preferences.
   - **Method:** GET
   - **Endpoint:** `/blogs/filter`
   - **Query Parameters:**
     - `tags` (optional): A comma-separated list of tags to filter blogs by. Example: `technology,programming`.
     - `startDate` (optional): The start date for filtering blog posts. The date should be in RFC 3339 format. Example: `2024-01-01T00:00:00Z`.
     - `endDate` (optional): The end date for filtering blog posts. The date should be in RFC 3339 format. Example: `2024-12-31T23:59:59Z`.
     - `sortBy` (optional): The criteria by which to sort the results. Example: `date` to sort by creation date, or `views` to sort by view count.
   - **Response:**
     - **Success:**
       - **Status Code:** `200 OK`
       - **Example:**
         ```json
         [
           {
             "id": "1",
             "title": "Introduction to Go",
             "content": "Go is an open-source programming language designed for simplicity...",
             "tags": ["programming", "golang"],
             "author": "John Doe",
             "createdAt": "2024-02-15T10:00:00Z",
             "updatedAt": "2024-03-01T15:30:00Z",
             "viewCount": 123,
             "likes": 45,
             "dislikes": 2,
             "comments": ["Great post!", "Very informative."]
           }
         ]
         ```
     - **Error:**
       - **Status Code:** `500 Internal Server Error`
       - **Example:**
         ```json
         {
           "error": "An unexpected error occurred while processing the request."
         }
         ```

       ```


### 11. **Chat**
   - **Description:** Handles user input and interacts with the chat service to get a response based on the provided prompt.
   - **Method:** POST
   - **Endpoint:** `/chat`
   - **Request Body:**
     - **Content-Type:** `application/json`
     - **Body:**
       ```json
       {
         "prompt": "chat prompt here"
       }
       ```
   - **Response:**
     - **Success:**
       - **Status Code:** `200 OK`
       - **Example:**
         ```json
         {
           "message": "Chat response based on the prompt"
         }
         ```
     - **Error:**
       - **Status Code:** `400 Bad Request`
       - **Example:**
         ```json
         {
           "error": "Invalid input"
         }
         ```
       - **Status Code:** `500 Internal Server Error`
       - **Example:**
         ```json
         {
           "error": "Error while processing the chat request."
         }
         ```

# Project Folder Structure Description

## Delivery/
Contains files related to the delivery layer, handling incoming requests and responses.

- **config/config.go**: Manages configuration settings, such as environment variables and application-specific configurations.
- **main.go**: Sets up the HTTP server, initializes dependencies, and defines the routing configuration.
- **controllers/**: Handles incoming HTTP requests and invokes the appropriate use case methods.
  - **aiController.go**: Manages requests related to AI functionalities, such as generating content suggestions.
  - **blogController.go**: Manages requests related to blog operations, including creating, updating, deleting, and retrieving blog posts.
  - **userController.go**: Manages user-related requests, including registration, login, and profile management.
- **router/router.go**: Sets up the routes and initializes the Gin router, connecting endpoints with corresponding controllers.

## Domain/
Defines the core business entities and logic.

- **blog.go**: Contains the `Blog` struct and related methods, representing blog posts in the system.
- **chat.go**: Manages entities related to chat or messaging within the platform.
- **comment.go**: Defines the structure and behavior of comments on blog posts.
- **token.go**: Manages token-related entities, including JWT tokens for authentication.
- **user.go**: Contains the `User` struct and related methods, representing users in the system.

## Infrastructure/
Implements external dependencies and services.

- **auth_middleWare.go**: Middleware to handle authentication and authorization using JWT tokens.
- **jwt_service.go**: Functions to generate and validate JWT tokens for secure user sessions.
- **email_services.go**: Manages the sending of emails, such as password resets or account activations.
- **utils.go**: Provides utility functions that are used across various parts of the application.
- **password_service.go**: Functions for hashing and comparing passwords to ensure secure storage of user credentials.

## Repositories/
Abstracts the data access logic, interfacing with the database or other storage systems.

- **blogRepository.go**: Interface and implementation for blog data access operations, including CRUD functionality.
- **userRepository.go**: Interface and implementation for user data access operations, such as retrieving and storing user information.

## Usecases/
Contains the application-specific business rules.

- **AIUsecases.go**: Implements use cases related to AI features, such as generating content or suggestions.
- **blogUsecases.go**: Implements use cases related to blog operations, including creating, updating, deleting, and searching for blog posts.
- **userUsecases.go**: Implements use cases related to users, including registration, authentication, and profile management.

