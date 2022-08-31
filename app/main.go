package main

import (
	"fmt"
	"github.com/gocantodev/server/app/database"
	_ "github.com/lib/pq"
	"os"
)

func main() {
	connection, err := database.Make(os.Getenv("POSTGRES_URL"))

	if err != nil {
		fmt.Println("There was an issue connecting to the DB: ", err)

		return
	}

	defer connection.Close()

	if err := connection.Ping(); err != nil {
		fmt.Println("There was an issue pinging to the DB: ", err)
		return
	}
}
