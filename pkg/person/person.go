// Package core containts the core domain of the domain model.
package core

// Person defines a person. What a weird description
type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// NewPerson creates a new instance of person.
func NewPerson(id int, name string) *Person {
	return &Person{
		ID:   id,
		Name: name,
	}
}

// PersonRepository provides access a person store.
type PersonRepository interface {
	Create(person *Person) error
	Find(id int) (*Person, error)
	FindAll() []*Person
}
