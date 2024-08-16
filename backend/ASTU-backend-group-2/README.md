# Blog Starter Project!

This repository contains the backend API for a blog platform that allows users to create, read, update, and delete blog posts, manage user profiles, and perform advanced search and filtering operations.

## Overview

The goal of this project is to develop a scalable, secure, and efficient backend API for a blog platform with the following key features:

- `User Management`: Register, log in, manage profiles, and authenticate users.
- `Blog Management`: CRUD operations for blog posts, including search, filtration, and popularity tracking.
- `AI Integration`: Generate content suggestions or enhancements based on user input.
- `Role-Based Access Control`: Different user roles (e.g., Admin, User) with specific permissions.

## Project Structure

The project follows the Clean Architecture pattern and is organized as follows:

```bash
├── api
│   ├── controller
│   ├── middleware
│   └── route
├── bootstrap
│   ├── app.go
│   ├── database.go
│   └── env.go
├── cmd
│   └── main.go
├── domain
├── go.mod
├── go.sum
├── internal
├── mongo
├── repository
└── usecase
```

## Features

### User Management

- `User Registration`: Register with an email, password, and profile details.
- `Login & Authentication`: Secure user login with JWT tokens.
- `Password Reset`: Forgot password functionality.
- `Profile Management`: Update user profile details.
- `Role Management`: Admins can promote or demote users.

### Blog Management

- `Create Blog Posts`: Authenticated users can create blog posts.
- `Retrieve Blogs`: View all blog posts with pagination and popularity metrics.
- `Update & Delete Blogs`: Authenticated users can update or delete their own posts.
- `Search & Filter`: Search blog posts by title, author, tags, etc.

### AI Integration

- `Content Suggestions`: Generate blog content based on user-provided keywords or topics.

## Getting Started

### Prerequisites

- `Go`: Ensure you have Go installed on your system.
- `MongoDB`: Set up a MongoDB instance for data storage.

### Installation

1. Clone the repository:

```bash
git clone github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2
cd blog-starter-project
```

2. Install dependencies:

```bash
go mod tidy
```

3. Set up environment variables:

```bash
cp .env.example .env
```

4. Run the application:

```bash
go run cmd/main.go
```

## Naming Conventions

#### Package Names:

- Style: Lowercase, short, and meaningful.
- Example: `domain`.

#### File Names:

- Style: Lowercase, with words separated by underscores.
- Example: `task.go`, `task_repository.go`, `task_usecase.go`.

#### Interfaces:

- Style: CamelCase, with descriptive names.
- Example: `TaskRepository`, `TaskUsecase`.

#### Struct Names:

- Style: CamelCase, start with a noun.
- Example: `Task`.

#### Function Names:

- Exported Functions: CamelCase, start with an uppercase letter.
- Examples: `Create`, `FetchByUserID`.
- Unexported Functions: camelCase, start with a lowercase letter.
- Example: `parseTask`, `validateTask` (if applicable).

#### Variable Names:

- Style: camelCase for regular variables, UPPERCASE with underscores for constants.
- Examples: `task`, `userID`, `collectionTask`.

#### Constants:

- Style: UPPERCASE with underscores.
- Example: `COLLECTION_TASK`.

## Contributing

We welcome contributions! Please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Commit your changes.
4. Push to your branch and create a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
