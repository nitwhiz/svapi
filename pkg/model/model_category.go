package model

import (
	"github.com/hashicorp/go-memdb"
	"github.com/nitwhiz/api2go/v2/jsonapi"
)

const TypeCategory = "categories"

type Category struct {
	ID         string          `json:"-"`
	InternalID string          `json:"internalId"`
	Items      []*Item         `json:"-" include:"items"`
	Names      []*CategoryName `json:"-" include:"names"`
}

func (c Category) SearchIndexContents() []string {
	return []string{c.InternalID}
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
			Type:         TypeCategoryName,
			Name:         "names",
			Relationship: jsonapi.ToManyRelationship,
		},
		{
			Type:         TypeItem,
			Name:         "items",
			Relationship: jsonapi.ToManyRelationship,
		},
	}
}

func (c Category) GetReferencedIDs() []jsonapi.ReferenceID {
	return BuildReferencedIDs(c)
}

func (c Category) GetReferencedStructs(include []string) []jsonapi.MarshalIdentifier {
	return BuildIncluded(include, c)
}
