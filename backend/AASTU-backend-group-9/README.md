# Blog Project

Welcome to the BLog API Project! This repository contains the backend Api for Blog project, developed by multiple groups. The project aims to provide a comprehensive solution for managing blogs and user interactions.

## Table of Contents

1. [Introduction](#introduction)
2. [Getting Started](#getting-started)

## Introduction

The Blog Starter Project is designed to develop a backend API for a blog platform. This platform allows users to create, read, update, and delete blog posts, manage profiles, perform advanced searches, and leverage AI for content suggestions. The API also integrates user authentication, authorization, and role management, ensuring secure and efficient operations.

## Getting Started

### Prerequisites

Before you begin, ensure you have the following installed on your machine:

- [Go](https://golang.org/doc/install) (version 1.16 or higher)
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