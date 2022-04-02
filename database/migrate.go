package main

import (
	"fmt"
	"github.com/gocantodev/server/database/connection"
	"github.com/gocantodev/server/model"
	"github.com/voyago/environment"
)

func main() {
	env, err := environment.Make("server")

	if err != nil {
		fmt.Println("Error resolving env")
		return
	}

	conn, err := connection.Make(env)

	if err != nil {
		fmt.Println("Error connecting to db")
		return
	}

	conn.GetDB().Migrator().DropTable(model.Author{}, model.Post{})

	err = conn.GetDB().AutoMigrate(model.Author{}, model.Post{})

	if err != nil {
		fmt.Println("Error migrating to db")
		return
	}

	fmt.Println("done...")
}
