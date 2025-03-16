package model

import (
	"github.com/hashicorp/go-memdb"
	"github.com/nitwhiz/api2go/v2/jsonapi"
)

const TypeRecipeIngredientGroup = "recipeIngredientGroups"

type RecipeIngredientGroup struct {
	ID          string              `json:"-"`
	Items       []*Item             `json:"-" include:"items"`
	Ingredients []*RecipeIngredient `json:"-" include:"ingredients"`
}

func (i RecipeIngredientGroup) TableName() string {
	return TypeRecipeIngredientGroup
}

func (i RecipeIngredientGroup) SearchIndexContents() []string {
	return nil
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
			Type:         TypeItem,
			Name:         "items",
			Relationship: jsonapi.ToManyRelationship,
		},
		{
			Type:         TypeItem,
			Name:         "ingredients",
			Relationship: jsonapi.ToManyRelationship,
		},
	}
}

func (i RecipeIngredientGroup) GetReferencedIDs() []jsonapi.ReferenceID {
	return BuildReferencedIDs(i)
}

func (i RecipeIngredientGroup) GetReferencedStructs(include []string) []jsonapi.MarshalIdentifier {
	return BuildIncluded(include, i)
}
