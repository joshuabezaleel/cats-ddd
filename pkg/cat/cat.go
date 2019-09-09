// Package cat defines cat as the core domain of the domain model.
package cat

// Cat defines a cat. Weird description.
type Cat struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Breed string `json:"breed"`
}

// NewCat creates a new instance of cat.
func NewCat(id int, name, breed string) *Cat {
	return &Cat{
		ID:    id,
		Name:  name,
		Breed: breed,
	}
}

// CatRepository provides access a cat store.
type CatRepository interface {
	Create(cat *Cat) error
	Find(id int) (*Cat, error)
	FindAll() []*Cat
}
