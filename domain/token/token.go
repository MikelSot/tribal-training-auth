package token

import "github.com/MikelSot/tribal-training-auth/model"

type UseCase interface {
	Generate(m model.User) (string, error)
}
