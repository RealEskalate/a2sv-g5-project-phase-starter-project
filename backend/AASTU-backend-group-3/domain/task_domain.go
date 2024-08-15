package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Status      string             `bson:"status" json:"status"`
	CreaterID   string             `bson:"creater_id" json:"creater_id"`
}

type TaskRepository interface {
	CreateTask(task Task) (Task, error)
	GetTasks() ([]Task, error)
	GetTaskByID(id string, creater string, Rol_ string) (Task, error)
	GetMyTasks(usename string) ([]Task, error)
	DeleteTask(id string) (Task, error)
	UpdateTask(id string, task Task) (Task, error)
	

}

type TaskUsecase interface {
	CreateTask(task Task) error
	GetTasks() ([]Task, error)
	GetTaskByID(id string, creater string, Role_ string) (Task, error)
	GetMyTasks(username string) ([]Task, error)
	DeleteTask(id string) error
	UpdateTask(id string, task Task) error
}

type TaskController interface {
	CreateTask(task Task) error
	GetTasks() ([]Task, error)
	GetTaskByID(id string, creater string, Role_ string) (Task, error)
	GetMyTasks(username string) ([]Task, error)
	DeleteTask(id string) error
	UpdateTask(id string, task Task) error
}


