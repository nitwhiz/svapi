package model

import (
	"fmt"
	"github.com/hashicorp/go-memdb"
	"github.com/manyminds/api2go/jsonapi"
	"github.com/nitwhiz/svapi/pkg/flags"
)

const TypeRecipe = "recipes"

type Recipe struct {
	ID          string              `json:"-"`
	Ingredients []*RecipeIngredient `json:"-"`
	Name        string              `json:"name"`
	Flags       []*flags.Flag       `json:"flags"`
	OutputItems []*Item             `json:"-"`
	OutputYield int                 `json:"outputYield"`
}

func (r Recipe) TableName() string {
	return TypeRecipe
}

func (r Recipe) SearchIndexContents() []string {
	res := []string{
		r.Name,
		fmt.Sprintf("%d", r.OutputYield),
	}

	res = flags.AppendToIndex(res, r.Flags)

	return res
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
			Type:         TypeRecipeIngredient,
			Name:         "ingredients",
			Relationship: jsonapi.ToManyRelationship,
		},
		{
			Type:         TypeItem,
			Name:         "outputItems",
			Relationship: jsonapi.ToManyRelationship,
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
