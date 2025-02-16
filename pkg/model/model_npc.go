package model

import (
	"fmt"
	"github.com/hashicorp/go-memdb"
	"github.com/manyminds/api2go/jsonapi"
	"github.com/nitwhiz/svapi/internal/data"
)

const TypeNpc = "npcs"

type Npc struct {
	ID             string       `json:"-"`
	InternalID     string       `json:"internalId"`
	TextureName    string       `json:"-"`
	BirthdaySeason string       `json:"birthdaySeason"`
	BirthdayDay    int          `json:"birthdayDay"`
	Names          []*NpcName   `json:"-"`
	GiftTastes     []*GiftTaste `json:"-"`
}

func (n Npc) SearchIndexContents() []string {
	return []string{n.InternalID, n.BirthdaySeason, fmt.Sprintf("%d", n.BirthdayDay)}
}

func (n Npc) TableName() string {
	return TypeNpc
}

func (n Npc) Indexes() map[string]*memdb.IndexSchema {
	return map[string]*memdb.IndexSchema{
		"internalId": {
			Name:    "internalId",
			Unique:  true,
			Indexer: &memdb.StringFieldIndex{Field: "InternalID"},
		},
	}
}

func (n Npc) GetID() string {
	return n.ID
}

func (n Npc) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type:         TypeNpcName,
			Name:         "names",
			IsNotLoaded:  true,
			Relationship: jsonapi.ToManyRelationship,
		},
		{
			Type:         TypeGiftTaste,
			Name:         "giftTastes",
			IsNotLoaded:  true,
			Relationship: jsonapi.ToManyRelationship,
		},
	}
}

func (n Npc) GetReferencedIDs() []jsonapi.ReferenceID {
	return nil
}

func (n Npc) GetCustomLinks(string) jsonapi.Links {
	return jsonapi.Links{
		"texture": {
			Href: fmt.Sprintf("%s/textures/portraits/%c/%s.png", data.Version, n.TextureName[0], n.TextureName),
		},
	}
}
