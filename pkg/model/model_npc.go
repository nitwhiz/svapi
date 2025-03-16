package model

import (
	"fmt"
	"github.com/hashicorp/go-memdb"
	"github.com/nitwhiz/api2go/v2/jsonapi"
	"github.com/nitwhiz/svapi/internal/data"
)

const TypeNpc = "npcs"

type Npc struct {
	ID             string       `json:"-"`
	InternalID     string       `json:"internalId"`
	TextureName    string       `json:"-"`
	BirthdaySeason string       `json:"birthdaySeason"`
	BirthdayDay    int          `json:"birthdayDay"`
	Names          []*NpcName   `json:"-" include:"names"`
	GiftTastes     []*GiftTaste `json:"-" include:"giftTastes"`
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
			Relationship: jsonapi.ToManyRelationship,
		},
		{
			Type:         TypeGiftTaste,
			Name:         "giftTastes",
			Relationship: jsonapi.ToManyRelationship,
		},
	}
}

func (n Npc) GetReferencedIDs() []jsonapi.ReferenceID {
	return BuildReferencedIDs(n)
}

func (n Npc) GetCustomLinks(string) jsonapi.Links {
	return jsonapi.Links{
		"texture": {
			Href: fmt.Sprintf("%s/textures/portraits/%c/%s.png", data.Version, n.TextureName[0], n.TextureName),
		},
	}
}

func (n Npc) GetReferencedStructs(include []string) []jsonapi.MarshalIdentifier {
	return BuildIncluded(include, n)
}
