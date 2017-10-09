package main

import "github.com/andrewarrow/simplecoin.life/client"
import "os"

//import "github.com/andrewarrow/simplecoin.life/crypto"
import "github.com/andrewarrow/simplecoin.life/peer"
import "time"

//import "fmt"

func main() {
	args := os.Args[1:]
	port := "9376"
	peerUrl := ""
	gui := true
	db := client.UserHomeDir() + "/.scl.db"

	key := ""
	for i, a := range args {
		if i%2 == 0 {
			key = a
		} else {
			if key == "--port" {
				port = a
			} else if key == "--peer" {
				peerUrl = a
			} else if key == "--db" {
				db = client.UserHomeDir() + "/" + a
			} else if key == "--gui" {
				gui = false
			}
		}
	}
	go peer.Listen(port)
	go peer.SayHello(peerUrl)
	client.SetDbPath(db)
	if gui {
		client.Setup()
	} else {
		for {
			time.Sleep(1)
		}
	}
}
