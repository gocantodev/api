package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("App: ", os.Getenv("APP_NAME"))
}
