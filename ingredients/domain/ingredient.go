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

type AccessIngredient interface {
	FindByName(name string) (ingredients []Ingredient, err error)
	Delete() (err error)
	Update() (err error)
	GetAlternateIngredients() (ingredients []Ingredient, err error)
}

func NewIngredient(name string, description string) (ingredient Ingredient) {
	ing := Ingredient{id: uuid.New(), Name: name, Description: description, alternates: []uuid.UUID{}}
	return ing
}