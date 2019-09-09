package adding

import (
	"errors"

	"github.com/joshuabezaleel/cats-ddd/pkg/core"
)

// ErrDuplicate is used when a cat already exists.
var ErrDuplicate = errors.New("Cat already exists")

// Service provides cat adding operations.
type Service interface {
	AddCat(cat *Cat) error
	AddPerson(person *Person) error
}

type service struct {
	cats   core.CatRepository
	people core.PersonRepository
}

func NewService(cats core.CatRepository, people core.PersonRepository) Service {
	return &service{
		cats:   cats,
		people: people,
	}
}

func (s *service) AddCat(cat *Cat) error {
	return nil
}

func (s *service) AddPerson(person *Person) error {
	return nil
}

type Cat struct {
}

type Person struct {
}
