package main

import (
	"fmt"
	"github.com/gocantodev/server/config"
)

func main() {
	configuration, err := config.New()

	if err != nil {
		fmt.Println("error: %w", err)
		return
	}

	fmt.Println(configuration)
}
