package user

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/AJRDRGZ/db-query-builder/models"
	"golang.org/x/crypto/bcrypt"

	"github.com/MikelSot/tribal-training-auth/model"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

const _defaultMinLenPassword = 5

type User struct {
	storage Storage
}

func New(s Storage) User {
	return User{s}
}

func (u User) Create(m model.User) (model.User, error) {
	if err := model.ValidateStructNil(m); err != nil {
		return model.User{}, fmt.Errorf("user: %w", err)
	}

	m.Email = strings.ToLower(m.Email)

	if !emailRegex.MatchString(m.Email) {
		return model.User{}, fmt.Errorf("upps! el email no es válido")
	}

	if len(m.Password) < _defaultMinLenPassword {
		return model.User{}, fmt.Errorf("upps! la contraseña no es válida")
	}

	if strings.TrimSpace(m.FirstName) == "" || strings.TrimSpace(m.Lastname) == "" || strings.TrimSpace(m.Email) == "" {
		return model.User{}, fmt.Errorf("upp! los datos no son válidos")
	}

	passwordBcrypt, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, fmt.Errorf("upps ocurrió un error inesperado")
	}
	m.Password = string(passwordBcrypt)

	err = u.storage.Create(m)
	if err != nil {
		return model.User{}, fmt.Errorf("user: %w", err)
	}
	m.Password = ""

	return m, nil
}

func (u User) GetByEmail(email string) (model.User, error) {
	if email == "" {
		return model.User{}, fmt.Errorf("user: email is required")
	}

	user, err := u.storage.GetWhere(models.FieldsSpecification{
		Filters: models.Fields{{Name: "email", Value: email}},
	})
	if err != nil {
		return model.User{}, fmt.Errorf("user: %w", err)
	}

	return user, nil
}
