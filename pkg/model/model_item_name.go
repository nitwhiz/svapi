package model

import (
	"github.com/hashicorp/go-memdb"
	"github.com/manyminds/api2go/jsonapi"
)

const TypeItemName = "itemNames"

type ItemName struct {
	ID       string    `json:"-"`
	Item     *Item     `json:"-"`
	Language *Language `json:"-"`
	Name     string    `json:"name"`
}

func (n ItemName) SearchIndexContents() [][]string {
	return [][]string{{n.Item.ID, n.Language.ID, n.Name}}
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
			Type: TypeItem,
			Name: "item",
		},
		{
			Type: TypeLanguage,
			Name: "language",
		},
	}
}

func (n ItemName) GetReferencedIDs() []jsonapi.ReferenceID {
	res := []jsonapi.ReferenceID{}

	res = append(res, jsonapi.ReferenceID{
		ID:   n.Item.ID,
		Type: TypeItem,
		Name: "item",
	})

	res = append(res, jsonapi.ReferenceID{
		ID:   n.Language.ID,
		Type: TypeLanguage,
		Name: "language",
	})

	return res
}
