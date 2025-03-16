package model

import (
	"github.com/hashicorp/go-memdb"
	"github.com/nitwhiz/api2go/v2/jsonapi"
)

const TypeNpcName = "npcNames"

type NpcName struct {
	ID       string    `json:"-"`
	Npc      *Npc      `json:"-" include:"npc"`
	Language *Language `json:"-" include:"language"`
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
	return BuildReferencedIDs(n)
}

func (n NpcName) GetReferencedStructs(include []string) []jsonapi.MarshalIdentifier {
	return BuildIncluded(include, n)
}
