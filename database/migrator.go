package database

import (
	"fmt"
	"gocantoserver/database/connection"
)

type Migrator struct {
	connection connection.Connection
}

func MakeMigrator(connection connection.Connection) Migrator {
	return Migrator{connection: connection}
}

func (receiver Migrator) Migrate() {
	fmt.Println("migrating:", receiver.connection.GetDB().Migrator().CurrentDatabase())
}
