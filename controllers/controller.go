package controllers

import (
	"apiantrean/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error {
	var antrean []models.Antrean

	models.DB.Find(&antrean)

	return c.Status(fiber.StatusOK).JSON(antrean)
}

func Show(c *fiber.Ctx) error{

	id := c.Params("id")

	var antrean models.Antrean

	if err := models.DB.Model(&antrean).Where("id = ?", id).First(&antrean).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Id tidak ditemukan",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Server sedang mengalami gangguan",
		})
	}

	return c.JSON(antrean)

}

func Create(c *fiber.Ctx) error {
	var antrean models.Antrean

	if err := c.BodyParser(&antrean); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	antrean.Id = "1"

	if err := models.DB.Create(&antrean).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Server sedang mengalami gangguan",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil membuat data antrean",
	})
}

func Update(c *fiber.Ctx) error {
	id := c.Params("id")

	var antrean models.Antrean

	if err := c.BodyParser(&antrean); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if models.DB.Where("id = ?", id).Updates(&antrean).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Id tidak ditemukan",
		})
	}

	return c.JSON(fiber.Map{
		"message" : "Berhasil memperbaharui data",
	})
}