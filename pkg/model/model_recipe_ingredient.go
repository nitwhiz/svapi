package model

import (
	"github.com/hashicorp/go-memdb"
	"github.com/manyminds/api2go/jsonapi"
)

const TypeRecipeIngredient = "recipeIngredients"

type RecipeIngredient struct {
	ID              string                 `json:"-"`
	Recipe          *Recipe                `json:"-"`
	IngredientGroup *RecipeIngredientGroup `json:"-"`
	Quantity        int                    `json:"quantity"`
}

func (i RecipeIngredient) TableName() string {
	return TypeRecipeIngredient
}

func (i RecipeIngredient) Indexes() map[string]*memdb.IndexSchema {
	return map[string]*memdb.IndexSchema{}
}

func (i RecipeIngredient) SearchIndexContents() [][]string {
	return [][]string{{i.Recipe.ID, i.IngredientGroup.ID}}
}

func (i RecipeIngredient) GetID() string {
	return i.ID
}

func (i RecipeIngredient) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: TypeRecipe,
			Name: "recipe",
		},
		{
			Type: TypeRecipeIngredientGroup,
			Name: "ingredientGroup",
		},
	}
}

func (i RecipeIngredient) GetReferencedIDs() []jsonapi.ReferenceID {
	return []jsonapi.ReferenceID{
		{
			ID:   i.Recipe.ID,
			Type: TypeRecipe,
			Name: "recipe",
		},
		{
			ID:   i.IngredientGroup.ID,
			Type: TypeRecipeIngredientGroup,
			Name: "ingredientGroup",
		},
	}
}
