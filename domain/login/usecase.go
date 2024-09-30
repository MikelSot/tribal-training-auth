package login

import (
	"database/sql"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/MikelSot/tribal-training-auth/model"
)

type Login struct {
	user  UserUseCase
	token TokenUseCase
}

func New(user UserUseCase, token TokenUseCase) Login {
	return Login{user, token}
}

func (l Login) Login(m model.User) (interface{}, error) {
	user, err := l.user.GetByEmail(m.Email)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("user not found")
	}
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(m.Password))
	if err != nil {
		return nil, fmt.Errorf("invalid password")
	}
	m.Password = ""
	user.Password = ""

	token, err := l.token.Generate(m)
	if err != nil {
		return nil, fmt.Errorf("unexpected error")
	}

	mr := struct {
		model.User `json:"user"`
		Token      string `json:"token"`
	}{
		user,
		token,
	}

	return mr, nil
}
