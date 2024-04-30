package task

type TaskDto struct {
	Body string `json:"body" validate:"required"`
}
