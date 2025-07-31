package registry

import (
	"internship-mini-project/internal/domain"
)

type Model struct {
	Model any
}

func RegistryModels() []Model {
	return []Model{
		{Model: domain.User{}},
		{Model: domain.Toko{}},
		{Model: domain.Category{}},
		{Model: domain.Address{}},
	}
}
