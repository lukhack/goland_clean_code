package use_cases

import (
	dao "app/src/infraestructure/db/dao"
	entities "app/src/modules/task/domain/entities"
	repository "app/src/modules/task/domain/repositories"
	"github.com/gofiber/fiber/v3/log"
)

type FindAllTaskUseCase struct {
	repo repository.TaskRepository
}

func NewFindAllTaskUseCase(repo *dao.PostgreTaskDao) *FindAllTaskUseCase {
	return &FindAllTaskUseCase{repo: repo}
}

func (puc *FindAllTaskUseCase) Execute(page int, limit int) (tasks *[]entities.Task, err error) {
	tasks, err = puc.repo.FindAll(page, limit)
	log.Info("Tarea Findd: ", tasks)
	return
}
