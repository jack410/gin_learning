package main

import "class09/routers"

func main() {
	server := routers.NewServer()

	server.Run(":8888")
}
