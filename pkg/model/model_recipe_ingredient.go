package model

import (
	"fmt"
	"github.com/hashicorp/go-memdb"
	"github.com/nitwhiz/api2go/v2/jsonapi"
)

const TypeRecipeIngredient = "recipeIngredients"

type RecipeIngredient struct {
	ID              string                 `json:"-"`
	Recipe          *Recipe                `json:"-" include:"recipe"`
	IngredientGroup *RecipeIngredientGroup `json:"-" include:"ingredientGroup"`
	Quantity        int                    `json:"quantity"`
}

func (i RecipeIngredient) TableName() string {
	return TypeRecipeIngredient
}

func (i RecipeIngredient) Indexes() map[string]*memdb.IndexSchema {
	return map[string]*memdb.IndexSchema{}
}

func (i RecipeIngredient) SearchIndexContents() []string {
	return []string{fmt.Sprintf("%d", i.Quantity)}
}

func (i RecipeIngredient) GetID() string {
	return i.ID
}

func (i RecipeIngredient) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type:         TypeRecipe,
			Name:         "recipe",
			Relationship: jsonapi.ToOneRelationship,
		},
		{
			Type:         TypeRecipeIngredientGroup,
			Name:         "ingredientGroup",
			Relationship: jsonapi.ToOneRelationship,
		},
	}
}

func (i RecipeIngredient) GetReferencedIDs() []jsonapi.ReferenceID {
	return BuildReferencedIDs(i)
}

func (i RecipeIngredient) GetReferencedStructs(include []string) []jsonapi.MarshalIdentifier {
	return BuildIncluded(include, i)
}
