# Blog Starter Backend Project

![Go](https://img.shields.io/badge/Go-1.18+-blue.svg)
![Gin](https://img.shields.io/badge/Gin-1.7+-green.svg)
![MongoDB](https://img.shields.io/badge/MongoDB-5.0+-brightgreen.svg)
![Redis](https://img.shields.io/badge/Redis-6.2+-red.svg)
![Docker](https://img.shields.io/badge/Docker-20.10+-blue.svg)
![Gemini AI](https://img.shields.io/badge/Gemini_AI-Integrated-orange.svg)
![License](https://img.shields.io/badge/License-MIT-yellow.svg)

## Overview

The **Blog Starter Backend Project** is a backend solution for a blogging platform, built with scalability, security, and maintainability in mind. It integrates modern technologies and offers advanced features like AI-assisted content creation, caching, OAuth authentication, and more.

## Features

- **Concurrency:** Efficient request handling using Go's goroutines.
- **Caching:** Redis integration for quick data retrieval and rate limiting.
- **OAuth & JWT Authentication:** Secure user authentication with JWT tokens.
- **Rate Limiting:** Prevent abuse with request rate limiting.
- **AI Integration:** Utilize AI for blog content generation and improvement.
- **Device-Aware Registration:** OTP for mobile, and link-based password setting for web users.
- **CRUD Operations:** Full Create, Read, Update, Delete functionality for blog posts.
- **Commenting System:** Users can comment on blog posts.
- **Like/Dislike System:** Users can like or dislike posts.
- **Advanced Search:** Search blogs by view count, likes, dislikes, and tags.
- **Role-Based Authorization:** User and admin roles with different access levels.
- **User Profile Management:** Comprehensive profile management for users.

## Tech Stack

- **Go**: Main programming language.
- **Gin**: Web framework for building APIs.
- **MongoDB**: NoSQL database for data storage.
- **Redis**: In-memory database for caching and rate limiting.
- **Docker**: Containerization platform for deployment.
- **Gemini AI**: AI integration for content generation.
- **Cloudinary**: Service for managing and storing images.

## Project Structure

```bash
.
├── cmd
│   └── main.go
├── Config
├── Delivery
│   ├── Controllers
│   ├── middlewares
│   └── Routers
├── Domain
├── Infrastructure
├── Mocks
├── Repository
├── Tests
├── UseCases
├── Utils
└── tmp
```

## Getting Started

### Prerequisites

- Go 1.18+
- Docker 20.10+
- MongoDB 5.0+
- Redis 6.2+

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/yourusername/blog-starter-backend.git
   cd blog-starter-backend
   ```

2. **Set up environment variables:**
   Create a `.env` file in the root directory and add your environment variables as shown below:

   ```env
   DB_NAME=StarterProject
   JWT_SECRET=your_jwt_secret
   SMTP_SERVER=smtp.gmail.com
   SMTP_PORT=465
   SMTP_USERNAME=your_email@gmail.com
   SMTP_PASSWORD=your_password
   BASE_URL=http://localhost:8081
   MONGO_URI=mongodb://localhost:27017
   SERVER_ADDRESS=localhost:8081
   CONTEXT_TIMEOUT=24
   ACCESS_TOKEN_EXPIRY_HOUR=24
   REFRESH_TOKEN_EXPIRY_HOUR=24
   REDIS_BLOG_KEY=blogs
   REDIS_DB_ADDRESS=localhost:6379
   GEMINI_API_KEY=your_gemini_api_key
   OAUTH_CLIENT_ID=your_oauth_client_id
   OAUTH_CLIENT_SECRET=your_oauth_client_secret
   OAUTH_REDIRECT_URL=http://localhost:8081/auth/callback
   CLOUDINARY_CLOUD_NAME=your_cloudinary_cloud_name
   CLOUDINARY_API_KEY=your_cloudinary_api_key
   CLOUDINARY_API_SECRET=your_cloudinary_api_secret
   ```

3. **Run the application using Docker:**
   ```bash
   docker-compose up --build
   ```

### Usage
- **Postman Documentation**: You can view API documentation at `https://documenter.getpostman.com/view/32287741/2sAXjGcEHS`.

## Testing

Run tests using the following command:

```bash
go test ./...
```

## Contributing

Contributions are welcome! Please fork the repository and create a pull request.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/NewFeature`)
3. Commit your Changes (`git commit -m 'Add some NewFeature'`)
4. Push to the Branch (`git push origin feature/NewFeature`)
5. Open a Pull Request


