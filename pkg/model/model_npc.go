package model

import "github.com/manyminds/api2go/jsonapi"

type Npc struct {
	ID             string    `gorm:"primaryKey" json:"-"`
	BirthdaySeason string    `json:"birthdaySeason"`
	BirthdayDay    int       `json:"birthdayDay"`
	DisplayNames   []NpcName `gorm:"constraint:OnDelete:CASCADE" json:"-"`
	DisplayNameIDs []string  `gorm:"-" json:"-"`
}

func (n Npc) GetID() string {
	return n.ID
}

func (n Npc) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "npcNames",
			Name: "names",
		},
	}
}

func (n Npc) GetReferencedIDs() []jsonapi.ReferenceID {
	var result []jsonapi.ReferenceID

	for _, npcName := range n.DisplayNames {
		result = append(result, jsonapi.ReferenceID{
			ID:   npcName.ID,
			Type: "npcNames",
			Name: "names",
		})
	}

	return result
}
