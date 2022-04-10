package model

import (
	"github.com/manyminds/api2go/jsonapi"
)

const TasteDislike = "dislike"
const TasteHate = "hate"
const TasteLike = "like"
const TasteLove = "love"
const TasteNeutral = "neutral"

type GiftTaste struct {
	ID     string `gorm:"primaryKey" json:"-"`
	ItemID string `gorm:"uniqueIndex:idx_gift_taste_item_npc_id" json:"-"`
	Item   Item   `gorm:"constraint:OnDelete:CASCADE" json:"-"`
	NpcID  string `gorm:"uniqueIndex:idx_gift_taste_item_npc_id" json:"-"`
	Npc    Npc    `gorm:"constraint:OnDelete:CASCADE" json:"-"`
	Taste  string `json:"taste"`
}

func (g GiftTaste) GetID() string {
	return g.ID
}

func (g GiftTaste) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "items",
			Name: "item",
		},
		{
			Type: "npcs",
			Name: "npc",
		},
	}
}

func (g GiftTaste) GetReferencedIDs() []jsonapi.ReferenceID {
	return []jsonapi.ReferenceID{
		{
			ID:   g.ItemID,
			Type: "items",
			Name: "item",
		},
		{
			ID:   g.NpcID,
			Type: "npcs",
			Name: "npc",
		},
	}
}
