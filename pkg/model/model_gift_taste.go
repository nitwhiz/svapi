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

func (g GiftTaste) SearchIndexContents() []string {
	return []string{g.Taste}
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
			Type:         TypeItem,
			Name:         "item",
			IsNotLoaded:  true,
			Relationship: jsonapi.ToOneRelationship,
		},
		{
			Type:         TypeNpc,
			Name:         "npc",
			IsNotLoaded:  true,
			Relationship: jsonapi.ToOneRelationship,
		},
	}
}

func (g GiftTaste) GetReferencedIDs() []jsonapi.ReferenceID {
	return nil
}
