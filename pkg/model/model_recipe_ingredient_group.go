package model

import (
	"github.com/hashicorp/go-memdb"
	"github.com/manyminds/api2go/jsonapi"
)

const TypeRecipeIngredientGroup = "recipeIngredientGroups"

type RecipeIngredientGroup struct {
	ID          string              `json:"-"`
	Items       []*Item             `json:"-"`
	Ingredients []*RecipeIngredient `json:"-"`
}

func (i RecipeIngredientGroup) TableName() string {
	return TypeRecipeIngredientGroup
}

func (i RecipeIngredientGroup) SearchIndexContents() [][]string {
	var res [][]string

	// todo?

	return res
}

func (i RecipeIngredientGroup) Indexes() map[string]*memdb.IndexSchema {
	return map[string]*memdb.IndexSchema{}
}

func (i RecipeIngredientGroup) GetID() string {
	return i.ID
}

func (i RecipeIngredientGroup) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: TypeItem,
			Name: "items",
		},
		{
			Type: TypeItem,
			Name: "ingredients",
		},
	}
}

func (i RecipeIngredientGroup) GetReferencedIDs() []jsonapi.ReferenceID {
	var res []jsonapi.ReferenceID

	for _, item := range i.Items {
		res = append(res, jsonapi.ReferenceID{
			ID:   item.ID,
			Type: TypeItem,
			Name: "items",
		})
	}

	for _, ingredient := range i.Ingredients {
		res = append(res, jsonapi.ReferenceID{
			ID:   ingredient.ID,
			Type: TypeRecipeIngredient,
			Name: "ingredients",
		})
	}

	return res
}
