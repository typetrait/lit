package model

type Permission struct {
	ID  int64  `gorm:"primaryKey;auto_increment;not null"`
	Key string `gorm:"unique;not null"`
}
