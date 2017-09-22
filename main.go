package main

//import "github.com/andrewarrow/simplecoin.life/client"
import "github.com/andrewarrow/simplecoin.life/crypto"
import "fmt"

func main() {
	//client.Setup()
	//n, e, d := crypto.GenKeys()
	//fmt.Println(e, n, d)
	fmt.Println(crypto.Encode("hello", 828719, 412211))
	fmt.Println(crypto.Decode("hello", 828719, 252347))
}
