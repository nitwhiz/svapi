package model

import (
	"fmt"
	"github.com/hashicorp/go-memdb"
	"github.com/manyminds/api2go/jsonapi"
	"github.com/nitwhiz/svapi/internal/data"
	"github.com/nitwhiz/svapi/pkg/flags"
)

const TypeItem = "items"

type Item struct {
	ID               string                   `json:"-"`
	InternalID       string                   `json:"internalId"`
	TextureName      string                   `json:"-"`
	Category         *Category                `json:"-"`
	Type             string                   `json:"type"`
	Flags            []*flags.Flag            `json:"flags"`
	Names            []*ItemName              `json:"-"`
	GiftTastes       []*GiftTaste             `json:"-"`
	IngredientGroups []*RecipeIngredientGroup `json:"-"`
	SourceRecipes    []*Recipe                `json:"-"`
}

func (i Item) SearchIndexContents() []string {
	catId := ""

	if i.Category != nil {
		catId = i.Category.ID
	}

	res := []string{
		i.InternalID,
		i.Type,
		catId,
	}

	res = flags.AppendToIndex(res, i.Flags)

	return res
}

func (i Item) TableName() string {
	return TypeItem
}

func (i Item) Indexes() map[string]*memdb.IndexSchema {
	return map[string]*memdb.IndexSchema{
		"internalId": {
			Name:    "internalId",
			Unique:  true,
			Indexer: &memdb.StringFieldIndex{Field: "InternalID"},
		},
	}
}

func (i Item) GetID() string {
	return i.ID
}

func (i Item) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type:         TypeItemName,
			Name:         "names",
			IsNotLoaded:  true,
			Relationship: jsonapi.ToManyRelationship,
		},
		{
			Type:         TypeCategory,
			Name:         "category",
			IsNotLoaded:  true,
			Relationship: jsonapi.ToOneRelationship,
		},
		{
			Type:         TypeGiftTaste,
			Name:         "giftTastes",
			IsNotLoaded:  true,
			Relationship: jsonapi.ToManyRelationship,
		},
		{
			Type:         TypeRecipeIngredientGroup,
			Name:         "ingredientGroups",
			IsNotLoaded:  true,
			Relationship: jsonapi.ToManyRelationship,
		},
		{
			Type:         TypeRecipeIngredientGroup,
			Name:         "sourceRecipes",
			IsNotLoaded:  true,
			Relationship: jsonapi.ToManyRelationship,
		},
	}
}

func (i Item) GetReferencedIDs() []jsonapi.ReferenceID {
	return nil
}

func (i Item) GetCustomLinks(string) jsonapi.Links {
	return jsonapi.Links{
		"texture": {
			Href: fmt.Sprintf("/%s/textures/items/%c/%s.png", data.Version, i.TextureName[0], i.TextureName),
		},
	}
}
