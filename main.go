package main

import (
	"Scouting-2022/secrets"
	"Scouting-2022/server"
)

func main() {
	_, err := server.NewServer(secrets.TOA_API_KEY, "2122-ISR-IIS1")

	if err != nil {
		panic(err)
	}

}
