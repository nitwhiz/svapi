package model

import (
	"github.com/hashicorp/go-memdb"
	"github.com/nitwhiz/api2go/v2/jsonapi"
)

const TypeCategoryName = "categoryNames"

type CategoryName struct {
	ID       string    `json:"-"`
	Category *Category `json:"-" include:"category"`
	Language *Language `json:"-" include:"language"`
	Name     string    `json:"name"`
}

func (n CategoryName) SearchIndexContents() []string {
	return []string{n.Name}
}

func (n CategoryName) TableName() string {
	return TypeCategoryName
}

func (n CategoryName) Indexes() map[string]*memdb.IndexSchema {
	return map[string]*memdb.IndexSchema{}
}

func (n CategoryName) GetID() string {
	return n.ID
}

func (n CategoryName) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type:         TypeCategory,
			Name:         "category",
			Relationship: jsonapi.ToOneRelationship,
		},
		{
			Type:         TypeLanguage,
			Name:         "language",
			Relationship: jsonapi.ToOneRelationship,
		},
	}
}

func (n CategoryName) GetReferencedIDs() []jsonapi.ReferenceID {
	return BuildReferencedIDs(n)
}

func (n CategoryName) GetReferencedStructs(include []string) []jsonapi.MarshalIdentifier {
	return BuildIncluded(include, n)
}
