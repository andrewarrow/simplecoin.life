package main

//import "github.com/andrewarrow/simplecoin.life/client"
import "github.com/andrewarrow/simplecoin.life/crypto"

//import "fmt"

func main() {
	//client.Setup()
	//n, e, d := crypto.GenKeys()
	//fmt.Println(e, n, d)
	//crypto.Encode("hello", 828719, 412211)
	//crypto.Encode("hello", 828719, 412211)
	c := crypto.Encode("hello", 91, 5)
	crypto.Decode(c, 91, 29)
}
