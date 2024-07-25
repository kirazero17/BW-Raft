package main

import (
	"flag"
	"fmt"
)

func main() {
	nodetype := flag.String("type", "server", "Node type")
	var add = flag.String("address", "", "servers's address")
	var mems = flag.String("members", "", "other members' address")

	flag.Parse()

	switch *nodetype {
	case "server":
		server(*add, *mems)
	case "client":
		client()
	default:
		fmt.Println("Invalid node type ! Exiting...")
	}
}
