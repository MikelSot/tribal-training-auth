package login

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"

	"github.com/MikelSot/tribal-training-auth/domain/login"
	"github.com/MikelSot/tribal-training-auth/model"
)

type handler struct {
	useCase login.UseCase
}

func newHandler(uc login.UseCase) handler {
	return handler{useCase: uc}
}

func (h handler) Login(c *fiber.Ctx) error {
	m := model.User{}

	if err := c.BodyParser(&m); err != nil {
		log.Warn("¡Uy! Error al leer el cuerpo de la solicitud", err.Error())

		return c.Status(fiber.StatusBadRequest).JSON(`{"errors": [{"code": "bind_failed", "message": "Error al leer el cuerpo de la solicitud"}]}`)
	}

	data, err := h.useCase.Login(m)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(data)
}
