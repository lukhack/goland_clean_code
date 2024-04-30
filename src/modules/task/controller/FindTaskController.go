package task

import (
	config "app/src/common/config"
	r "app/src/common/responses"
	usecase "app/src/modules/task/use_cases"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"strconv"
)

type FindTaskController struct {
	UseCase *usecase.FindTaskUseCase
	Result  *r.Result
}

func NewFindTaskController(usecase *usecase.FindTaskUseCase, r *r.Result) *FindTaskController {
	return &FindTaskController{
		UseCase: usecase,
		Result:  r,
	}
}

func (ph *FindTaskController) validateRequest(c fiber.Ctx) (id int, err error) {
	ids := c.Params("id")
	log.Info("id: %v", ids)
	id, err = strconv.Atoi(ids)

	if err != nil {
		return 0, fmt.Errorf(config.BAD_REQUES, "Id no valido "+ids)
	}

	return
}

func (ph *FindTaskController) Run(c fiber.Ctx) (err error) {
	id, err := ph.validateRequest(c)
	if err != nil {
		return ph.Result.BadRequest(c, "Error at validation "+err.Error())
	}

	task, err := ph.UseCase.Execute(id)
	if err != nil {
		return ph.Result.BadRequest(c, "Error at execution "+err.Error())
	}

	_ = ph.Result.Ok(c, task)
	return
}
