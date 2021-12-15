package main

import (
	"fmt"
	"gocantoserver/database/connection"
)

func main() {
	con, err := connection.MakeConnection()

	if err != nil {
		fmt.Println("Connection error")
	}

	fmt.Println(con)
}
