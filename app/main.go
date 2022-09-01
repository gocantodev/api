package main

import (
	"fmt"
	"github.com/gocantodev/server/app/support"
	_ "github.com/lib/pq"
)

func main() {
	seed := "secret"
	//result := "$2a$14$ajq8Q7fbtFRQvXpdCq7Jcuy.Rx1h/L4J60Otx.gyNLbAYctGMJ9tK"

	pass, err := support.MakePassword(seed)

	if err != nil {
		fmt.Println("pass error", err)
		return
	}

	fmt.Println("seed: ", seed)
	fmt.Println("seed: ", pass.GetHash())
	fmt.Println("Valid? ", pass.Is(seed))
	fmt.Println("Length: ", len(pass.GetHash()))
	//
	//
	//connection, err := database.Make(os.Getenv("POSTGRES_URL"))
	//
	//if err != nil {
	//	fmt.Println("There was an issue connecting to the DB: ", err)
	//
	//	return
	//}
	//
	//defer connection.Close()
	//
	//if err := connection.Ping(); err != nil {
	//	fmt.Println("There was an issue pinging to the DB: ", err)
	//	return
	//}
}
