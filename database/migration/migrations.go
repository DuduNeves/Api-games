package migration

import (
	"github.com/DuduNeves/dev-games.git/models"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	db.AutoMigrate(models.Games{})
}
