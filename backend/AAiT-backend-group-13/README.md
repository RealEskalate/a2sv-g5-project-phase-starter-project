# Blog System API

A RESTful API for managing blog posts, comments, reactions, and users. The system includes functionalities for liking and disliking posts, managing user accounts, and resetting passwords.

## Features

- Create, update, delete, and retrieve blog posts.
- Add comments to blog posts.
- Like or dislike blog posts.
- Manage user accounts with registration, login, and password reset functionality.

## Installation

1. **Clone the repository:**

2. **Set up the environment:**

   - Create a `.env` file in the project root and add the necessary environment variables.
   - Refer to `example.env` for required variables.

3. **Install dependencies:**

   ```bash
   go mod tidy
   ```

4. **Run the application:**
   ```bash
   make run
   ```

## Usage

### Running the server

Start the server using the command:

```bash
make run
```

The API will be available at `http://localhost:8080`.

### Running Tests

Run tests using the following command:

```bash
make test
```

## API Documentation

### Authentication Endpoints

- **POST /api/v1/auth/signup** - User registration.
- **POST /api/v1/auth/login** - User login.
- **POST /api/v1/auth/resetPasswordCode** - Request a password reset code.
- **POST /api/v1/auth/validateCode** - Validate password reset code.
- **POST /api/v1/auth/resetPassword** - Reset user password.
- **POST /api/v1/auth/validateEmail** - Validate user email.
- **POST /api/v1/auth/:username/logout** - Log out user.

### Blog Endpoints

- **POST /api/v1/blogs** - Create a new blog post.
- **PUT /api/v1/blogs/:id** - Update a blog post.
- **DELETE /api/v1/blogs/:id** - Delete a blog post.
- **GET /api/v1/blogs/** - Retrieve all blog posts.
- **GET /api/v1/blogs/:id** - Retrieve a blog post by ID.

### Comment Endpoints

- **POST /api/v1/blogs/:id/comments** - Add a comment to a blog post.
- **GET /api/v1/blogs/:id/comments** - Retrieve comments for a blog post.

### Reaction Endpoints

- **POST /api/v1/blogs/:id/reactions** - Like or dislike a blog post.
- **GET /api/v1/blogs/:id/reactions** - Retrieve reactions for a blog post.

### User Management Endpoints

- **PUT /api/v1/auth/update/:id** - Update user profile.
- **POST /api/v1/users/:username/promote** - Promote a user to admin.
- **POST /api/v1/users/:username/demote** - Demote an admin to a regular user.

### Additional Endpoints

- **POST /api/v1/recommendation** - Submit a recommendation.
- **POST /api/v1/review** - Submit a review.

For more detailed API documentation, refer to the [API Definition](./docs/api_definition.md).

## Project Structure

The project follows a layered architecture:

- **api/**: Handles HTTP routes and controllers.
- **application/**: Contains the application logic, use cases, and services.
- **domain/**: Includes domain models and business rules.
- **infrastructure/**: Handles database access, external services, and third-party libraries.
- **config/**: Manages configuration settings.
