package model

import (
	"github.com/hashicorp/go-memdb"
	"github.com/manyminds/api2go/jsonapi"
)

const TypeCategoryName = "categoryNames"

type CategoryName struct {
	ID       string    `json:"-"`
	Category *Category `json:"-"`
	Language *Language `json:"-"`
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
			IsNotLoaded:  true,
			Relationship: jsonapi.ToOneRelationship,
		},
		{
			Type:         TypeLanguage,
			Name:         "language",
			IsNotLoaded:  true,
			Relationship: jsonapi.ToOneRelationship,
		},
	}
}

func (n CategoryName) GetReferencedIDs() []jsonapi.ReferenceID {
	return nil
}
