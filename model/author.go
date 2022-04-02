package model

import (
	"database/sql"
	"time"
)

type Author struct {
	ID        uint   `gorm:"primaryKey"`
	Uuid      string `gorm:"size:32"`
	Nickname  string
	CreatedAt time.Time    `gorm:"index"`
	UpdatedAt time.Time    `gorm:"index"`
	DeletedAt sql.NullTime `gorm:"index"`
}
