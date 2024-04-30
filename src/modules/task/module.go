package task

import (
	types "app/src/common/types"
	dao "app/src/infraestructure/db/dao"
	controller "app/src/modules/task/controller"
	"app/src/modules/task/use_cases"
	"net/http"

	"go.uber.org/fx"
)

func configureModuleRoutes(
	ctrlCreateTask *controller.CreateTaskController,
	ctrlFindTask *controller.FindTaskController,
	ctrlFindAllTask *controller.FindAllTaskController,
	h *types.HandlersStore,
) {
	handlersModuleProducts := &types.SliceHandlers{
		Prefix: "task",
		Routes: []types.HandlerModule{
			{
				Route:   "/",
				Method:  http.MethodPost,
				Handler: ctrlCreateTask.Run,
			},
			{
				Route:   "/:id",
				Method:  http.MethodGet,
				Handler: ctrlFindTask.Run,
			},
			{
				Route:   "/page/:page/limit/:limit",
				Method:  http.MethodGet,
				Handler: ctrlFindAllTask.Run,
			},
		},
	}

	h.Handlers = append(h.Handlers, *handlersModuleProducts)

}

func ModuleProviders() []fx.Option {
	return []fx.Option{
		fx.Provide(dao.NewPostgreTaskDao),
		fx.Provide(controller.NewCreateTaskController),
		fx.Provide(controller.NewFindTaskController),
		fx.Provide(controller.NewFindAllTaskController),

		fx.Provide(use_cases.NewCreateTaskUseCase),
		fx.Provide(use_cases.NewFindTaskUseCase),
		fx.Provide(use_cases.NewFindAllTaskUseCase),
		fx.Invoke(configureModuleRoutes),
	}
}
