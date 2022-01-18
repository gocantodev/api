package model

type Migration struct {
	id    uint `gorm:"primaryKey"`
	name  string
	batch uint
}
