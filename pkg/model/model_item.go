package model

import (
	"github.com/manyminds/api2go/jsonapi"
)

type Item struct {
	ID           string     `gorm:"primaryKey" json:"-"`
	InternalID   int        `json:"internalId"`
	Category     int        `json:"category"`
	Type         string     `json:"type"`
	DisplayNames []ItemName `gorm:"constraint:OnDelete:CASCADE" json:"-"`
}

func (i Item) GetID() string {
	return i.ID
}

func (i Item) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "itemNames",
			Name: "names",
		},
	}
}

func (i Item) GetReferencedIDs() []jsonapi.ReferenceID {
	var result []jsonapi.ReferenceID

	for _, itemName := range i.DisplayNames {
		result = append(result, jsonapi.ReferenceID{
			ID:   itemName.ID,
			Type: "itemNames",
			Name: "names",
		})
	}

	return result
}
