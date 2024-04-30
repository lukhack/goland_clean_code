package use_cases

import (
	dao "app/src/infraestructure/db/dao"
	dto "app/src/modules/task/domain/dto"
	entities "app/src/modules/task/domain/entities"
	repository "app/src/modules/task/domain/repositories"
	"github.com/gofiber/fiber/v3/log"
)

type CreateTaskUseCase struct {
	repo repository.TaskRepository
}

func NewCreateTaskUseCase(repo *dao.PostgreTaskDao) *CreateTaskUseCase {
	return &CreateTaskUseCase{repo: repo}
}

func (puc *CreateTaskUseCase) Execute(data *dto.TaskDto) (err error) {
	task, err := puc.repo.Create(&entities.Task{
		Body: data.Body,
	})

	log.Info("Tarea created: ", task)
	return
}
