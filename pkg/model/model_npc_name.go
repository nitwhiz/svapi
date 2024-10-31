package model

import (
	"github.com/hashicorp/go-memdb"
	"github.com/manyminds/api2go/jsonapi"
)

const TypeNpcName = "npcNames"

type NpcName struct {
	ID       string    `json:"-"`
	Npc      *Npc      `json:"-"`
	Language *Language `json:"-"`
	Name     string    `json:"name"`
}

func (n NpcName) SearchIndexContents() []string {
	return []string{n.Name}
}

func (n NpcName) TableName() string {
	return TypeNpcName
}

func (n NpcName) Indexes() map[string]*memdb.IndexSchema {
	return map[string]*memdb.IndexSchema{}
}

func (n NpcName) GetID() string {
	return n.ID
}

func (n NpcName) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type:         TypeNpc,
			Name:         "npc",
			Relationship: jsonapi.ToOneRelationship,
		},
		{
			Type:         TypeLanguage,
			Name:         "language",
			Relationship: jsonapi.ToOneRelationship,
		},
	}
}

func (n NpcName) GetReferencedIDs() []jsonapi.ReferenceID {
	var res []jsonapi.ReferenceID

	res = append(res, jsonapi.ReferenceID{
		ID:   n.Npc.ID,
		Type: TypeNpc,
		Name: "npc",
	})

	res = append(res, jsonapi.ReferenceID{
		ID:   n.Language.ID,
		Type: TypeLanguage,
		Name: "language",
	})

	return res
}
