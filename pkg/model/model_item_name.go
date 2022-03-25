package model

type ItemName struct {
	ID       string `gorm:"primaryKey" json:"-"`
	ItemID   string `gorm:"uniqueIndex:idx_item_name_id_lang" json:"-"`
	Language string `gorm:"uniqueIndex:idx_item_name_id_lang" json:"language"`
	Name     string `json:"name"`
}

func (n ItemName) GetID() string {
	return n.ID
}
