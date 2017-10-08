package main

import "github.com/andrewarrow/simplecoin.life/client"
import "os"

//import "github.com/andrewarrow/simplecoin.life/crypto"
import "github.com/andrewarrow/simplecoin.life/peer"

//import "fmt"

func main() {
	args := os.Args[1:]
	port := "9376"
	db := client.UserHomeDir() + "/.scl.db"
	if len(args) == 2 && args[0] == "--db" {
		db = client.UserHomeDir() + "/" + args[1]
	} else if len(args) == 2 && args[0] == "--port" {
		port = args[1]
	}
	go peer.Listen(port)
	client.Setup(db)
}
