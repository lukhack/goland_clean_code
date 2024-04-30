package task

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Body string
}
