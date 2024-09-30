package register

import (
	"fmt"
	"github.com/MikelSot/tribal-training-auth/model"
)

type Register struct {
	user  UserUseCase
	token TokenUseCase
}

func New(u UserUseCase, t TokenUseCase) Register {
	return Register{u, t}
}

func (r Register) Register(m model.User) (interface{}, error) {
	user, err := r.user.Create(m)
	if err != nil {
		return nil, err
	}

	token, err := r.token.Generate(m)
	if err != nil {
		return nil, fmt.Errorf("error generating token: %w", err)
	}

	user.Password = ""

	mr := struct {
		model.User `json:"user"`
		Token      string `json:"token"`
	}{
		user,
		token,
	}

	return mr, nil
}
