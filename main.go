package main

//import "github.com/andrewarrow/simplecoin.life/client"
import "github.com/andrewarrow/simplecoin.life/crypto"

import "fmt"

func main() {
	//client.Setup()
	//n, e, d := crypto.GenKeys()
	n := uint64(828719)
	e := uint64(412211)
	d := uint64(252347)

	n = uint64(91)
	e = uint64(5)
	d = uint64(29)

	//fmt.Println(e, n, d)
	//crypto.Encode("hello", 828719, 412211)
	//crypto.Encode("hello", 828719, 412211)
	c := crypto.Encode(10, n, e)
	fmt.Println("10*", c, crypto.Decode(c, n, d))
	c = crypto.Encode(90, n, e)
	fmt.Println("90*", c, crypto.Decode(c, n, d))
	c = crypto.Encode(89, n, e)
	fmt.Println("89*", c, crypto.Decode(c, n, d))
	//crypto.Decode(c, n, d)
}
