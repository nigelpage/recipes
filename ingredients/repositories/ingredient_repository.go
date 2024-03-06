package repositories

import (
	//"github.com/ajstarks/svgo"
	//"github.com/google/uuid"
)

type IngredientStore struct {
	name string
}

func NewIngredientStore(name string) *IngredientStore {
	return &IngredientStore {name: name}
	
}