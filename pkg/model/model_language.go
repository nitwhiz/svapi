package model

import (
	"github.com/hashicorp/go-memdb"
	"github.com/nitwhiz/api2go/v2/jsonapi"
)

const TypeLanguage = "languages"

type Language struct {
	ID            string          `json:"-"`
	Code          string          `json:"code"`
	CategoryNames []*CategoryName `json:"-" include:"categoryNames"`
	ItemNames     []*ItemName     `json:"-" include:"itemNames"`
	NpcNames      []*NpcName      `json:"-" include:"npcNames"`
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
			Relationship: jsonapi.ToManyRelationship,
		},
		{
			Type:         TypeItemName,
			Name:         "itemNames",
			Relationship: jsonapi.ToManyRelationship,
		},
		{
			Type:         TypeNpcName,
			Name:         "npcNames",
			Relationship: jsonapi.ToManyRelationship,
		},
	}
}

func (l Language) GetReferencedIDs() []jsonapi.ReferenceID {
	return BuildReferencedIDs(l)
}

func (l Language) GetReferencedStructs(include []string) []jsonapi.MarshalIdentifier {
	return BuildIncluded(include, l)
}
