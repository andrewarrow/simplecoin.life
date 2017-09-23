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
	/*
		n0 := uint64(91)
		e0 := uint64(5)
		d0 := uint64(29)
	*/
	n1 := uint64(44817313)
	e1 := uint64(22946869)
	d1 := uint64(34788097)

	n2 := uint64(828719)
	e2 := uint64(412211)
	//d2 := uint64(252347)

	/*
		n3 := uint64(828719)
		e3 := uint64(412211)
		d3 := uint64(252347)
	*/

	fmt.Printf("Public Account Number: %d-%d\n", n1, e1)
	fmt.Printf("Private Account Number: %d-%d\n", n1, d1)

	// CryptoJS.SHA256(index + previousHash + timestamp + data).toString();
	//t0 := crypto.NewTransaction(0)

	t1 := crypto.NewTransaction(1)
	t1.Hash = fmt.Sprintf("%d-%d", n1, e1)
	t1.OwnersPublicAccount = []uint64{n1, e1}
	t1.PreviousOwnersSignature = []uint64{} // 0's sig of hash
	fmt.Println(t1.Encode())

	t2 := crypto.NewTransaction(2)
	t2.Hash = t1.Hash + fmt.Sprintf(",%d-%d", n2, e2)
	t2.OwnersPublicAccount = []uint64{n2, e2}
	t2.PreviousOwnersSignature = crypto.EncodeString(t2.Hash, n1, e1) // 1's sig of hash
	fmt.Println(t2.Encode())

	/*
		list := crypto.EncodeString("AHELLO", n, e)
		fmt.Println(list)
		for _, c := range list {
			p := crypto.Decode(c, n, d)
			fmt.Println(p)
		}
	*/
}
