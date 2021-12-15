package connection

import (
	"fmt"
	"gorm.io/gorm"
)

type Connection struct {
	db *gorm.DB
}

func MakeConnection() (Connection, error) {
	db, err := gorm.Open(GetDriver())

	if err != nil {
		return Connection{}, fmt.Errorf("The database connection failed: %v\n", err)
	}

	return Connection{
		db: db,
	}, nil
}
