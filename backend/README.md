````markdown
# AASTU Backend Group 6 Project 🎉

Welcome to the AASTU Backend Group 6 project! This repository contains the backend implementation for our blog platform, developed using Go, Gin, and MongoDB. Our project is designed to handle user authentication, blog management, and more, with a clean architecture approach. 🚀

## Project Structure 🗂️

```plaintext
backend
 ┣ AASTU-backend-group-6
 ┃ ┣ cmd
 ┃ ┃ ┣ tmp
 ┃ ┃ ┃ ┗ main.exe
 ┃ ┃ ┣ .env
 ┃ ┃ ┗ main.go
 ┃ ┣ Delivery
 ┃ ┃ ┣ controllers
 ┃ ┃ ┣ routers
 ┃ ┃ ┗ tmp
 ┃ ┣ doc
 ┃ ┃ ┗ doc.md
 ┃ ┣ Domain
 ┃ ┣ Infrastructure
 ┃ ┣ mocks
 ┃ ┣ mongo
 ┃ ┣ Repositories
 ┃ ┣ Test
 ┃ ┣ Usecases
 ┃ ┣ Utils
 ┃ ┣ .env
 ┃ ┣ .gitignore
 ┃ ┣ air.cfg
 ┃ ┣ bash.sh
 ┃ ┣ go.mod
 ┃ ┣ go.sum
 ┃ ┣ README.md
 ┃ ┗ run.sh
```
````

## Environment Variables 🌍

To run this project, you will need to add the following environment variables to your `.env` file:

```plaintext
DATABASE_URL=mongodb+srv://<username>:<password>@cluster0.mongodb.net/
PORT=8080
JWT_SECRET=<your-jwt-secret>
DB_NAME=G6_Blog
user_collection=user
blog_collection=blog
CONTEXT_TIMEOUT=10
ACCESS_TOKEN_EXPIRY_HOUR=2
REFRESH_TOKEN_EXPIRY_HOUR=168
ACCESS_TOKEN_SECRET=<your-access-token-secret>
REFRESH_TOKEN_SECRET=<your-refresh-token-secret>
CLIENT_ID=<your-client-id>
CLIENT_SECRET=<your-client-secret>
REDIRECT_URI=<your-callback-url>
OAUTH_STATE_STRING=oauthStateString
ACTIVE_USER_COLLECTION=active_user
GEMINI_API_KEY=<your-api-key>
```

## Endpoints 🛤️

### Blog Endpoints 📝

- `GET /blogs/` - Retrieve all blogs
- `GET /blogs/:id` - Retrieve a blog by ID
- `GET /blogs/search` - Search blogs by title and author
- `GET /blogs/filter` - Filter blogs by tag
- `GET /blogs/my` - Retrieve authenticated user's blogs (Protected)
- `GET /blogs/my/:id` - Retrieve a specific blog of the authenticated user (Protected)
- `POST /blogs/create` - Create a new blog (Protected)
- `PUT /blogs/update/:id` - Update a blog by ID (Protected)
- `DELETE /blogs/delete/:id` - Delete a blog by ID (Protected)
- `POST /blogs/comment/create` - Comment on a blog (Protected)
- `POST /blogs/react/:id` - React to a blog (Protected)

### User Registration Endpoints 👤

- `POST /auth/signup` - Register a new user
- `POST /auth/verify` - Verify a user's email address using OTP
- `POST /auth/resend` - Resend a verification email
- `POST /auth/reset` - Send a password reset email
- `POST /auth/reset/token` - Reset a user's password
- `POST /auth/login` - Login a user
- `POST /auth/logout` - Logout a user (Protected)
- `Post /auth/google` - Login with Google
- `GET /auth/callback` - Google OAuth callback
- `POST /auth/refresh` - Refresh access token (Protected)
-

### User Endpoints 👤

- `PUT /users/update/:id` - Update user information
- `POST /users/promote/:id` - Promote a user to a higher role

## How to Run 🏃‍♂️

1. Clone the repository: `git clone https://github.com/RealEskalate/a2sv-g5-project-phase-starter-project`
2. Navigate to the project directory: `cd backend/AASTU-backend-group-6`
3. Set up your environment variables in the `.env` file.
4. Run the application: `go run cmd/main.go`

## Contributing 🤝

We welcome contributions! Please fork the repository and create a pull request with your changes.

## License 📄

This project is licensed under the MIT License. See the `LICENSE` file for more details.

---

Happy coding! 💻✨

```

Feel free to adjust the emojis or text as needed!
```
