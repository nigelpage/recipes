package ingredients

import (
	//"github.com/ajstarks/svgo"
	"github.com/google/uuid"
)

type ds struct {
	store IngredientStore
}

func NewIngredient(name string, description string) (ingredient Ingredient) {
	ing := Ingredient{id: uuid.New(), Name: name, Description: description, alternates: []uuid.UUID{}}
	return ing
}

func (i Ingredient) FindByName(n string) (ingredients []Ingredient, err error) {
	return nil, nil
}