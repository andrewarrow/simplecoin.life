package main

//import "github.com/andrewarrow/simplecoin.life/client"
import "github.com/andrewarrow/simplecoin.life/crypto"
import "fmt"

func main() {
	//client.Setup()
	for {
		_, e, _ := crypto.GenKeys()
		fmt.Println(e)
	}
}
