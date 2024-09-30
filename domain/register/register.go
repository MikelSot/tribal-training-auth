package register

import "github.com/MikelSot/tribal-training-auth/model"

type UseCase interface {
	Register(m model.User) (interface{}, error)
}

type UserUseCase interface {
	Create(m model.User) (model.User, error)
}

type TokenUseCase interface {
	Generate(m model.User) (string, error)
}
