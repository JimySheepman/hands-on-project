package controllers

import (
	"shortest-distances/service"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService service.UserService
}

func NewController(service service.UserService) *UserController {
	return &UserController{
		userService: service,
	}
}

func (uc *UserController) Healthcheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "Success",
		"message": "The system is healthy.",
	})
}

func (uc *UserController) Login(c *fiber.Ctx) {

}

func (uc *UserController) Register(c *fiber.Ctx) {

}

/*
func (rc *RepoController) SelectShippingCompany(c *fiber.Ctx) error {
	var deliveries models.Deliveries

	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := c.BodyParser(&deliveries); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	shipments, err := rc.repoService.SelectShippingCompany(customContext, &deliveries)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"shipments": shipments,
	})
}
*/
