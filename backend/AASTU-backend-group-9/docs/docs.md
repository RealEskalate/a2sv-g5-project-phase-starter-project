# Blog Starter Project API Documentation

## Introduction

### Overview
The Blog Starter Project is a backend API designed to power a blog platform. This platform allows users to create, read, update, and delete blog posts, manage user profiles, and perform advanced search and filtering operations. Key features include user authentication, authorization, role management, and AI integration for content suggestions.

### Product Goals and Objectives
- Developed a RESTful API with clear and intuitive endpoints.
- Implemented core CRUD operations for blog posts.
- Integrateed user authentication and authorization mechanisms.
- Provideed functionalities for blog tags, filtration, and search.
- Incorporateed AI capabilities for generating blog content based on user input.
- Ensured high performance, reliability, and scalability.

## Functional Requirements

### 1. User Management

#### 1.1 User Registration
- **Description:** Users can register on the platform with their email, password, and profile details.
- **Actors:** Visitor
- **Flow of Events:**
  1. The user submits necessary details (username, email, password, etc.).
  2. The system validates the information (email format, password strength, etc.).
  3. The system checks for existing users with the same email or username.
  4. If validation passes, a new user account is created.
  5. The system may send an activation email or link for account verification.

#### 1.2 Login
- **Description:** Registered users can log in to access the platform's features.
- **Actors:** User & Admin
- **Flow of Events:**
  1. The user provides their username (or email) and password.
  2. The system validates the credentials.
  3. If valid, the system generates and returns access and refresh tokens.

#### 1.3 Authentication
- **Description:** Allows users to access the system without needing to log in every time.
- **Actors:** User & Admin
- **Flow of Events:**
  1. The user sends the access token with each request.
  2. The server verifies the token.
  3. If expired, the system validates the refresh token and issues a new access token.

#### 1.4 Forgot Password
- **Description:** Users can reset their password if forgotten.
- **Actors:** User & Admin
- **Flow of Events:**
  1. User requests a password reset by providing their email.
  2. The system sends a password reset link via email.
  3. The user resets the password using the link.
  4. The system updates the password and confirms the change.

#### 1.5 Logout
- **Description:** Allows users to log out or be forcefully logged out.
- **Actors:** User & Admin
- **Flow of Events:**
  1. The user requests to log out or the system identifies expired tokens.
  2. The system invalidates the tokens and removes them from the database.
  3. A confirmation response is sent to the user.

#### 1.6 User Promotion and Demotion
- **Description:** Admins can promote a user to admin or demote them back to a regular user.
- **Actors:** Admin
- **Flow of Events:**
  1. Admin sends a request to promote or demote a user.
  2. The system verifies the admin's authorization.
  3. The system updates the user's role accordingly and returns a response.

### 2. Blog Management

#### 2.1 Blog Creation
- **Description:** Authenticated users can create new blog posts.
- **Actors:** User & Admin
- **Flow of Events:**
  1. User sends a request with blog details (title, content, tags, etc.).
  2. The system validates the input and authorization.
  3. If valid, the system creates a new blog post and returns the details.

#### 2.2 Blog Retrieval
- **Description:** Users can view blog posts with pagination and popularity metrics.
- **Actors:** User & Admin
- **Flow of Events:**
  1. User requests to retrieve blog posts with optional pagination and sorting.
  2. The system retrieves and returns the blog posts with associated metrics.

#### 2.3 Blog Update
- **Description:** Users can update the details of an existing blog post.
- **Actors:** User
- **Flow of Events:**
  1. User sends a request with updated blog details.
  2. The system validates the input and checks if the user is the author.
  3. If valid, the system updates the blog post and returns the updated details.

#### 2.4 Blog Deletion
- **Description:** Authenticated users or admins can delete blog posts.
- **Actors:** User & Admin
- **Flow of Events:**
  1. User sends a request to delete a blog post.
  2. The system validates the request and checks authorization.
  3. The blog post is deleted, and a confirmation message is returned.

#### 2.5 Blog Search
- **Description:** Users can search for blog posts based on titles or author names.
- **Actors:** User
- **Flow of Events:**
  1. User sends a request with search criteria.
  2. The system retrieves and returns matching blog posts.

#### 2.6 Blog Popularity Tracking
- **Description:** The system tracks popularity metrics for each blog post.
- **Actors:** System
- **Flow of Events:**
  1. The system tracks and updates view counts, likes, and comments.
  2. Duplicate actions are prevented to ensure accurate tracking.

#### 2.7 Blog Filtration
- **Description:** Users can filter blog posts by tags, date, or popularity.
- **Actors:** User
- **Flow of Events:**
  1. User sends a request with filter criteria.
  2. The system retrieves and returns the filtered blog posts.

### 3. AI Integration

#### 3.1 Content Suggestions
- **Description:** The system suggests improvements or generates content ideas based on user-provided keywords or topics.
- **Actors:** User
- **Flow of Events:**
  1. User requests content suggestions.
  2. The system processes the input and returns AI-generated content suggestions.

### 4. Profile Management

#### 4.1 Profile Update
- **Description:** Users can update their profile details.
- **Actors:** User
- **Flow of Events:**
  1. User sends a request with updated profile details.
  2. The system validates and updates the user's profile.
  3. The updated profile details are returned.

## API Documentation

This API documentation provides a detailed guide on how to interact with the Blog Starter Project API. All endpoints are documented with example requests and responses, error codes, and descriptions.

https://documenter.getpostman.com/view/36737395/2sAXjDevNy