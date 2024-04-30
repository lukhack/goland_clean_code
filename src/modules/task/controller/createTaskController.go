package task

import (
	config "app/src/common/config"
	r "app/src/common/responses"
	dto "app/src/modules/task/domain/dto"
	usecase "app/src/modules/task/use_cases"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type CreateTaskController struct {
	UseCase *usecase.CreateTaskUseCase
	Result  *r.Result
}

func NewCreateTaskController(usecase *usecase.CreateTaskUseCase, r *r.Result) *CreateTaskController {
	return &CreateTaskController{
		UseCase: usecase,
		Result:  r,
	}
}

func (ph *CreateTaskController) validateRequest(c fiber.Ctx) (task dto.TaskDto, err error) {
	if err = c.Bind().Body(&task); err != nil {
		return task, fmt.Errorf(config.BAD_REQUES, err.Error())
	}

	if err = validator.New().Struct(&task); err != nil {
		return task, fmt.Errorf(config.BAD_REQUES, err.Error())
	}

	return
}

func (ph *CreateTaskController) Run(c fiber.Ctx) (err error) {
	data, err := ph.validateRequest(c)
	if err != nil {
		return ph.Result.BadRequest(c, "Error at validation "+err.Error())
	}

	err = ph.UseCase.Execute(&data)
	if err != nil {
		return ph.Result.BadRequest(c, "Error at execution "+err.Error())
	}

	_ = ph.Result.Ok(c, data)
	return
}
