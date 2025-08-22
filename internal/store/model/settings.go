package model

type Settings struct {
	ID    int64  `gorm:"primary_key"`
	Name  string `gorm:"unique;not null"`
	Value string `gorm:"unique;not null"`
}
