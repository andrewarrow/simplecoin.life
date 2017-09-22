package main

//import "github.com/andrewarrow/simplecoin.life/client"
import "github.com/andrewarrow/simplecoin.life/crypto"

//import "fmt"

func main() {
	//client.Setup()
	//n, e, d := crypto.GenKeys()
	n := int64(828719)
	e := int64(412211)
	d := int64(252347)

	n = int64(91)
	e = int64(5)
	d = int64(29)

	//fmt.Println(e, n, d)
	//crypto.Encode("hello", 828719, 412211)
	//crypto.Encode("hello", 828719, 412211)
	c := crypto.Encode("hello", n, e)
	crypto.Decode(c, n, d)
}
