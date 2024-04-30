package use_cases

import (
	dao "app/src/infraestructure/db/dao"
	entities "app/src/modules/task/domain/entities"
	repository "app/src/modules/task/domain/repositories"
	"github.com/gofiber/fiber/v3/log"
)

type FindTaskUseCase struct {
	repo repository.TaskRepository
}

func NewFindTaskUseCase(repo *dao.PostgreTaskDao) *FindTaskUseCase {
	return &FindTaskUseCase{repo: repo}
}

func (puc *FindTaskUseCase) Execute(id int) (task *entities.Task, err error) {
	task, err = puc.repo.FindById(id)
	log.Info("Tarea Findd: ", task)
	return
}
