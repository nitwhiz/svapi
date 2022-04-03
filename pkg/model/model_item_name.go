package model

import "github.com/manyminds/api2go/jsonapi"

type ItemName struct {
	ID       string `gorm:"primaryKey" json:"-"`
	ItemID   string `gorm:"uniqueIndex:idx_item_name_id_lang" json:"-"`
	Language string `gorm:"uniqueIndex:idx_item_name_id_lang" json:"language"`
	Name     string `json:"name"`
}

func (n ItemName) GetID() string {
	return n.ID
}

func (n ItemName) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "items",
			Name: "item",
		},
	}
}

func (n ItemName) GetReferencedIDs() []jsonapi.ReferenceID {
	return []jsonapi.ReferenceID{
		{
			ID:   n.ItemID,
			Type: "items",
			Name: "item",
		},
	}
}
