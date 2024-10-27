package model

import (
	"fmt"
	"github.com/hashicorp/go-memdb"
	"github.com/manyminds/api2go/jsonapi"
	"github.com/nitwhiz/svapi/pkg/util"
)

const TypeItem = "items"

type Item struct {
	ID               string                   `json:"-"`
	InternalID       string                   `json:"internalId"`
	TextureName      string                   `json:"-"`
	Category         *Category                `json:"-"`
	Type             string                   `json:"type"`
	IsGiftable       bool                     `json:"isGiftable"`
	IsBigCraftable   bool                     `json:"isBigCraftable"`
	Names            []*ItemName              `json:"-"`
	GiftTastes       []*GiftTaste             `json:"-"`
	IngredientGroups []*RecipeIngredientGroup `json:"-"`
}

func (i Item) SearchIndexContents() [][]string {
	res := [][]string{}

	catId := ""

	if i.Category != nil {
		catId = i.Category.ID
	}

	// todo: correct filtering

	res = append(res, []string{
		fmt.Sprintf("category:%s", catId),
		fmt.Sprintf("isBigCraftable:%s", util.BoolAsString(i.IsBigCraftable)),
	})

	return res

	//return [][]string{{catId, i.InternalID, i.Type}}
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
			Type: TypeItemName,
			Name: "names",
		},
		{
			Type: TypeCategory,
			Name: "category",
		},
		{
			Type: TypeGiftTaste,
			Name: "giftTastes",
		},
		{
			Type: TypeRecipeIngredientGroup,
			Name: "ingredientGroups",
		},
	}
}

func (i Item) GetReferencedIDs() []jsonapi.ReferenceID {
	var result []jsonapi.ReferenceID

	for _, itemName := range i.Names {
		result = append(result, jsonapi.ReferenceID{
			ID:   itemName.ID,
			Type: TypeItemName,
			Name: "names",
		})
	}

	result = append(result, jsonapi.ReferenceID{
		ID:   i.Category.ID,
		Type: TypeCategory,
		Name: "category",
	})

	for _, giftTaste := range i.GiftTastes {
		result = append(result, jsonapi.ReferenceID{
			ID:   giftTaste.ID,
			Type: TypeGiftTaste,
			Name: "giftTastes",
		})
	}

	for _, recipeIngredientGroup := range i.IngredientGroups {
		result = append(result, jsonapi.ReferenceID{
			ID:   recipeIngredientGroup.ID,
			Type: TypeRecipeIngredientGroup,
			Name: "ingredientGroups",
		})
	}

	return result
}

func (i Item) GetCustomLinks(string) jsonapi.Links {
	return jsonapi.Links{
		"texture": {
			Href: fmt.Sprintf("/v1/textures/items/%c/%s.png", i.TextureName[0], i.TextureName),
		},
	}
}
