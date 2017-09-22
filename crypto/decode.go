package crypto

import "fmt"

//import "encoding/binary"
import "math/big"

func Decode(bc big.Int, n, d int64) string {

	//Plaintext = Cd mod n
	// C = Pe mod n
	bd := big.NewInt(d)

	var limit big.Int
	pow := limit.Exp(&bc, bd, nil)
	fmt.Println(pow)

	return ""
}
