package initializers

import (
	"github.com/robinmuhia/GoJWTAuth/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}