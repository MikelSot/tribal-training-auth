package handler

import (
	"github.com/MikelSot/tribal-training-auth/infrastructure/handler/login"
	"github.com/MikelSot/tribal-training-auth/infrastructure/handler/register"
	"github.com/MikelSot/tribal-training-auth/model"
)

func InitRoutes(spec model.RouterSpecification) {
	// L
	login.NewRouter(spec)

	// R
	register.NewRouter(spec)
}
