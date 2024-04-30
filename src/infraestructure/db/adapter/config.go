package infraestruture

import (
	"app/src/common/config"
	entities "app/src/modules/task/domain/entities"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConnections struct {
	DB *gorm.DB
}

func NewDBConnection(cfg *common.Config) *DBConnections {
	var err error
	var db *gorm.DB
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.UserName,
		cfg.DB.Password,
		cfg.DB.Database,
	)

	fmt.Println(psqlInfo)

	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("connected to database")

	//Migrar base de datos
	db.AutoMigrate(&entities.Task{})
	fmt.Println("database migrated")
	return &DBConnections{db}
}
