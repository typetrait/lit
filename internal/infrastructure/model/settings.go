package model

type Settings struct {
	ID    int64  `gorm:"primary_key;auto_increment"`
	Name  string `gorm:"unique;not null"`
	Value string `gorm:"unique;not null"`
}
