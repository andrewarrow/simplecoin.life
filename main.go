package main

import "github.com/andrewarrow/simplecoin.life/words"
import "github.com/andrewarrow/simplecoin.life/sql"
import "fmt"

var version = "0.1"

func main() {
	sql.InsertTransaction("genesis", "genesis", "root", "", 2779530283277761, 0, 1)
	bundle_id := words.BigWords()
	txid := words.BigWords()
	// Each input transaction contains a signature fragment that proves that you own the private key
	sql.InsertTransaction(txid, bundle_id, "root", "siseneg", -1*2779530283277761, 0, 1151449)
	//sql.InsertBundleInput(bundle_id, txid, 0)
	value := int64(2413941289)
	i := 0
	for {
		if i >= 1151449 {
			break
		}
		fmt.Println(i)
		txid = words.BigWords()
		address := words.BigWords()
		sql.InsertTransaction(txid, bundle_id, address, "", value, i, 1151449)
		//sql.InsertBundleOutput(bundle_id, txid, i)
		i++
	}
}

/*
func main() {
	args := os.Args[1:]
	port := "8666"
	gui := false
	db := sql.UserHomeDir() + "/.scl.db"

	key := ""
	for i, a := range args {
		if i%2 == 0 {
			key = a
		} else {
			if key == "--port" {
				port = a
			} else if key == "--db" {
				db = sql.UserHomeDir() + "/" + a
			} else if key == "--gui" {
				gui = true
			}
		}
	}
	sql.SetDbPath(db)
	go peer.Listen(port)
	go peer.SayHello(version)
	if gui {
		client.Setup()
	} else {
		for {
			time.Sleep(1)
		}
	}
}*/
