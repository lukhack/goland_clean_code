package common

type PaginationDto struct {
	Page  int `json:"page" validate:"required"`
	Limit int `json:"limit" validate:"required"`
}
