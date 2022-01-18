package main

import (
	"fmt"
	"github.com/gocantodev/server/database/connection"
	"github.com/voyago/environment"
)

type Migrator struct {
	connection connection.Connection
}

func main() {
	env, err := environment.Make("server")

	if err != nil {
		panic("The given env [server] is invalid.")
	}

	conn, err := connection.Make(env)

	if err != nil {
		panic(err)
	}

	migrator := MakeMigrator(conn)
	migrator.Migrate()
}

func MakeMigrator(connection connection.Connection) Migrator {
	return Migrator{connection: connection}
}

func (receiver Migrator) Migrate() {
	fmt.Println("migrating:", receiver.connection.GetDB().Migrator().CurrentDatabase())
}
