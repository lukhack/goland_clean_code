package taks

import entities "app/src/modules/task/domain/entities"

type TaskRepository interface {
	FindAll(page int, limit int) (*[]entities.Task, error)
	Create(task *entities.Task) (*entities.Task, error)
	FindById(id int) (*entities.Task, error)
	EditTasks(t *entities.Task, id int) (*entities.Task, error)
	DeleteTasks(id int) error
}
