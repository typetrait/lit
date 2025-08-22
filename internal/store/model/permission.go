package model

type Permission struct {
	ID  int64  `gorm:"primaryKey;not null"`
	Key string `gorm:"unique;not null"`
}
