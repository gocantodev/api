package command

import (
	"errors"
	"fmt"
	"github.com/gocantodev/server/database/connection"
	"github.com/gocantodev/server/database/schema"
)

const FRESH = "fresh"

type Runner struct {
	Conn     connection.Connection
	Commands map[string]func(con connection.Connection)
}

func MakeRunner(conn connection.Connection) Runner {
	runner := Runner{}
	commands := make(map[string]func(con connection.Connection), 0)

	commands[FRESH] = Fresh()

	runner.Commands = commands
	runner.Conn = conn

	return runner
}

func (receiver Runner) Run(option string) error {
	command, exists := receiver.Commands[option]

	if !exists {
		return errors.New(fmt.Sprintf("The given option [%v] does not exist", option))
	}

	command(receiver.Conn)

	return nil
}

func Fresh() func(conn connection.Connection) {
	return func(conn connection.Connection) {
		models := schema.Make().GetModels()
		migrator := conn.GetDB().Migrator()

		if err := migrator.DropTable(models...); err != nil {
			fmt.Println(fmt.Sprintf("There was an error dropping the DB tables: %v", err))
		}

		if err := migrator.AutoMigrate(models...); err != nil {
			fmt.Println(fmt.Sprintf("There was an error migrating the DB: %v", err))
		}

		fmt.Println("The DB was successfully migrated...")
	}
}
