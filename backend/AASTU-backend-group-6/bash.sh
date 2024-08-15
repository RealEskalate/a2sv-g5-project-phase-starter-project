
# Create directories for each component
mkdir -p Delivery/controllers
mkdir -p Delivery/routers
mkdir -p Domain
mkdir -p Infrastructure
mkdir -p Repositories
mkdir -p Usecases

# Create the Go files
touch Delivery/main.go
touch Delivery/controllers/controller.go
touch Delivery/routers/router.go
touch Domain/domain.go
touch Infrastructure/auth_middleWare.go
touch Infrastructure/jwt_service.go
touch Infrastructure/password_service.go
touch Repositories/task_repository.go
touch Repositories/user_repository.go
touch Usecases/task_usecases.go
touch Usecases/user_usecases.go

echo "Task Manager folder structure created successfully."
