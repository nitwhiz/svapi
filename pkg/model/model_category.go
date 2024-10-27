package model

import (
	"github.com/hashicorp/go-memdb"
	"github.com/manyminds/api2go/jsonapi"
)

const TypeCategory = "categories"

type Category struct {
	ID         string          `json:"-"`
	InternalID string          `json:"internalId"`
	Items      []*Item         `json:"-"`
	Names      []*CategoryName `json:"-"`
}

func (c Category) SearchIndexContents() [][]string {
	return [][]string{{c.InternalID}}
}

func (c Category) TableName() string {
	return TypeCategory
}

func (c Category) Indexes() map[string]*memdb.IndexSchema {
	return map[string]*memdb.IndexSchema{
		"internalId": {
			Name:    "internalId",
			Indexer: &memdb.StringFieldIndex{Field: "InternalID"},
		},
	}
}

func (c Category) GetID() string {
	return c.ID
}

func (c Category) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: TypeCategoryName,
			Name: "names",
		},
		{
			Type: TypeItem,
			Name: "items",
		},
	}
}

func (c Category) GetReferencedIDs() []jsonapi.ReferenceID {
	var result []jsonapi.ReferenceID

	for _, catName := range c.Names {
		result = append(result, jsonapi.ReferenceID{
			ID:   catName.ID,
			Type: TypeCategoryName,
			Name: "names",
		})
	}

	for _, itemName := range c.Items {
		result = append(result, jsonapi.ReferenceID{
			ID:   itemName.ID,
			Type: TypeItem,
			Name: "items",
		})
	}

	return result
}
