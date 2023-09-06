package ingredients

import (
	"github.com/ajstarks/svgo"
	"github.com/google/uuid"
)

type Ingredient struct {
	id uuid.UUID;
	Name string;
	Image svg.SVG;
	Description string;
	alternates []uuid.UUID
}

type IngredientAccess interface {
	FindByName(name string) (ingredients []Ingredient, err error)
	Delete() (err error)
	Update() (err error)
	GetAlternateIngredients() (ingredients []Ingredient, err error)
}

type IngredientStore interface {
	Find(name string) (ingredients []Ingredient, err error)
	Insert() (err error)
	Update() (err error)
	Delete() (err error)
}