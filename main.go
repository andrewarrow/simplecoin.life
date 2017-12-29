package main

import "github.com/andrewarrow/simplecoin.life/words"
import "github.com/andrewarrow/simplecoin.life/sql"
import "fmt"
import "math"

var version = "0.1"

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func simpleBundle(from, to string, value, balance int64) []map[string]string {
	trunk := items[512]
	branch := items[9003]

	items := []map[string]string{}

	bundle := words.BigWords()
	tx := words.BigWords()
	item := map[string]string{tx: tx, bundle: bundle}
	sql.InsertTransaction(tx, bundle, to, "", value, 0, 2)
	items = append(items, item)

	tx = words.BigWords()
	item = map[string]string{tx: tx, bundle: bundle}
	sql.InsertTransaction(tx, bundle, from, Reverse(from), -1*value, 1, 2)
	items = append(items, item)

	tx = words.BigWords()
	new_address = words.BigWords()
	sql.InsertTransaction(tx, bundle, new_address, "", balance-value, 2, 2)
	item = map[string]string{tx: tx, bundle: bundle, change: new_address}
	items = append(items, item)

	sql.UpdateTransaction(items[0], items[1], trunk)
	sql.UpdateTransaction(items[1], items[2], trunk)
	sql.UpdateTransaction(items[2], trunk, branch)
	sql.InsertBundle(bundle, 2)

	return items
}

func main() {
	sql.InsertTransaction("genesis", "genesis", "root", "", 2779530283277761, 0, 0)

	bundle_id := words.BigWords()
	txid := words.BigWords()
	// Each input transaction contains a signature fragment that proves that you own the private key
	sql.InsertTransaction(txid, bundle_id, "root", "siseneg", -1*2779530283277761, 0, 1151448)
	sql.UpdateTransaction(txid, "genesis", "genesis")
	value := int64(2413941289)
	i := 0
	var address string
	items := []string{}
	for {
		if i >= 1151449 {
			break
		}
		fmt.Println(i)
		txid = words.BigWords()
		address = words.BigWords()
		items = append(items, txid)
		sql.InsertTransaction(txid, bundle_id, address, "", value, i, 1151448)
		sql.UpdateTransaction(txid, "genesis", "genesis")
		i++
	}
	sql.InsertBundle(bundle_id, 1151448)

	/*
		"Trunk" and "Branch" are hashes of other transactions, namely the two transactions that were approved by the transaction you are currently looking at.  */

	trunk := items[512]
	branch := items[9003]

	items = []string{}
	bundle_id = words.BigWords()
	txid = words.BigWords()
	sql.InsertTransaction(txid, bundle_id, address, Reverse(address), -1*2413941289, 0, 2)
	items = append(items, txid)
	txid = words.BigWords()
	new_address := words.BigWords()
	sql.InsertTransaction(txid, bundle_id, new_address, "", 1000, 1, 2)
	items = append(items, txid)
	txid = words.BigWords()
	new_address = words.BigWords()
	sql.InsertTransaction(txid, bundle_id, new_address, "", 2413940289, 2, 2)
	items = append(items, txid)

	sql.UpdateTransaction(items[0], items[1], trunk)
	sql.UpdateTransaction(items[1], items[2], trunk)
	sql.UpdateTransaction(items[2], trunk, branch)
	sql.InsertBundle(bundle_id, 2)

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
