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
			IsNotLoaded:  true,
			Relationship: jsonapi.ToManyRelationship,
		},
		{
			Type:         TypeItem,
			Name:         "ingredients",
			IsNotLoaded:  true,
			Relationship: jsonapi.ToManyRelationship,
		},
	}
}

func (i RecipeIngredientGroup) GetReferencedIDs() []jsonapi.ReferenceID {
	return nil
}
