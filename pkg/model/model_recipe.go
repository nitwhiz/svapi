package model

import (
	"fmt"
	"github.com/hashicorp/go-memdb"
	"github.com/manyminds/api2go/jsonapi"
)

const TypeRecipe = "recipes"

type Recipe struct {
	ID          string              `json:"-"`
	Ingredients []*RecipeIngredient `json:"-"`
	Name        string              `json:"name"`
	IsCooking   bool                `json:"isCooking"`
	OutputItems []*Item             `json:"-"`
	OutputYield int                 `json:"outputYield"`
}

func (r Recipe) TableName() string {
	return TypeRecipe
}

func (r Recipe) SearchIndexContents() [][]string {
	// todo: this needs to be able to process multiple index values; it does now. use it for filter[]
	return [][]string{{r.Name, fmt.Sprintf("%v", r.IsCooking)}}
}

func (r Recipe) Indexes() map[string]*memdb.IndexSchema {
	return map[string]*memdb.IndexSchema{}
}

func (r Recipe) GetID() string {
	return r.ID
}

func (r Recipe) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: TypeRecipeIngredient,
			Name: "ingredients",
		},
		{
			Type: TypeItem,
			Name: "outputItems",
		},
	}
}

func (r Recipe) GetReferencedIDs() []jsonapi.ReferenceID {
	var result []jsonapi.ReferenceID

	for _, ingredient := range r.Ingredients {
		result = append(result, jsonapi.ReferenceID{
			ID:   ingredient.ID,
			Type: TypeRecipeIngredient,
			Name: "ingredients",
		})
	}

	for _, item := range r.OutputItems {
		result = append(result, jsonapi.ReferenceID{
			ID:   item.ID,
			Type: TypeItem,
			Name: "outputItems",
		})
	}

	return result
}
