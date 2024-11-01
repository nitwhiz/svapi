package model

import (
	"github.com/hashicorp/go-memdb"
	"github.com/manyminds/api2go/jsonapi"
)

const TypeLanguage = "languages"

type Language struct {
	ID            string          `json:"-"`
	Code          string          `json:"code"`
	CategoryNames []*CategoryName `json:"-"`
	ItemNames     []*ItemName     `json:"-"`
	NpcNames      []*NpcName      `json:"-"`
}

func (l Language) SearchIndexContents() []string {
	return []string{l.Code}
}

func (l Language) TableName() string {
	return TypeLanguage
}

func (l Language) Indexes() map[string]*memdb.IndexSchema {
	return map[string]*memdb.IndexSchema{
		"code": {
			Name:    "code",
			Unique:  true,
			Indexer: &memdb.StringFieldIndex{Field: "Code"},
		},
	}
}

func (l Language) GetID() string {
	return l.ID
}

func (l Language) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type:         TypeCategoryName,
			Name:         "categoryNames",
			IsNotLoaded:  true,
			Relationship: jsonapi.ToManyRelationship,
		},
		{
			Type:         TypeItemName,
			Name:         "itemNames",
			IsNotLoaded:  true,
			Relationship: jsonapi.ToManyRelationship,
		},
		{
			Type:         TypeNpcName,
			Name:         "npcNames",
			IsNotLoaded:  true,
			Relationship: jsonapi.ToManyRelationship,
		},
	}
}

func (l Language) GetReferencedIDs() []jsonapi.ReferenceID {
	return nil
}
