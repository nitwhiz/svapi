package model

import "github.com/manyminds/api2go/jsonapi"

type NpcName struct {
	ID       string `gorm:"primaryKey" json:"-"`
	NpcID    string `gorm:"uniqueIndex:idx_npc_name_id_lang" json:"-"`
	Language string `gorm:"uniqueIndex:idx_npc_name_id_lang" json:"language"`
	Name     string `json:"name"`
}

func (n NpcName) GetID() string {
	return n.ID
}

func (n NpcName) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "npcs",
			Name: "npc",
		},
	}
}

func (n NpcName) GetReferencedIDs() []jsonapi.ReferenceID {
	return []jsonapi.ReferenceID{
		{
			ID:   n.NpcID,
			Type: "npcs",
			Name: "npc",
		},
	}
}
