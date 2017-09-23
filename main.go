package main

//import "github.com/andrewarrow/simplecoin.life/client"
import "github.com/andrewarrow/simplecoin.life/crypto"

import "fmt"

func main() {
	//client.Setup()
	//n, e, d := crypto.GenKeys()
	//n := uint64(828719)
	//e := uint64(412211)
	//d := uint64(252347)

	n := uint64(44817313)
	e := uint64(22946869)
	d := uint64(34788097)

	fmt.Printf("Public Account Number: %d-%d\n", n, e)
	fmt.Printf("Private Account Number: %d-%d\n", n, d)

	t := crypto.NewTransaction()
	t.Hash = "BLAH"
	t.OwnersPublicAccount = []uint64{n, e}
	fmt.Println(t.Encode())

	/*
		list := crypto.EncodeString("AHELLO", n, e)
		fmt.Println(list)
		for _, c := range list {
			p := crypto.Decode(c, n, d)
			fmt.Println(p)
		}
	*/
}
