package infraestruture

import (
	db "app/src/infraestructure/db/adapter"
	entities "app/src/modules/task/domain/entities"
	"gorm.io/gorm"
)

type PostgreTaskDao struct {
	sql *gorm.DB
}

func NewPostgreTaskDao(connection *db.DBConnections) *PostgreTaskDao {
	return &PostgreTaskDao{sql: connection.DB}
}

func (dao *PostgreTaskDao) Create(t *entities.Task) (*entities.Task, error) {
	dbs := dao.sql
	err := dbs.Create(&t).Error
	return t, err
}

func (dao *PostgreTaskDao) FindAll(page int, limit int) (*[]entities.Task, error) {
	dbs := dao.sql
	var tasks []entities.Task
	err := dbs.Find(&tasks).Limit(page).Offset(limit).Error
	return &tasks, err
}

func (dao *PostgreTaskDao) FindById(id int) (task *entities.Task, err error) {
	dbs := dao.sql
	err = dbs.First(&task, id).Error
	return
}

func (dao *PostgreTaskDao) EditTasks(t *entities.Task, id int) (*entities.Task, error) {
	dbs := dao.sql
	task, err := dao.FindById(id)
	if err != nil {
		return nil, err
	}
	task.Body = t.Body
	dbs.Save(&task)
	return task, nil
}

func (dao *PostgreTaskDao) DeleteTasks(id int) error {
	dbs := dao.sql
	task, err := dao.FindById(id)
	if err != nil {
		return err
	}

	dbs.Delete(&task)
	return nil
}
