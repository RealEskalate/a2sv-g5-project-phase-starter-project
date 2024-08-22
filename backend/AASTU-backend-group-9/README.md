# Blog Starter Project

## Introduction

The **Blog Starter Project** is a backend API designed to power a comprehensive blog platform. This project provides a foundation for creating, managing, and interacting with blog content. It includes features like user authentication, blog management, advanced search and filtering, and AI-driven content suggestions.

## Features

- **User Authentication & Authorization:**
  - User registration, login, and profile management.
  - Role-based access control (admin/user).
  - Secure JWT-based token management with access and refresh tokens.

- **Blog Management:**
  - CRUD operations for blog posts.
  - Tagging, filtering, and search functionality.
  - Blog popularity tracking with view counts, likes, dislikes, and comments.

- **AI Integration:**
  - AI-powered content suggestions based on user input.

- **Advanced Features:**
  - Pagination for efficient data retrieval.
  - Profile management for users.
  - Robust error handling and validation.

## Tech Stack

- **Programming Language:** Go
- **Framework:** Gin
- **Database:** MongoDB
- **Authentication:** JWT (JSON Web Tokens)
- **Documentation:** Postman

### Key Components:

- **Delivery:** Contains the controllers and routers, handling HTTP requests and responses.
- **Domain:** Contains the core business logic and domain models.
- **Infrastructure:** Handles external services and infrastructure-related code such as authentication and database connections.
- **Repositories:** Implements the data access layer, interacting with the MongoDB database.
- **Usecases:** Contains the application-specific business logic.

## Getting Started

### Prerequisites

- Go 1.19 or higher
- MongoDB
- Postman (for API testing)

API Documentation
The API is documented using Postman. You can import the Postman collection to explore the endpoints.

https://documenter.getpostman.com/view/36737395/2sAXjDevNy

## Getting Started

### Prerequisites

Before you begin, ensure you have the following installed on your machine:

- [Go](https://golang.org/doc/install) (version 1.19 or higher)
- [MongoDB](https://www.mongodb.com/)

### Setup

1. **Clone the repository:**

    ```sh
    git clone https://github.com/RealEskalate/a2sv-g5-project-phase-starter-project.git
    cd aastu-project/backend/AASTU-backend-group-9
    ```
2. **Set up environment variables:**

    Copy the `.env.example` file to `.env` and update the environment variables as needed.

    ```sh
    cp cmd/.env.example cmd/.env
    ```

3. **Install dependencies:**

    ```sh
    go mod download
    ```

4. **Run the backend server:**

    ```sh
    go run cmd/main.go
    ```