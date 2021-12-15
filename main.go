package main

import "gocantoserver/api"

func main() {
	router := api.MakeRouter()

	router.Start()
}
