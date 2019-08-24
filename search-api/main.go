package main

import (
	"fmt"
	search "app/search"
	server "app/server"
)

func main() {
	fmt.Println("starting elastic client..")
	search.Start()
	fmt.Println("starting server..")
	server.Serve()
}