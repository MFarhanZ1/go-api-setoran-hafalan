package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mfarhanz1/go-api-setoran-hafalan/app/services"
	"github.com/mfarhanz1/go-api-setoran-hafalan/app/types"
)

type GlobalHandler struct {
	service *services.GlobalService
}

func (gh *GlobalHandler) Introduce(c fiber.Ctx) error {
	introduce, err := gh.service.Introduce()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&types.CommonHandlerType{
			Response: false,
			Message:  err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(introduce)
}

func (gh *GlobalHandler) NotFound(c fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(types.CommonHandlerType{
		Response: false,
		Message:  "Waduh, mau nyari apa kamu mas? 😅",
	})
}

func NewGlobalHandler(service *services.GlobalService) *GlobalHandler {
	return &GlobalHandler{service: service}
}