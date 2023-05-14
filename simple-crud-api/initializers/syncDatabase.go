package initializers

import "github.com/salindae/simple-crud-api/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.Product{})
}
