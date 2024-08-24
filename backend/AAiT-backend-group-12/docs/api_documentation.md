# Getting Started

Welcome to the documentation for the Blog API. This document describes routes and features of the API, their authorization level and all the necessary details one needs to run, build and use the application.

> Note: All the commands listed here are supposed to be executed in the __root directory of the project__. 

To facilitate API testing and continuous development with updated documents, [Postman](https://www.postman.com/) has been used to document the API endpoints. The details of the API are presented with their descriptions, authorization levels, sample request and response bodies and relevant examples.

- https://documenter.getpostman.com/view/37574343/2sAXjF6thE

## Quickstart
To run the application locally, use the following command:
```bash
go run ./delivery
```

To build the project, use the following command:
```bash
go build -o ./path/filename ./delivery
```
The build can then be executed using:
```bash
./path/filename 
```

## Features
### Authentication and Authorization
- Signup using a unique username, a unique email and a password.
- Signup using Google OAuth. The unique username and password is required to complete the oauth signup process.
- Login using either the email or password A
- Email verification for users that signup using the standard procedure.
- Password reset via email for users who have verified accounts.
- Authorization system based on refresh and access tokens. The refresh tokens are hashed and stored in the database to enable server-side logouts.
- Forced logouts.

### Users
- Support for profile pictures. Accounts that use google OAuth to signup will have their profile picture from their google account by default.
- Update user details. Updates to a user's bio, phone number and profile picture are supported by the system.
- User promotion by `Admin` roles.
- Admin demotion by the `Root` user.

### Blog
- Users can create blog posts.
- Users can update the contents of previously created blog posts.
- Users can delete blog posts.
- Users can retrieve blog posts based on a wide variety of filter metrics. These metrics are:
    - __Searching__
        - Search by title
        - Search by author
    - __Filtering__
        - Filter by tags
        - Filter by date (before/after)
        - Filter by minimum likes
        - Filter by minimum dislikes
        - Filter by minimum comment count
        - Filter by minimum view count
    - __Sorting__ (asc/dsc)
        - Sort by creation date
        - Sort by popularity metrics
            - Like count
            - Dislike count
            - View count
- The retrieval functions naturally support pagination. The definitions of the paging system, meaning the `document per page` and the `page number`, can be provided for usecases that require a dynamic pagination system.

### AI Integration
Users can use AI to supercharge their blogs in three manners:
- Use the LLM to read and modify a completed blog. This may be used to proof-read and potentially add extra content to an existing blog.
- Use the LLM to generate content for a provided title.
- Use the LLM to get relevant, trending content ideas for a particular field of study.

Accordingly, the blog API provides these AI integration features:

- `suggest improvement`: This endpoint takes the content and title of a blog post and performs improvements, returning the enhanced data.
- `create content`: This endpoint takes the content and title of a blog post and generates the blog content.
- `generate content ideas`: This endpoint queries for a good blog topic. 

### Comments
- Users can add comments to a blog site.
- Users can update previously submitted comments.
- Users can delete their comments.

## Development
In the root folder, config files for [air](https://github.com/air-verse/air) and [fresh](https://github.com/zzwx/fresh) can be found. These packages allow the application to re-build and re-launch whenever changes to the tracked files are made. To use fresh, for instance, download the package on your local machine and execute the `fresh` command in the root of the project.
```bash
go install github.com/zzwx/fresh@latest
```
In the root of the project:
```bash
fresh
```

## Class Diagram
![alt text](<./Class Diagram.drawio.svg>)

## Architecture and Folder Structure 
The project follows clean architecture principles, using interfaces defined in domains for dependency inversion. The folder structure is as follows:

```
root
|   .air.toml
|   .env
|   .fresh.yaml
|   .gitignore
|   go.mod
|   go.sum
|   README.md
|   Dockerfile
|   sample.env
|
+---delivery
|   |   main.go
|   |
|   +---controllers
|   |       auth_controller.go
|   |       blog_controller.go
|   |       comment_controller.go
|   |       oauth_controller.go
|   |
|   +---env
|   |       env.go
|   |
|   \---router
|           auth_router.go
|           blog_router.go
|           oauth_router.go
|           router.go
|
+---docs
|       api_documentation.md
|       Class Diagram.drawio
|
+---domain
|   |   domain.go
|   |   errors.go
|   |   services.go
|   |
|   \---dtos
|           blog_dtos.go
|           user_dtos.go
|
+---infrastructure
|   +---ai
|   |       ai_service.go
|   |
|   +---cryptography
|   |       hashing_service.go
|   |
|   +---db
|   |       config.go
|   |       root_user.go
|   |
|   +---fs
|   |       fs.go
|   |
|   +---jwt
|   |       jwt_service.go
|   |
|   +---mail
|   |       mail_service.go
|   |
|   +---middleware
|   |       auth_middleware.go
|   |
|   +---oauth
|   |       google_auth.go
|   |
|   +---redis
|   |       redis_service.go
|   |
|   \---utils
|           generate_token.go
|           validate_blogs.go
|
+---mocks
|       AIModelInterface.go
|       AIServicesInterface.go
|       BlogRepositoryInterface.go
|       BlogUseCaseInterface.go
|       CacheRepositoryInterface.go
|       CodedError.go
|       HashingServiceInterface.go
|       JWTServiceInterface.go
|       MailServiceInterface.go
|       UserRepositoryInterface.go
|       UserUsecaseInterface.go
|
+---repository
|       blog_repository.go
|       cache_repository.go
|       user_repository.go
|
+---tests
|       ai_service_test.go
|       blog_controller_test.go
|       blog_repository_test.go
|       blog_usecase_test.go
|       cache_repository_test.go
|       comment_controller_test.go
|       hashing_service_test.go
|       jwt_service_test.go
|       mail_service_test.go
|       middleware_test.go
|       oauth_service_test.go
|       user_repository_test.go
|       user_usecase_test.go
|       utils_test.go
|
\---usecase
        blog_usecase.go
        user_usecase.go
```

### Application Components

> Delivery: Contains files related to the delivery layer, handling incoming requests and responses, including the controllers and routers. This directory also contains the entry point to the application and the setup file for loading environment variables.

- `main.go` Sets up the HTTP server, initializes dependencies, and defines the routing configuration.

- `controllers/` Contains the handlers for incoming HTTP requests and invokes the appropriate use case methods.

- `routers/` Sets up the routes and initializes the Gin router. `router.go` call all the other router files and is responsible for initalizing the API endpoints.

> Domain: Defines the core business entities, relevant data transfer objects and the interfaces for the various layers of the application.

- `domain.go` Defines the core business entities, the interface definitions for controllers, usecases and repositories.

- `errors.go` Defines a custom error model used to communicate errors throughout the application.

- `services.go` Defines interfaces for the services used by the application.

- `dtos/` Defines the data transfer objects used by the application with all the infrastructure-specific tags.

> Infrastructure: Implements external dependencies and services that support the application functionality.

- `ai/ai_service.go` Contains the implementation of AI services used in the application. This includes functions for utilizing the LLM to read and modify completed blogs, generate content for a provided title, and get relevant, trending content ideas for a particular field of study.

- `cryptography/hashing_service.go` Contains functions for hashing strings and comparing and validating previously hashed strings with against their plaintext candidates.

- `db/config.go` Contains the functions used to connect to a MongoDB cluster, create all the indices in the DB and disconnect from the previously connected cluster.

- `db/root_user.go` Contains the function call for creating the root user of the application.

- `fs/` Contains all the filesystem operations used by the application.

- `jwt/jwt_service.go` Contains functions for generating and validating JWT tokens, the primary authentication system used in this project.

- `mail/mail_service.go` Contains the logic for sending emails using a specified email server. It also provides methods for composing mails using specified templated.

- `middleware/auth_middleWare.go` Contains the middleware responsible for handling authentication and authorization using JWT tokens. It ensures that only authenticated users can access certain routes or perform specific actions.


- `oauth/google_auth.go` Contains the implementation of Google OAuth authentication. It provides functions for authenticating users using Google credentials.

- `redis/redis_service.go` Contains the functions used to connect to a Redis store and disconnect from the previously connected redis client.

- `utils/` Defines the miscellaneous functions used by the application. Their purposes range from adding extra validation for certain fields to generating random tokens.


> Repositories: Encapsulate data access logic and abstracts the database technology and functionality from the usecases.

- `blog_repository.go`: Interface implementation for blog data access operations.

- `cache_repository.go`: Interface implementation for cache data access operations.

- `user_repository.go`: Interface implementation for user data access operations.

> Usecases: Contains the application-specific business rules and logic. A complete list of the features of the application is detailed in the ___Features___ section.

- `blog_usecases.go`: Implements the usecases related to blogs, such as creating, updating, retrieving, and deleting tasks.
- `user_usecases.go`: Implements the usecases related to users, such as registering, logging in, password reset, etc...

### Testing Components
- `tests/`: Contains all the unit tests for the various components of the application

- `mocks/`: Contains all the mocked components used in the tests.

## Enviornment Variables
As usual, the environment variables are defined in a `.env` file at the root of the project. These files are processed using functions in `delivery/env/env.go` before the application runs. It is advised to simply copy the `sample.env` file located at the root of the project, rename the file to `.env` and fill in the values with your credentials. Incase the `sample.env` file isn't available, use this template for defining your environment variables.
```
DB_ADDRESS=
DB_NAME=
TEST_DB_NAME=
PORT=
JWT_SECRET_TOKEN=
ACCESS_TOKEN_LIFESPAN_MINUTES=
REFRESH_TOKEN_LIFESPAN_HOURS=
ROOT_USERNAME=
ROOT_PASSWORD=
ROUTE_PREFIX=
SMTP_GMAIL=
SMTP_PASSWORD=
REDIS_URL=
GOOGLE_CLIENT_ID=
GOOGLE_CLIENT_SECRET=
GEMINI_API_KEY=
CACHE_EXPIRATION=
```
Note: The `SMTP_GMAIL` and `SMTP_PASSWORD` variables are used for sending the mail using the `smpt.gmail.com`. To use a different provider, you have to change the code located in `mail_service.go`. The same files contains HTML templates for generating the mail that will be sent to users as well. Modification to those templates will change how the email verification and password request emails look like.

The `TEST_DB_NAME` variable will be used by the repository during testing. The tests will use the same database cluster defined under `DB_ADDRESS`. Not providing a name for the test database will cause the repository tests to fail.

The same redis store will be used for the cache repository tests. Again, failing to provide the redis connection url will cause the tests to fail.

## Sending requests using tokens

The authentication system is based on JWT. The token will be sent to the client when it makes a request to the login endpoint (or the oauth login endpoint) with the correct credentials. That token must be included in the Authorization header of any requests to protected routes. The format of the token follows the standard `bearer e...` format. Any deviation from this might cause the middleware to block the incoming request.

Sample request (with auth header):
```bash
curl --location --request DELETE 'http://localhost:8080/protected_endpoint' \
--header 'Authorization: bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Imt5c2sifQ.pIb58jAfa9Rd3u38AzTLdtU_hGR624P6by2epR_baMM'
```

## Testing

Note: The repository tests WILL NOT pass if a valid test database has not been setup. Make sure to check the environment variables section for more information on how to provide a test DB.

To run all the tests:
```bash 
go test -v blog_api/tests
```
Each file contains a suite that groups up all the related tests. To run one of suites contained in one of the files, run this command:
```bash
go test -run NAME_OF_THE_SUITE ./tests/file_name_test.go 

# example
go test -run TestUserUsecase ./tests/user_usecase_test.go
```
With timeout:
```bash
go test -timeout 30s -run NAME_OF_THE_SUITE ./tests/file_name_test.go 
```

The suites are usually run using the functions defined last, accepting a `*testing.T` as a parameter and running the test suite.

Coverage profiles can also be created. To run tests and generate the corresponding coverage file:
```bash
go test -v -coverprofile="OUTPUT" -coverpkg="PACKAGE"  ./tests/FILENAME

# example: repository coverage
go test -v -coverprofile="cover.out" -coverpkg="blog_api/repository"  ./tests/blog_repository_test.go ./tests/user_repository_test.go ./tests/cache_repository_test.go

# example: jwt service coverage
go test -v -coverprofile="cover.out" -coverpkg="blog_api/infrastructure/jwt"  ./tests/jwt_service_test.go

# example: usecase coverage
go test -v -coverprofile="cover.out" -coverpkg="blog_api/usecase"  ./tests/blog_usecase_test.go ./tests/user_usecase_test.go

# example: controller coverage
go test -v -coverprofile="cover.out" -coverpkg="blog_api/delivery/controllers"  ./tests/auth_controller_test.go ./tests/oauth_controller_test.go ./tests/blog_controller_test.go ./tests/comment_controller_test.go
```

The output of the commands above are not human-readable. To generate a `.html` file with all the data obtained from the coverage profile, use the following command:
```bash
go tool cover -html="cover.out" -o coverage.html
```
