package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nicobellanich/go_postgreSQL_docker/database"
	"github.com/nicobellanich/go_postgreSQL_docker/models"
)

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Instance.Find(&facts)

	return c.Status(200).JSON(facts)
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Instance.Create(&fact)

	return c.Status(200).JSON(fact)
}
