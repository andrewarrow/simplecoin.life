package main

//import "github.com/andrewarrow/simplecoin.life/client"
import "github.com/andrewarrow/simplecoin.life/crypto"
import "github.com/andrewarrow/simplecoin.life/words"

import "fmt"

func main() {
	w2 := words.BigWords()
	w3 := words.BigWords()
	w4 := words.BigWords()
	r2 := words.BigWords()
	r3 := words.BigWords()
	r4 := words.BigWords()

	coin := crypto.TransactionList{}
	tx := crypto.NewTransaction(0)
	tx.Owner = w2
	coin.Items = append(coin.Items, tx)
	tx = crypto.NewTransaction(1)
	tx.Owner = w3
	coin.Items = append(coin.Items, tx)
	tx = crypto.NewTransaction(2)
	tx.Owner = w4
	coin.Items = append(coin.Items, tx)
	fmt.Println(coin.Encode())

	coin = crypto.TransactionList{}
	tx = crypto.NewTransaction(0)
	tx.Owner = r2
	coin.Items = append(coin.Items, tx)
	tx = crypto.NewTransaction(1)
	tx.Owner = r3
	coin.Items = append(coin.Items, tx)
	tx = crypto.NewTransaction(2)
	tx.Owner = r4
	coin.Items = append(coin.Items, tx)
	fmt.Println(coin.Encode())
}
