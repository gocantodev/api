package main

import (
	"fmt"
	config2 "github.com/gocantodev/server/config"
	_ "github.com/lib/pq"
)

func main() {
	config, err := config2.MakeConfig()
	//seed := "secret"
	//result := "$2a$14$ajq8Q7fbtFRQvXpdCq7Jcuy.Rx1h/L4J60Otx.gyNLbAYctGMJ9tK"

	//pass, err := Entity.MakePassword(seed)

	if err != nil {
		fmt.Println("error", err)
		return
	}

	fmt.Println(config)

	//id := uuid.New().String()
	//fmt.Println(id, len(id), time.Now().UTC().String())

	//fmt.Println("seed: ", seed)
	//fmt.Println("seed: ", pass.GetHash())
	//fmt.Println("Valid? ", pass.Is(seed))
	//fmt.Println("Length: ", len(pass.GetHash()))
	//
	//
	//connection, err := Database.Make(os.Getenv("POSTGRES_URL"))
	//
	//if err != nil {
	//	fmt.Println("There was an issue connecting to the DB: ", err)
	//
	//	return
	//}
	//
	//repo := Repository.MakeUsersRepository(&connection)
	//
	//user, err := repo.FindByUuid("7d36e0d0-b579-48f2-878a-e4cb6534f15b")
	//
	//if err != nil {
	//	fmt.Println("There was an issue querying the DB: ", err)
	//
	//	return
	//}
	//
	//fmt.Println("user:", user)

	//

	//
	//defer connection.Close()
	//
	//if err := connection.Ping(); err != nil {
	//	fmt.Println("There was an issue pinging to the DB: ", err)
	//	return
	//}
}
