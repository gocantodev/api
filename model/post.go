package model

import (
	"database/sql"
	"time"
)

type Post struct {
	ID        uint   `gorm:"primaryKey"`
	Uuid      string `gorm:"size:32"`
	Author    Author
	AuthorID  uint
	Slug      string `gorm:"unique"`
	Title     string
	Body      string       `gorm:"type:mediumtext"`
	CreatedAt time.Time    `gorm:"index"`
	UpdatedAt time.Time    `gorm:"index"`
	DeletedAt sql.NullTime `gorm:"index"`
}
