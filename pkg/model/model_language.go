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
	var res []jsonapi.ReferenceID

	for _, catName := range l.CategoryNames {
		res = append(res, jsonapi.ReferenceID{
			ID:   catName.ID,
			Type: TypeCategoryName,
			Name: "categoryNames",
		})
	}

	for _, catName := range l.CategoryNames {
		res = append(res, jsonapi.ReferenceID{
			ID:   catName.ID,
			Type: TypeItemName,
			Name: "itemNames",
		})
	}

	for _, catName := range l.NpcNames {
		res = append(res, jsonapi.ReferenceID{
			ID:   catName.ID,
			Type: TypeNpcName,
			Name: "npcNames",
		})
	}

	return res
}
