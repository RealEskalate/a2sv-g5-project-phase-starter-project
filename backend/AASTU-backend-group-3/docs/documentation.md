# 🧪 Teasted Clean Architecture RESTful API with a MongoDB Database and JWT User Authentication

This repository contains a backend RESTful API built using the principles of Clean Architecture in Go with the Gin web framework and MongoDB as the database.

## ✨ Features

- 👤 **User Management**: Register, login, and manage user profiles with role-based access control.
- 📝 **Task Management**: Create, read, update, and delete tasks, with personal and admin views.
- 🔒 **Secure**: Authentication and authorization using middleware.
- ⚙️ **Scalable and Maintainable**: Separation of concerns with clear boundaries between different layers of the application.

## 📚 API Documentation

For detailed API endpoints and usage instructions, please refer to the [API Documentation](https://documenter.getpostman.com/view/32898780/2sA3s1oruU).

## ⚙️ Installation and Setup

### ✅ Prerequisites

- 🐹 Go 1.16+
- 🍃 MongoDB instance

### 🛠️ Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/teklumt/A2SV-Backend-Tasks-2024.git
   cd A2SV-Backend-Tasks-2024/Task7-%20Clean%20Architecture
   ```

2. **Install dependencies**:

   ```bash
   go mod tidy
   ```

3. **Setup environment variables**:

   - Create a `.env` file in the root directory.
   - Add the necessary environment variables such as MongoDB URI.

4. **Run the application**:
   ```bash
   go run main.go
   ```

## 🚀 Usage

- **User Operations**:

  - `GET /users`: Get all users (Admin only) 👨‍💻.
  - `GET /users/:id`: Get a user by ID 🆔.
  - `GET /users/me`: Get the authenticated user's profile 🧑‍💼.
  - `DELETE /users/:id`: Delete a user by ID ❌.

- **Task Operations**:
  - `POST /tasks`: Create a new task ➕.
  - `GET /tasks`: Get all tasks (Admin only) 📄.
  - `GET /tasks/:id`: Get a task by ID 🆔.
  - `GET /tasks/me`: Get tasks assigned to the authenticated user 🗒️.
  - `DELETE /tasks/:id`: Delete a task by ID ❌.
  - `PUT /tasks/:id`: Update a task by ID ✏️.
