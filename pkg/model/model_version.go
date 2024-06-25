package model

type Version struct {
	ID      string `gorm:"primaryKey" json:"-"`
	Version string `json:"-"`
}
