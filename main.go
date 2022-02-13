package main

import "Scouting-2022/server"

func main() {
	_, err := server.NewServer()
	if err != nil {
		panic(err)
	}
}
