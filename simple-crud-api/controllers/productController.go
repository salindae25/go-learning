package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/salindae/simple-crud-api/models"
	"github.com/salindae/simple-crud-api/services"
)

func ProductRoutes(api fiber.Router) {
	api.Get("", getProduct)
	api.Get("/:id", getProductById)
	api.Delete("/:id", deleteProduct)
	api.Post("", postProduct)
	api.Put("/:id", updateProduct)
}
func getProduct(c *fiber.Ctx) error {
	result, err := services.GetAllProducts()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(result)
}
func getProductById(c *fiber.Ctx) error {
	id, errParam := c.ParamsInt("id")
	if errParam != nil {
		return c.Status(500).SendString(errParam.Error())
	}
	result, err := services.GetProductById(id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(result)
}
func postProduct(c *fiber.Ctx) error {
	p := new(models.Product)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	product := models.Product{
		Name:        p.Name,
		Description: p.Description,
		Category:    p.Category,
		Price:       p.Price,
	}
	err := services.AddProductById(&product)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(product)
}
func updateProduct(c *fiber.Ctx) error {

	id, errParam := c.ParamsInt("id")
	if errParam != nil {
		return c.Status(500).SendString(errParam.Error())
	}
	p := new(models.Product)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	result, err := services.UpdateProductById(id, models.Product{
		Name:        p.Name,
		Description: p.Description,
		Category:    p.Category,
		Price:       p.Price,
	})
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(result)
}
func deleteProduct(c *fiber.Ctx) error {

	id, errParam := c.ParamsInt("id")
	if errParam != nil {
		return c.Status(500).SendString(errParam.Error())
	}
	err := services.DeleteProductById(id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(`{"message":"Success"}`)
}
