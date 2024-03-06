package domainInterfaces

import (
	"github.com/recipes/repositories"

	"github.com/google/uuid"
)


type ds struct {
	store IngredientStore
}

type Ingredient struct {
	id uuid.UUID
	Name string
	Description string
	alternates []uuid.UUID
}

func NewIngredient(name string, description string) (ingredient Ingredient) {
	ing := Ingredient{id: uuid.New(), Name: name, Description: description, alternates: []uuid.UUID{}}
	return ing
}

func (i Ingredient) FindByName(n string) (ingredients []Ingredient, err error) {
	return nil, nil
}