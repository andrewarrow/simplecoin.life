package main

import "github.com/andrewarrow/simplecoin.life/sql"
import "github.com/andrewarrow/simplecoin.life/client"
import "os"

//import "github.com/andrewarrow/simplecoin.life/crypto"
import "github.com/andrewarrow/simplecoin.life/peer"
import "time"

//import "fmt"

func main() {
	args := os.Args[1:]
	port := "8666"
	peerUrl := ""
	gui := false
	db := sql.UserHomeDir() + "/.scl.db"
	http := false

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
				db = sql.UserHomeDir() + "/" + a
			} else if key == "--gui" {
				gui = true
			} else if key == "--http" {
				http = true
			}
		}
	}
	sql.SetDbPath(db)
	go peer.Listen(port)
	if http {
		go peer.Listen80()
	}
	go peer.SayHello(peerUrl)
	if gui {
		client.Setup()
	} else {
		for {
			time.Sleep(1)
		}
	}
}
