package main

import (
	"Scouting-2022/secrets"
	"Scouting-2022/server"
)

func main() {
	s, err := server.NewServer(secrets.TOA_API_KEY, "2122-ISR-IIC")

	if err != nil {
		panic(err)
	}

	err = s.Run()
	if err != nil {
		panic(err)
	}
}
