package task

import (
	config "app/src/common/config"
	r "app/src/common/responses"
	usecase "app/src/modules/task/use_cases"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

type FindAllTaskController struct {
	UseCase *usecase.FindAllTaskUseCase
	Result  *r.Result
}

func NewFindAllTaskController(usecase *usecase.FindAllTaskUseCase, r *r.Result) *FindAllTaskController {
	return &FindAllTaskController{
		UseCase: usecase,
		Result:  r,
	}
}

func (ph *FindAllTaskController) validateRequest(c fiber.Ctx) (page int, limit int, err error) {
	l := c.Params("limit")
	p := c.Params("limit")

	page, err = strconv.Atoi(p)
	limit, err = strconv.Atoi(l)
	if err != nil {
		return 0, 0, fmt.Errorf(config.BAD_REQUES, "Pages and limit are required")
	}

	return
}

func (ph *FindAllTaskController) Run(c fiber.Ctx) (err error) {
	page, limit, err := ph.validateRequest(c)
	if err != nil {
		return ph.Result.BadRequest(c, "Error at validation "+err.Error())
	}

	task, err := ph.UseCase.Execute(page, limit)
	if err != nil {
		return ph.Result.BadRequest(c, "Error at execution "+err.Error())
	}

	_ = ph.Result.Ok(c, task)
	return
}
