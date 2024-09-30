package bootstrap

import (
	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2/log"

	"github.com/MikelSot/tribal-training-auth/infrastructure/handler"
	"github.com/MikelSot/tribal-training-auth/model"
)

func Run() {
	_ = godotenv.Load()

	app := newFiber()
	db := getConnectionDB()

	handler.InitRoutes(model.RouterSpecification{
		App:       app,
		DB:        db,
		ExpiresAt: getExpiresAtHours(),
		SignKey:   getSignKey(),
	})

	log.Fatal(app.Listen(getPort()))
}
