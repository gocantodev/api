package main

import (
	"fmt"
	"github.com/gocantodev/server/database/connection"
	"github.com/gocantodev/server/database/migration/cli"
	"os"
)

func main() {
	input, err := cli.Parse(os.Args)

	if err != nil {
		fmt.Println(fmt.Sprintf("There was an error [%v] while reading the command.", err))

		return
	}

	if input.ShouldRunMigrations() {
		fmt.Println("run migrations from:", input.GetFileName())

		return
	}

	if input.ShouldReject() {
		fmt.Println("the given command is invalid:", input.GetError())

		return
	}

	conn, err := connection.Make(input.GetEnv())

	if err != nil {
		fmt.Println(fmt.Sprintf("The was an error [%v] while connecting to the DB.", err))

		return
	}

	fmt.Println("value:", input.GetValue())
	fmt.Println("connected to:", conn.GetDB().Migrator().CurrentDatabase())
}
