package main

//import "github.com/andrewarrow/simplecoin.life/client"
import "github.com/andrewarrow/simplecoin.life/crypto"

import "fmt"

func main() {
	//client.Setup()
	n, e, d := crypto.GenKeys()
	//n := uint64(828719)
	//e := uint64(412211)
	//d := uint64(252347)
	fmt.Printf("Public Account Number: %d-%d\n", n, e)
	fmt.Printf("Private Account Number: %d-%d\n", n, d)

	/*
		list := crypto.EncodeString("AHELLO", n, e)
		fmt.Println(list)
		for _, c := range list {
			p := crypto.Decode(c, n, d)
			fmt.Println(p)
		}
	*/
}
