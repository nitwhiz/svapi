package model

import (
	"github.com/hashicorp/go-memdb"
	"github.com/manyminds/api2go/jsonapi"
)

const TypeGiftTaste = "giftTastes"

const TasteDislike = "dislike"
const TasteHate = "hate"
const TasteLike = "like"
const TasteLove = "love"
const TasteNeutral = "neutral"

type GiftTaste struct {
	ID    string `json:"-"`
	Npc   *Npc   `json:"-"`
	Item  *Item  `json:"-"`
	Taste string `json:"taste"`
}

func (g GiftTaste) Indexes() map[string]*memdb.IndexSchema {
	return map[string]*memdb.IndexSchema{}
}

func (g GiftTaste) SearchIndexContents() [][]string {
	return [][]string{{g.Npc.ID, g.Item.ID, g.Taste}}
}

func (g GiftTaste) TableName() string {
	return TypeGiftTaste
}

func (g GiftTaste) GetID() string {
	return g.ID
}

func (g GiftTaste) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: TypeItem,
			Name: "item",
		},
		{
			Type: TypeNpc,
			Name: "npc",
		},
	}
}

func (g GiftTaste) GetReferencedIDs() []jsonapi.ReferenceID {
	return []jsonapi.ReferenceID{
		{
			ID:   g.Item.ID,
			Type: TypeItem,
			Name: "item",
		},
		{
			ID:   g.Npc.ID,
			Type: TypeNpc,
			Name: "npc",
		},
	}
}
