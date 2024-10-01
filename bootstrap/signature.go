package bootstrap

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v4"

	"github.com/MikelSot/tribal-training-auth/infrastructure/handler/request"
)

func ValidateJWT(c *fiber.Ctx) error {

	tokenHeader, err := request.GetTokenFromHeader(c)
	if err != nil {
		log.Warn("Se encontró un error al tratar de leer el token")

		return fmt.Errorf("bootstrap: %w", err)
	}

	verifyFunction := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("método de firma no válido")
		}

		return []byte(getSignKey()), nil
	}

	token, err := jwt.Parse(tokenHeader, verifyFunction)
	if errors.Is(err, jwt.ErrTokenExpired) {
		log.Warn("Token de acceso expirado")

		return fmt.Errorf("token expirado")
	}
	if !token.Valid {
		log.Warn("Token de acceso no válido: %s", err.Error())

		return fmt.Errorf("token no válido")
	}
	if err != nil {
		log.Errorf("Error al procesar el token: %s", err.Error())

		return fmt.Errorf("error al procesar el token")
	}

	return c.Next()
}
