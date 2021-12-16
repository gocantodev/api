package main

import (
	"github.com/voyago/environment"
	"gocantoserver/database"
	"gocantoserver/database/connection"
)

func main() {
	env, err := environment.Make("server")

	if err != nil {
		panic("The given env [server] is invalid.")
	}

	conn, err := connection.Make(env)

	if err != nil {
		panic(err)
	}

	migrator := database.MakeMigrator(conn)
	migrator.Migrate()
}
