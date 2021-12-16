package main

import (
	"fmt"
	"github.com/voyago/environment"
	kernel "gocantoserver/app"
)

func main() {
	env, err := environment.Make("server")

	if err != nil {
		panic("The given env [server/.env] is invalid.")
	}

	app, err := kernel.MakeApp(env)

	if err != nil {
		message, _ := fmt.Printf("There was an issue creating the App: %v", err)
		panic(message)
	}

	app.Start()
}
