package model

type NpcName struct {
	ID       string `gorm:"primaryKey" json:"-"`
	NpcID    string `gorm:"uniqueIndex:idx_npc_name_id_lang" json:"-"`
	Language string `gorm:"uniqueIndex:idx_npc_name_id_lang" json:"language"`
	Name     string `json:"name"`
}

func (n NpcName) GetID() string {
	return n.ID
}
