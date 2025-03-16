package model

import (
	"github.com/hashicorp/go-memdb"
	"github.com/nitwhiz/api2go/v2/jsonapi"
)

const TypeItemName = "itemNames"

type ItemName struct {
	ID       string    `json:"-"`
	Item     *Item     `json:"-" include:"item"`
	Language *Language `json:"-" include:"language"`
	Name     string    `json:"name"`
}

func (n ItemName) SearchIndexContents() []string {
	return []string{n.Name}
}

func (n ItemName) TableName() string {
	return TypeItemName
}

func (n ItemName) Indexes() map[string]*memdb.IndexSchema {
	return map[string]*memdb.IndexSchema{}
}

func (n ItemName) GetID() string {
	return n.ID
}

func (n ItemName) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type:         TypeItem,
			Name:         "item",
			Relationship: jsonapi.ToOneRelationship,
		},
		{
			Type:         TypeLanguage,
			Name:         "language",
			Relationship: jsonapi.ToOneRelationship,
		},
	}
}

func (n ItemName) GetReferencedIDs() []jsonapi.ReferenceID {
	return BuildReferencedIDs(n)
}

func (n ItemName) GetReferencedStructs(include []string) []jsonapi.MarshalIdentifier {
	return BuildIncluded(include, n)
}
