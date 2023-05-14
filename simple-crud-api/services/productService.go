package services

import (
	"github.com/salindae/simple-crud-api/initializers"
	"github.com/salindae/simple-crud-api/models"
)

func GetAllProducts() ([]models.Product, error) {
	products := []models.Product{}
	results := initializers.DB.Limit(10).Find(&products)
	return products, results.Error
}
func GetProductById(id int) (models.Product, error) {
	product := models.Product{}
	results := initializers.DB.First(&product, id)
	return product, results.Error
}
func DeleteProductById(id int) error {
	results := initializers.DB.Delete(&models.Product{}, id)
	return results.Error
}
func UpdateProductById(id int, modifiedProduct models.Product) (models.Product, error) {
	product := models.Product{}
	results := initializers.DB.First(&product, id)
	product.Name = modifiedProduct.Name
	product.Category = modifiedProduct.Category
	product.Description = modifiedProduct.Description
	product.Price = modifiedProduct.Price
	initializers.DB.Save(&product)
	return product, results.Error
}
func AddProductById(product *models.Product) error {
	results := initializers.DB.Create(product)
	return results.Error
}
