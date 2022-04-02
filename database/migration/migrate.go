package main

import (
	"fmt"
	"github.com/gocantodev/server/database/connection"
	"github.com/gocantodev/server/database/migration/command"
	"github.com/voyago/environment"
	"os"
)

func main() {
	input, _ := command.Parse(os.Args)

	if input.ShouldReject() {
		fmt.Println("the given command is invalid:", input.GetError())

		return
	}

	env, _ := environment.Make("server")
	conn, err := connection.Make(env)

	if err != nil {
		fmt.Println(fmt.Sprintf("There was an error connecting to the DB: %v", err))
		return
	}

	runner := command.MakeRunner(conn)

	if err := runner.Run(input.GetValue()); err != nil {
		fmt.Println(fmt.Sprintf("There was an error running the given command %s: %v", input.GetValue(), err))
	}
}
