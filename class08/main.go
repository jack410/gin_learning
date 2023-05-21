package main

import "class08/routers"

func main() {
	server := routers.NewServer()

	server.Run(":8888")
}
